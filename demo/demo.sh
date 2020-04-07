#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd /home/dev-ubuntu/go/src/garcimore
go build -tags libsqlite3

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./garcimore cluster init cluster 127.0.0.1:9000 &
sleep 5
echo 'Join ClusterMember' 
./garcimore cluster join cluster 127.0.0.1:9001 127.0.0.1:9000 &
sleep 5
echo 'Join ClusterMember' 
./garcimore cluster join cluster 127.0.0.1:9002 127.0.0.1:9000 &
sleep 5

echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./garcimore aggregator init agg1 titi 127.0.0.1:8000 127.0.0.1:9000 &
./garcimore aggregator init agg2 titi 127.0.0.1:8100 127.0.0.1:9000 &
sleep 5
echo 'Join AggregatorMember Agg1 and Agg2'
./garcimore aggregator join agg1 titi 127.0.0.1:8001 127.0.0.1:9000 127.0.0.1:8000 &
./garcimore aggregator join agg2 titi 127.0.0.1:8101 127.0.0.1:9000 127.0.0.1:8100 &
sleep 5

echo 'Connector'
echo 'Init ConnectorMember Con1 and Con2' 
./garcimore connector init con1 titi 127.0.0.1:7000 127.0.0.1:7010 127.0.0.1:8000 &
./garcimore connector init con2 titi 127.0.0.1:7100 127.0.0.1:7110 127.0.0.1:8100 &
sleep 5
echo 'Join ConnectorMember Con1 and Con2' 
./garcimore connector join con1 titi 127.0.0.1:7001 127.0.0.1:7011 127.0.0.1:8000 127.0.0.1:7000 &
./garcimore connector join con2 titi 127.0.0.1:7101 127.0.0.1:7111 127.0.0.1:8100 127.0.0.1:7100 &
sleep 5

#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test
 
