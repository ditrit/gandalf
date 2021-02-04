#!/bin/bash

BINDIR=/usr/local/bin
COCKROACH=$BINDIR/cockroach
DATADIR=$1
CERTDIR=$DATADIR/certs
STORE=$2
LISTEN_ADDR=$3
HTTP_ADDR=$4
MEMBERS=$5

mkdir -p $DATADIR
cd $DATADIR
$COCKROACH start \
		--certs-dir=$CERTDIR \
		--store=$STORE \
		--listen-addr=$LISTEN_ADDR \
		--http-addr=$HTTP_ADDR \
		--join=$MEMBERS \
		--background