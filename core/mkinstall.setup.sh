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
	 echo "ERROR : can not change owner of $dir to $user"
	 exit 1
      fi
  fi
}

create_dir() {
  dir=$1
  if [ -n "$dir" ] $
  then 
    if [ ! -d "$dir" ] 
      sudo mkdir -p "$dir"
      if [ ! $? ]
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
    mv "$src" "$dst"
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
  cat >/etc/systemd/system/igandalf.service<<EOF
[Unit]
Description=Gandalf service
Requires=Network.target
After=Network.target

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=$INSTDIR/gandalf $mode

[Install]
WantedBy=multi-user.target
EOF
}

howto_start() {
  echo 
  echo "----"
  echo "Gandalf installed in mode ""$mode"""
  echo
  echo "Next steps :"
  echo "1. configure gandalf : "
  echo "   vim /etc/gandalf/gandalf.conf"
  echo "2. start the gandalf service : "
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
move_to ./certs $CONFDIR/
chown_dir $CONFDIR gandalf

# Create DATADIR
create_dir $DATADIR
chown_dir $DATADIR gandalf

# prepare log destination
create_dir $LOGFILE
chown_dir $LOGDIR gandalf

# install gandalf service
create_service
systemctl enable gandalf

# ready to go...
[ $? ] && howto_start