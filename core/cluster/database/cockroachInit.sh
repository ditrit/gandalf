#!/bin/bash

BINDIR=/usr/local/bin
DATADIR=$1
CERTDIR=$DATADIR/certs
COCKROACH=$BINDIR/cockroach
HOST=$2

$COCKROACH init \
		--certs-dir=$CERTDIR \
		--host=$HOST 