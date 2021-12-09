#!/bin/bash

cd /usr/local/bin
start_aggregator="/var/lib/gandalf/start_aggregator.sh"
echo -n Launching cluster 
gandalf start cluster  --offset 1 > /tmp/cluster &
nb=0

while [ "$secret" == "" -a "$called" == "" ]
do 
  secret=`grep "New Aggregator" /tmp/cluster | awk '{print $5}'`
  called=`grep "Cluster call done" /tmp/cluster`
  sleep 2 
  nb=`expr $nb + 2`
  echo -n "."
  if [ $nb -eq 60 ]; then secret="FIN"; fi
done

if [ "$secret" == "FIN" ]
then 
  echo 
  echo "ERROR runnning the cluster"
  sleep 1000
  exit 1
fi

if [ -n "$secret" ] 
then 
  cat > ${start_aggregator} <<EOF
#!/bin/sh
gandalf start aggregator --offset 4 -l gandalf -t gandalf --bind 0.0.0.0 --cluster 127.0.0.1:9100 --secret ${secret} > /tmp/aggregator &
EOF
chmod a+x ${start_aggregator}
fi

echo
echo -n Launching aggregator
${start_aggregator}
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

