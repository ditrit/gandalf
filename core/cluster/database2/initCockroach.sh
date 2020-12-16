#!/bin/bash

CERTDIR=$DATADIR/certs
COCKROACH=$BINDIR/cockroach
MEMBER_ADDR=localhost:26257
MEMBERS=$MEMBER_ADDR

$COCKROACH init \
		--certs-dir=$CERTDIR \
		--host=$MEMBERS 