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
cat > Dockerfile <<EOF
FROM ubuntu:21.04
# deploy gandalf filesystem 
ADD fs4docker.tgz /
# install dependencies
RUN ln -fs /usr/share/zoneinfo/Europe/Paris /etc/localtime
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y apt-utils 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates tzdata
# prepare ebnvironment
EXPOSE 9203
WORKDIR /usr/local/bin
# command to launch at the startup of the container 
CMD rungandalf.sh
EOF
sudo docker rm `sudo docker ps -a | grep "geandalfdocker" | cut -d" " -f1 | grep -v CONTAINER`
sudo docker rmi gandalfdocker
sudo docker build -t gandalfdocker .
