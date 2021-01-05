#!/bin/bash

BINDIR=~/bin
COCKROACH=$BINDIR/cockroach
DATADIR=$1
CERTDIR=$DATADIR/certs
STORE=$2
LISTEN_ADDR=$3
HTTP_ADDR=$4
MEMBERS=$5

mkdir $DATADIR/database
cd $DATADIR/database
$COCKROACH start \
		--certs-dir=$CERTDIR \
		--store=$STORE \
		--listen-addr=$LISTEN_ADDR \
		--http-addr=$HTTP_ADDR \
		--join=$MEMBERS \
		--background