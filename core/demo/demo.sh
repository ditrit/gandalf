#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd "${BASH_SOURCE%/*}/.."
go build -o gandalf

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./gandalf cluster -l Cluster --offset 1
./gandalf cluster -l Cluster --offset 2 --db_nodename node2 --join 127.0.0.1:9099 --secret TUTU
./gandalf cluster -l Cluster --offset 3 --db_nodename node3 --join 127.0.0.1:9099 --secret TITI

./gandalf cluster -l Cluster 
./gandalf cluster -l Cluster --join 127.0.0.1:9099 --secret TUTU
./gandalf cluster -l Cluster --join 127.0.0.1:9099 --secret TITI


echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./gandalf aggregator -l Aggregator1 -t tenant1 --port 10000 --cluster 127.0.0.1:9099 --secret TATA
sleep 5

echo 'Connector'
echo 'ConnectorMember Con1 and Con2' 
./gandalf connector -l Connector1 --port 10100 --aggregator 127.0.0.1:10000 --secret TOTO --class utils --product Custom

./gandalf -g connector -l Connector1 -b 127.0.0.1:7000 -r /tmp/ -a 127.0.0.1:8000 -y Utils -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector2 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Docker -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector3 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
sleep 5


echo 'Cli' 
./gandalf cli -e http://localhost:9200 login
./gandalf cli -e http://localhost:9200 create user <username> <email> <password> -t <token>
./gandalf cli -e http://localhost:9200 list user -t <token>
./gandalf cli -e http://localhost:9200 create tenant <tenant> -t <token>
./gandalf cli -e http://localhost:9200 list tenant -t <token>


#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test

#export AZURE_AUTH_LOCATION=/home/dev-ubuntu/connecteur_azure.auth
