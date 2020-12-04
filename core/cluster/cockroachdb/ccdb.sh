#!/bin/bash

INSTDIR=.
VERSION=v20.1.6

COCKROACH_PKG=cockroach-$VERSION.linux-amd64
COCKROACH_TGZ=$COCKROACH_PKG.tgz

echo Test de fonctionnement CockroachDB
echo ----------------------------------
echo
echo 

installation() {
echo Local installation 
echo 
wget https://binaries.cockroachdb.com/$COCKROACH_TGZ -O /tmp/$COCKROACH_TGZ
tar xf /tmp/$COCKROACH_TGZ -C /tmp/
cp /tmp/$COCKROACH_PKG/cockroach $INSTDIR/
rm -rf /tmp/$COCKROACH_PKG*
if [ -f "$INSTDIR/cockroach" ] 
then 
	echo cockroach installed 
else
	echo ERROR : cockroach installation failed
       	exit 1 
fi
}

build_cluster() {

echo Cockroach iterative cluster installation
echo ----------------------------------------

MEMBERS=""
for i in 1 2 3 4 5 6
do 
	echo "    - Start member $i"
	MEMBER_ADDR=localhost:$(expr 26256 \+ $i)
	if [ -z "$MEMBERS" ] 
	then 
		MEMBERS=$MEMBER_ADDR
	else
		MEMBERS=$MEMBERS,$MEMBER_ADDR
	fi
	cockroach start \
		--insecure \
		--store=node$i \
		--listen-addr=$MEMBER_ADDR \
		--http-addr=localhost:$(expr 8887 \+ $i) \
		--join=$MEMBERS \
		--background
	echo "    - INIT the cluster with the first member"
	[ $i -eq 1 ] && cockroach init --insecure --host=$MEMBERS 
	sleep 5
done
}

workload() {
	cockroach workload init movr 'postgresql://root@localhost:26257?sslmode=disable'
	(cockroach workload run movr --duration=6m 'postgresql://root@localhost:26257?sslmode=disable' > /dev/null ) &
}

failure_restart() {
	echo "!!!! One node failure (number $1)"
	cockroach quit --insecure --host=localhost:$(expr 26256 + $1)
	echo "failure during three minutes" 
	sleep 120 
	echo "node $1 restart"
	cockroach start \
		--insecure \
		--store=node$1 \
		--listen-addr=localhost:$(expr 26256 + $1) \
		--http-addr=localhost:$(expr 8887 \+ $1) \
		--join=$MEMBERS \
		--background
	sleep 120 
	echo "waiting three minutes" 
}

stop_cluster() {
for i in 1 2 3 4 5 6
do 
	echo "  Stopping node $i"
	cockroach quit --insecure --host=localhost:$(expr 26256 \+ $i) &
     
done
}

remove_datastores() {
echo "  Removing datastores"
for i in 1 2 3 4 5 6
do
	rm -r $INSTDIR/node$i
done
echo "  Datastores removed"
}

#installation
build_cluster
workload
failure_restart 3
stop_cluster
remove_datastores

exit 0

