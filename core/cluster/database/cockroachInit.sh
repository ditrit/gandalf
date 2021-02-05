#!/bin/bash

BINDIR=/usr/local/bin
COCKROACH=$BINDIR/cockroach
DATADIR=$1
CERTDIR=$2
HOST=$2

$COCKROACH init \
		--certs-dir=$CERTDIR \
		--host=$HOST 