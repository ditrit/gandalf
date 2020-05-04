#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd /home/dev-ubuntu/go/src/core
go build -tags libsqlite3

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./core cluster init cluster 127.0.0.1:9000 
sleep 5
echo 'Join ClusterMember' 
./core cluster join cluster 127.0.0.1:9001 127.0.0.1:9000 
sleep 5
echo 'Join ClusterMember' 
./core cluster join cluster 127.0.0.1:9002 127.0.0.1:9000 
sleep 5

echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./core aggregator Aggregator1 tenant1 127.0.0.1:8000 127.0.0.1:9000 
./core aggregator Aggregator2 tenant1 127.0.0.1:8100 127.0.0.1:9000 
./core aggregator Aggregator3 tenant1 127.0.0.1:8200 127.0.0.1:9000 
sleep 5

echo 'Connector'
echo 'ConnectorMember Con1 and Con2' 
./core connector Connector1 tenant1 127.0.0.1:7000 127.0.0.1:7010 127.0.0.1:8000 
./core connector Connector2 tenant1 127.0.0.1:7100 127.0.0.1:7110 127.0.0.1:8100 
./core connector Connector3 tenant1 127.0.0.1:7200 127.0.0.1:7210 127.0.0.1:8200 
sleep 5

#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test
 
