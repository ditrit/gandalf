#!/bin/bash

COCKROACH_VERSION=v20.1.6
COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
COCKROACH_TGZ=$COCKROACH_PKG.tgz
MKSELF_VERSION=2.4.2
MKSELF_NAME=makeself-$MKSELF_VERSION
MKSELF_RELURL=https://github.com/megastep/makeself/releases/download/

TMPDIR=/tmp/gandalf_install
rm -rf $TMPDIR
mkdir $TMPDIR

# Get makeself
wget $MKSELF_RELURL/release-$MKSELF_VERSION/$MKSELF_NAME.run -O /tmp/$MKSELF_NAME.run
chmod a+x /tmp/$MKSELF_NAME.run
pushd /tmp
./$MKSELF_NAME.run
popd

# Download cockroach
wget https://binaries.cockroachdb.com/$COCKROACH_TGZ -O /tmp/$COCKROACH_TGZ
tar xf /tmp/$COCKROACH_TGZ -C /tmp/
cp /tmp/$COCKROACH_PKG/cockroach $TMPDIR/
rm -rf /tmp/$COCKROACH_PKG*
if [ -f "$TMPDIR/cockroach" ] 
then 
	echo cockroach extracted 
else
	echo ERROR : cockroach extracted failed
       	exit 1 
fi

# Build gandalf
[ -f gandalf ] && rm gandalf
go build -o gandalf
if [ -f gandalf ]
then
	cp gandalf $TMPDIR/
else
	echo error building gandalf binary
	exit 1
fi

# copy certs into installation directory (until the internal PKI is ready)
cp -r certs $TMPDIR/
chmod 600 $TMPDIR/node.key
chmod 600 $TMPDIR/client.root.key

# Copy sutup script 
cp mkinstall.setup.sh $TMPDIR/setup.sh
chmod a+x $TMPDIR/setup.sh


/tmp/$MKSELF_NAME/makeself.sh $TMPDIR/ gandalf.sh "Gandalf installer" ./setup.sh

