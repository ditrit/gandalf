#!/bin/bash
set -x

TMPDIR=/tmp/gandalf_dev_docker
TMPINSTALL=/tmp/gandalf_install
INSTDIR=$TMPDIR/usr/local/bin
BDSCRIPTDIR=cluster/database
CONFDIR=$TMPDIR/etc/gandalf
LOGDIR=$TMPDIR/var/log/gandalf
VARDIR=$TMPDIR/var/lib/gandalf
VARDATADIR=$TMPDIR/var/lib/cockroach
[ -d $TMPDIR ] && rm -rf $TMPINSTALL
./mkinstall.sh
mkdir -p $INSTDIR $CONFDIR $LOGDIR $VARDIR $VARDATADIR $INSTDIR/$BDSCRIPTDIR
cp $TMPINSTALL/cockroach $INSTDIR/
cp $TMPINSTALL/gandalf $INSTDIR/
cp rungandalf.sh $INSTDIR/
cp ./$BDSCRIPTDIR/cockroach*.sh $INSTDIR/$BDSCRIPTDIR/
cp -r certs $CONFDIR/
cp -r certs $INSTDIR/
chmod go-rwx $CONFDIR/certs/*
tar czvf fs4docker.tgz -C $TMPDIR .
sudo docker rm `sudo docker ps -a | grep "geandalfdocker" | cut -d" " -f1 | grep -v CONTAINER`
sudo docker rmi gandalfdocker
sudo docker build -t gandalfdocker .
