#!/bin/bash
set -x

COCKROACH_VERSION=v20.1.6
COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
COCKROACH_TGZ=$COCKROACH_PKG.tgz

TMPDIR=/tmp/gandalf_install
rm -rf $TMPDIR
mkdir $TMPDIR

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

# Copy sutup script 
cp mkinstall.setup.sh $TMPDIR/setup.sh
chmod a+x $TMPDIR/setup.sh


makeself $TMPDIR/ gandalf.sh "Gandalf installer" ./setup.sh
