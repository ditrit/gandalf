#!/bin/bash

cmd=$1
suffix=${2:-0}

if [ -z $cmd ]
then 
  echo "Usage : testbackend cmd [storage]"
  echo
  echo "Allow to easily test Gandalf as a backend for Leto using Docker"
  echo 
  echo "Commands:"
  echo "  run     runs a gandalf system (cluster and aggregator) using the given storage"
  echo "  del     delete a storage"
  echo 
  echo "Options:"
  echo "  storage the name of the storage to use. If none, the name used is '0'" 
  echo
fi

volume_cockroach=cockroach_${suffix}
volume_gandalf=gandalf_${suffix}

if [ "$1" == "run" ]
then
  sudo docker run -p 127.0.0.1:9203:9203/tcp -v ${volume_gandalf}:/var/lib/gandalf -v ${volume_cockroach}:/var/lib/cockroach gandalfdocker
fi
if [ "$1" == "del" ]
then
  sudo docker volume rm ${volume_gandalf} 
  sudo docker volume rm ${volume_cockroach}
fi


