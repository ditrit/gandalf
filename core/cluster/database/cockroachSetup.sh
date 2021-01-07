#!/bin/bash

BINDIR=~/bin
COCKROACH=$BINDIR/cockroach
DATADIR=~/data

CERTDIR=$DATADIR/certs
DSNDIR=$(echo $CERTDIR | sed -e 's:/:%2F:g')
DATASTORE=$DATADIR/datastore
SAFECERTDIR=$DATASTORE/safe
mkdir -p $CERTDIR $DATASTORE $SAFECERTDIR

echo "Cockroach certificates creation"
$COCKROACH cert create-ca \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-node \
127.0.0.1 \
$(hostname) \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-client \
root \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key

$COCKROACH cert create-client \
hydra \
--certs-dir=$CERTDIR \
--ca-key=$SAFECERTDIR/ca.key