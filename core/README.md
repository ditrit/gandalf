# Gandalf Core
La solution Gandalf (Gandalf is A Natural Devops Application Life-cycle Framework) a pour unique objectif de faciliter l’adoption du DevOps sur tout le cycle de vie DevOps sans imposer de choix ou présupposés technologiques ou de produits.

https://ditrit.io/gandalf/

## Schema :
![alt text](images/schemagandalf.png "gandalf schéma")


## Architecture :

### Cluster :
Le cluster Gandalf trace et fait transiter les commandes et les événements.
### Aggregator :
Les agrégateurs Gandalf cloisonnent et simplifient l’architecture réseau.
### Connector : 
Les connecteurs Gandalf assurent la communication avec les briques du SI.   

## Installation

Gandalf peux être installé sur Ubuntu. Il peux aussi être installé sur les OS sur lesquelles [`dqlite`](https://dqlite.io/) peut être installé.

```
# Cloner ce repository
git clone https://github.com/ditrit/gandalf-core
cd gandalf-core

# Installer les dependences go
go get

## Installer gcc, pour sqlite et dqlite
sudo apt install build-essential

## Installer libsqlite
sudo apt install sqlite3 libsqlite3-dev

## Installer dqlite
### Ubuntu https://dqlite.io/
sudo add-apt-repository -y ppa:dqlite/stable && sudo apt install dqlite
```

## Build :

```bash
go build -tags libsqlite3
```

## Documentation

https://taiga.orness.com/project/xavier-namt/wiki/home


## CLI
L'ensemble d'une solution gandalf est piloté par un unique binaire **'gandalf'**.

gandalf mode command [options]
mode : connector|aggregator|cluster|agent

### Common options :
-c config_file
--config=config_file
config_file : default value is '/etc/gandalf.[json|ini|yaml]'

### Cluster mode usage :
usage:  
gandalf cluster init logical_name bind_address  
gandalf cluster join logical_name bind_address join_address  

*   init command is used to setup a new global Gandalf instance. Output provides the key to be used by super-admin.
*   join command is used to add a new member to an existing cluster


**Fichier de configuration gandalf en mode cluster (by exemple) :**

```bash
mode: cluster
logical_name: toto
bind_address: 192.168.22.10
[join_address : 192.168.22.11]
```

### Aggregator mode usage :
usage:  
gandalf aggregator logical_name tenant bind_address link_address  

**Fichier de configuration gandalf en mode aggregator (by exemple) :**

```bash
mode: aggregator
logical_name: toto
tenant: tata
bind_address: 192.168.22.10
link_address: 192.168.22.11
```

### Connector mode usage :
usage:  
gandalf connector  logical_name tenant bind_address grpc_bind_address link_address  

**Fichier de configuration gandalf en mode connector (by exemple) :**

```bash
mode: connector
logical_name: toto
tenant: tata
bind_address: 192.168.22.10
grpc_bind_address: 192.168.22.11
link_address: 192.168.22.12
```

## Demo
```bash
Cluster :
./gandalf-core cluster init cluster 127.0.0.1:9000 
./gandalf-core cluster join cluster 127.0.0.1:9001 127.0.0.1:9000 
./gandalf-core cluster join cluster 127.0.0.1:9002 127.0.0.1:9000 
```

```bash
Aggregator :
./gandalf-core aggregator Aggregator1 tenant1 127.0.0.1:8000 127.0.0.1:9000
./gandalf-core aggregator Aggregator2 tenant1 127.0.0.1:8100 127.0.0.1:9000
./gandalf-core aggregator Aggregator3 tenant1 127.0.0.1:8200 127.0.0.1:9000
./gandalf-core aggregator Aggregator4 tenant1 127.0.0.1:8300 127.0.0.1:9000
```

```bash
Clone workers and configure them :
cd ~/gandalf/workers/workflow
git clone https://github.com/ditrit/gandalf/connectors/goworkflow.git
go build gandalf-connectors-goworkflow

cd ~/gandalf/workers/Utils
git clone https://github.com/ditrit/gandalf/connectors/goutils.git
go build gandalf-connectors-goutils
```

```bash
Connector :
./gandalf-core connector Connector1 tenant1 127.0.0.1:7000 127.0.0.1:7010 127.0.0.1:8000 Utils
./gandalf-core connector Connector2 tenant1 127.0.0.1:7100 127.0.0.1:7110 127.0.0.1:8100 Workflow
./gandalf-core connector Connector3 tenant1 127.0.0.1:7200 127.0.0.1:7210 127.0.0.1:8200 Azure
./gandalf-core connector Connector4 tenant1 127.0.0.1:7300 127.0.0.1:7310 127.0.0.1:8300 Gitlab
```

```bash
Clone workflow demo and configure it :
git clone git@github.com:ditrit/gandalf-connectors-goworkflow-demo.git

Upload : localhost:3004/gandalf/workflow/upload
Start : localhost:3005/gandalf/app

```
## To Do

Test !!