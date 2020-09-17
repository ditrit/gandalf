#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd "${BASH_SOURCE%/*}/.."
go build -tags libsqlite3 -o gandalf

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./gandalf -g cluster -l Cluster -i Cluster1 -b 127.0.0.1:9000 
sleep 5
echo 'Join ClusterMember' 
./gandalf -g cluster -l Cluster -i Cluster2 -b 127.0.0.1:9001 -j 127.0.0.1:9000 
sleep 5
echo 'Join ClusterMember' 
./gandalf -g cluster -l Cluster -i Cluster3 -b 127.0.0.1:9002 -j 127.0.0.1:9000 
sleep 5

echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./gandalf -g aggregator -l Aggregator1 -i Aggregator1 -t tenant1 -b 127.0.0.1:8000 -c 127.0.0.1:9000
./gandalf -g aggregator -l Aggregator2 -i Aggregator2 -t tenant1 -b 127.0.0.1:8100 -c 127.0.0.1:9000
./gandalf -g aggregator -l Aggregator3 -i Aggregator3 -t tenant1 -b 127.0.0.1:8200 -c 127.0.0.1:9000
./gandalf -g aggregator -l Aggregator4 -i Aggregator4 -t tenant1 -b 127.0.0.1:8300 -c 127.0.0.1:9000
sleep 5

echo 'Connector'
echo 'ConnectorMember Con1 and Con2' 
./gandalf -g connector -l Connector1 -i Connector1 -t tenant1 -b 127.0.0.1:7000 -r 127.0.0.1:7010 -a 127.0.0.1:8000 -y Utils -p Custom -v 1,2 -w $HOME/gandalf/workers -z https://github.com/ditrit/gandalf-workers/raw/master
./gandalf -g connector -l Connector2 -i Connector2 -t tenant1 -b 127.0.0.1:7100 -r 127.0.0.1:7110 -a 127.0.0.1:8100 -y Workflow -p Custom -v 1 -w $HOME/gandalf/workers -z https://github.com/ditrit/gandalf-workers/raw/master
./gandalf -g connector -l Connector3 -i Connector3 -t tenant1 -b 127.0.0.1:7200 -r 127.0.0.1:7210 -a 127.0.0.1:8200 -y Azure
./gandalf -g connector -l Connector4 -i Connector4 -t tenant1 -b 127.0.0.1:7300 -r 127.0.0.1:7310 -a 127.0.0.1:8300 -y VersionControlSystems -p Gitlab -v 1 -w $HOME/gandalf/workers -z https://github.com/ditrit/gandalf-workers/raw/master
sleep 5


#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test

#export AZURE_AUTH_LOCATION=/home/dev-ubuntu/connecteur_azure.auth
