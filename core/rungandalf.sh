#!/bin/bash

cd /usr/local/bin
echo -n Launching cluster 
gandalf start cluster  --offset 1 > /tmp/cluster &
nb=0
while [ "$secret" == "" ]
do 
  secret=`grep "New Aggregator" /tmp/cluster | awk '{print $5}'`
  sleep 2 
  nb=`expr $nb + 2`
  echo -n "."
  if [ $nb -eq 60 ]; then secret="FIN"; fi
done
if [ "$secret" == "FIN" ]
then 
  echo 
  echo "ERROR runnning the cluster"
  exit 1
fi
echo
echo -n Launching aggregator
gandalf start aggregator --offset 4 -l gandalf -t gandalf --bind 0.0.0.0 --cluster 127.0.0.1:9100 --secret $secret > /tmp/aggregator &
nb=0
while [ "$ready" == "" ]
do 
  ready=`grep "Aggregator call done" /tmp/aggregator`
  sleep 2 
  nb=`expr $nb + 2`
  echo -n "."
  if [ $nb -eq 60 ]; then ready="NO"; fi
done
if [ "$ready" == "NO" ]
then 
  echo 
  echo "ERROR runnning the aggregator"
  exit 1
fi
echo 
echo "Gandalf is ready, the API is available on port 9203"
while true; do sleep 3; done

