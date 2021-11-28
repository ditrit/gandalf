#!/bin/bash

BINDIR=/usr/local/bin
COCKROACH=$BINDIR/cockroach
CERTDIR=/etc/gandalf/certs/
HOST=127.0.0.1:9300
DATABASE="gandalf"
PASSWORD="gandalf"

$COCKROACH sql 	\
		--certs-dir=$CERTDIR \
		--host=$HOST \
		--database=$DATABASE