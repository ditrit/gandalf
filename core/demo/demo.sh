#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd "${BASH_SOURCE%/*}/.."
go build -o gandalf

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./gandalf -g cluster -l Cluster -b 127.0.0.1:9000 
sleep 5
echo 'Join ClusterMember' 
./gandalf -g cluster -l Cluster -b 127.0.0.1:9001 -j 127.0.0.1:9000
sleep 5
echo 'Join ClusterMember' 
./gandalf -g cluster -l Cluster -b 127.0.0.1:9002 -j 127.0.0.1:9000 
sleep 5

echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./gandalf -g aggregator -l Aggregator1 -t tenant1 -b 127.0.0.1:8000 -c 127.0.0.1:9000 -s TATA
sleep 5

echo 'Connector'
echo 'ConnectorMember Con1 and Con2' 
./gandalf -g connector -l Connector1 -b 127.0.0.1:7000 -a 127.0.0.1:8000 -s TOTO
./gandalf -g connector -l Connector1 -b 127.0.0.1:7000 -r /tmp/ -a 127.0.0.1:8000 -y Utils -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector2 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Docker -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector3 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
sleep 5


#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test

#export AZURE_AUTH_LOCATION=/home/dev-ubuntu/connecteur_azure.auth
