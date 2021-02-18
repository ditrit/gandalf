#!/bin/bash
echo 'Running Demo'   

echo 'Build' 
cd "${BASH_SOURCE%/*}/.."
go build -o gandalf

sleep 5
echo 'Cluster' 
echo 'Init ClusterMember' 
./gandalf cluster --offset 1 -l Cluster
./gandalf cluster --offset 2 -l Cluster --join 127.0.0.1:9100 --secret <secret>
./gandalf cluster --offset 3 -l Cluster --join 127.0.0.1:9100 --secret <secret>

echo 'Aggregator' 
echo 'Init AggregatorMember Agg1 and Agg2'
./gandalf aggregator --offset 4 -l Aggregator1 -t tenant1 --cluster 127.0.0.1:9100 --secret <secret>
sleep 5

echo 'Connector'
echo 'ConnectorMember Con1 and Con2' 
./gandalf connector -offset 5 -l Connector1 --aggregator 127.0.0.1:9104 --secret <secret> --class utils --product Custom


./gandalf -g connector -l Connector1 -b 127.0.0.1:7000 -r /tmp/ -a 127.0.0.1:8000 -y Utils -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector2 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Docker -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
./gandalf -g connector -l Connector3 -b 127.0.0.1:7100 -r /tmp/ -a 127.0.0.1:8000 -y Workflow -p Custom -v 1.0 -w $HOME/gandalf/workers -z https://github.com/ditrit/workers/raw/master -s TOTO
sleep 5


echo 'Cli' 
./gandalf cli -e http://localhost:9200 login <login> <password>
./gandalf cli -e http://localhost:9200 create user <username> <email> <password> -t <token>
./gandalf cli -e http://localhost:9200 list user -t <token>
./gandalf cli -e http://localhost:9200 create tenant <tenant> -t <token>
./gandalf cli -e http://localhost:9200 list tenant -t <token>
# CREATE TENANT ADMIN
./gandalf cli -e http://localhost:9200 list tenant -t <token>
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
./gandalf cli -e http://localhost:9200 declare aggregator name <tenant> <name> -t <token>
./gandalf cli -e http://localhost:9200 declare aggregator member <tenant> <name> -t <token>
./gandalf cli -e http://localhost:9200 declare connector name <tenant> <name> -t <token>
./gandalf cli -e http://localhost:9200 declare connector member <tenant> <name> -t <token>


#echo 'Worker'
#./garcimore test send cmd test test
#./garcimore test receive cmd test

#./garcimore test send evt test test test
#./garcimore test receive evt test test

#export AZURE_AUTH_LOCATION=/home/dev-ubuntu/connecteur_azure.auth
