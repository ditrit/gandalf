#!/bin/bash

BINDIR=/usr/local/bin
COCKROACH=$BINDIR/cockroach
CERTDIR=$1
HOST=$2

$COCKROACH init \
		--certs-dir=$CERTDIR \
		--host=$HOST 