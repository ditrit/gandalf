#!/bin/bash

CERTDIR=$DATADIR/certs
BINDIR=~/bin
COCKROACH=$BINDIR/cockroach
DATADIR=~/data
MEMBER_ADDR=localhost:26257
MEMBERS=$MEMBER_ADDR

pushd $DATADIR/datastore
$COCKROACH start \
		--certs-dir=$CERTDIR \
		--store=node1 \
		--listen-addr=$MEMBER_ADDR \
		--http-addr=localhost:8888 \
		--join=$MEMBERS \
		--background