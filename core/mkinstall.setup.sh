#!/bin/bash

usage() {
  echo "USAGE : $0 <mode>"
  echo "   mode=cluster|aggregator|connector|cli"
}

chown_dir() {
  dir=$1
  user=$2
  if [ -n "$dir" -a -n "$user" ] 
  then
      sudo chown -R $user $dir
      if [ ! $? ]
      then
	 echo "ERROR : can not change owner of $dir to $user"
	 exit 1
      fi
  fi
}

create_dir() {
  dir=$1
  if [ -n "$dir" ] 
  then 
    if [ ! -d "$dir" ] 
    then
      sudo mkdir -p "$dir"
      if [ ! $? ]
      then
	 echo "ERROR : can not create $dir"
	 exit 1
      fi
    fi
  fi 
}

move_to() {
  src=$1
  dst=$2
  if [ -n $1 -a -n $2 ] 
  then
    sudo mv "$src" "$dst"
    if [ ! $? ]
    then 
      echo "ERROR : can not move ""$src"" to ""$dst"""
      exit 1
    fi 
  fi
}

create_user() {
  username=$1
  userhome=$2
  if [ -n "$username" -a -n "$userhome" ] 
  then
    sudo useradd -s /bin/bash -d $userhome -m $username
    if [ ! $? ]
    then 
      echo ERROR : can not create user $username
      exit 1
    fi 
  else 
    [ -z "$username" ] && echo "ERROR : no username provided for user creation"
    [ -z "$userhome" ] && echo "ERROR : no home provided for user $username creation"
    exit 1
  fi
}

INSTDIR=/usr/local/bin
CONFDIR=/etc/gandalf
LOGDIR=/var/log/gandalf
DATADIR=/var/lib/cockroach

# Validate argument
case $1 in 
  "cluster"|"aggregator"|"connector"|"cli")
    mode=$1 
    ;;
  *)
    echo "argument can only be 'cluster', 'aggregator', 'connector' or 'cli'"
    usage
    exit 1
esac

create_service() {
cat >/tmp/gandalf.service<<EOF
[Unit]
Description=Gandalf service
Requires=network.target
After=network.target

[Service]
User=gandalf
RemainAfterExit=yes
ExecStart=$INSTDIR/gandalf $mode

[Install]
WantedBy=multi-user.target
EOF
sudo mv /tmp/gandalf.service /etc/systemd/system/
}

howto_start() {
  echo 
  echo "----"
  echo "Gandalf installed in mode ""$mode"""
  echo
  echo "Next steps :"
  echo "1. configure gandalf : "
  echo "   vim /etc/gandalf/gandalf.conf"
  echo "2. enable the gandalf service : "
  echo "     systemctl enable gandalf"
  echo "3. start the gandalf service : "
  echo "     systemctl start gandailf"
  echo
}

# install binaries
create_dir $INSTDIR
move_to ./gandalf $INSTDIR/
[ "$mode" == "cluster" ] && move_to ./cockroach $INSTDIR

# create user
create_user gandalf /home/gandalf

# create and populate config directory
create_dir $CONFDIR
[ -d $CONFDIR/certs ] && sudo rm -rf $CONFDIR/certs
move_to ./certs $CONFDIR/
chown_dir $CONFDIR gandalf

# Create DATADIR
create_dir $DATADIR
chown_dir $DATADIR gandalf

# prepare log destination
create_dir $LOGDIR
chown_dir $LOGDIR gandalf

# install gandalf service
create_service

# ready to go...
[ $? ] && howto_start

