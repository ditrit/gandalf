#!/bin/bash

BINDIR=/usr/local/bin
COCKROACH=$BINDIR/cockroach
DATADIR=$1
CERTDIR=$2
STORE=$3
LISTEN_ADDR=$4
HTTP_ADDR=$5
MEMBERS=$6

mkdir -p $DATADIR
cd $DATADIR
$COCKROACH start \
		--certs-dir=$CERTDIR \
		--store=$STORE \
		--listen-addr=$LISTEN_ADDR \
		--http-addr=$HTTP_ADDR \
		--join=$MEMBERS \
		--background