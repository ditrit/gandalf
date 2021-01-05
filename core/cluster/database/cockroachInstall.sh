#!/bin/bash

BINDIR=~/bin
COCKROACH=$BINDIR/cockroach
DATADIR=~/data

if ! [ -f $COCKROACH ]
	then
	BINDIR=~/bin
	COCKROACH_VERSION=v20.1.6
	COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
	COCKROACH_TGZ=$COCKROACH_PKG.tgz
	wget https://binaries.cockroachdb.com/$COCKROACH_TGZ -O /tmp/$COCKROACH_TGZ
	tar xf /tmp/$COCKROACH_TGZ -C /tmp/
	cp /tmp/$COCKROACH_PKG/cockroach $BINDIR/
	rm -rf /tmp/$COCKROACH_PKG*
	fi
	
	rm -r $DATADIR/*