package database2

func installCockroach() {
	`if ! [ -f $COCKROACH ]
	then
	COCKROACH_VERSION=v20.1.6
	COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
	COCKROACH_TGZ=$COCKROACH_PKG.tgz
	wget https://binaries.cockroachdb.com/$COCKROACH_TGZ -O /tmp/$COCKROACH_TGZ
	tar xf /tmp/$COCKROACH_TGZ -C /tmp/
	cp /tmp/$COCKROACH_PKG/cockroach $BINDIR/
	rm -rf /tmp/$COCKROACH_PKG*
	fi
	
	rm -r $DATADIR/*`
}

func setupCoackroach() {
	`echo "Cockroach setup"
	MEMBER_ADDR=localhost:26257
	MEMBERS=$MEMBER_ADDR
	
	CERTDIR=$DATADIR/certs
	DSNDIR=$(echo $CERTDIR | sed -e 's:/:%2F:g')
	DATASTORE=$DATADIR/datastore
	SAFECERTDIR=$DATASTORE/safe
	mkdir -p $CERTDIR $DATASTORE $SAFECERTDIR
	A
	echo "Cockroach certificates creation"
	$COCKROACH cert create-ca \
	--certs-dir=$CERTDIR \
	--ca-key=$SAFECERTDIR/ca.key
	
	$COCKROACH cert create-node \
	localhost \
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
	--ca-key=$SAFECERTDIR/ca.key`
}

func startCoackroach() {
	`echo "Cockroach start"
	pushd $DATADIR/datastore
	$COCKROACH start \
			--certs-dir=$CERTDIR \
			--store=node1 \
			--listen-addr=$MEMBER_ADDR \
			--http-addr=localhost:8888 \
			--join=$MEMBERS \
			--background`
}
