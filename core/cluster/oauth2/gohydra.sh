#!/bin/bash

BINDIR=~/bin
HYDRA=$BINDIR/hydra
COCKROACH=$BINDIR/cockroach
DATADIR=~/data

echo "Complete reinitialisation"
killall hydra
killall cockroach
sleep 10

echo "check binaries"
if ! [ -f $HYDRA ]
then 
  echo "build Hydra binary"
  ### BUILD HYDRA as explained in hydra docs :
  #go get -d -u github.com/ory/hydra
  ## workaround to install astutil package (doesn't find the right path)
  ### mkdir -p /home/gandalf/go/src/golang.org/x/tools/go/ast
  ### ln -s /usr/lib/go-1.13/src/cmd/vendor/golang.org/x/tools/go/ast/astutil /home/gandalf/go/src/golang.org/x/tools/go/ast/astutil
  #$ go install github.com/gobuffalo/packr/v2/packr2
  #$ cd $(go env GOPATH)/src/github.com/ory/hydra
  #$ GO111MODULE=on make install-stable
  #$ $(go env GOPATH)/bin/hydra help
  ## Warning : no version given this way
  ## Simpler way is to use install.sh on a new branch using the version tag (v1.8.5) :
  pushd $(go env GOPATH)
  git clone https://github.com/ory/hydra.git
  cd hydra
  git checkout -b v1.8.5 v1.8.5
  bash install.sh
  mv bin/hydra $HYDRA
  popd
fi

if ! [ -f $COCKROACH ]
then
COCKROACH_VERSION=v20.1.6
COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
COCKROACH_TGZ=$COCKROACH_PKG.tgz
wget https://binaries.cockroachdb.com/$COCKROACH_TGZ -O /tmp/$COCKROACH_TGZ
tar xf /tmp/$COCKROACH_TGZ -C /tmp/
cp /tmp/$COCKROACH_PKG/cockroach $BINDIR/
rm -rf /tmp/$COCKROACH_PKG*
fi

rm -r $DATADIR/*

echo "Cockroach setup"
MEMBER_ADDR=localhost:26257
MEMBERS=$MEMBER_ADDR

CERTDIR=$DATADIR/certs
DSNDIR=$(echo $CERTDIR | sed -e 's:/:%2F:g')
DATASTORE=$DATADIR/datastore
SAFECERTDIR=$DATASTORE/safe
mkdir -p $CERTDIR $DATASTORE $SAFECERTDIR
A
echo "Cockroach certificates creation"
$COCKROACH cert create-ca \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-node \
localhost \
$(hostname) \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-client \
root \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-client \
hydra \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

echo "Cockroach start"
pushd $DATADIR/datastore
$COCKROACH start \
		--certs-dir=$CERTDIR \
		--store=node1 \
		--listen-addr=$MEMBER_ADDR \
		--http-addr=localhost:8888 \
		--join=$MEMBERS \
		--background

echo "Cockroach cluster initialisation"
$COCKROACH init \
		--certs-dir=$CERTDIR \
		--host=$MEMBERS 

echo "Hydra database and user creation"
$COCKROACH sql 	\
		--certs-dir=$CERTDIR \
		 --host=$MEMBER_ADDR <<EOF
CREATE DATABASE IF NOT EXISTS hydra;
CREATE USER IF NOT EXISTS hydra WITH PASSWORD 'hydra';
GRANT ALL ON DATABASE hydra TO hydra;
EOF


echo "HYDRA configuration"

export DSN=cockroach://hydra:hydra@localhost:26257/hydra?sslcert=$DSNDIR%2Fclient.hydra.crt\&sslkey=$DSNDIR%2Fclient.hydra.key\&sslmode=verify-full\&sslrootcert=$DSNDIR%2Fca.crt
export SECRETS_SYSTEM=$(export LC_CTYPE=C; cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
export URLS_SELF_ISSUER=https://localhost:9000/ 
export URLS_CONSENT=http://localhost:9020/consent 
export URLS_LOGIN=http://localhost:9020/login 

echo "Hydra database setup"
$HYDRA migrate sql $DSN --yes 

echo "Hydra start"

$HYDRA serve all


popd
exit 0

