# Gandalf Core
La solution Gandalf (Gandalf is A Natural Devops Application Life-cycle Framework) a pour unique objectif de faciliter l’adoption du DevOps sur tout le cycle de vie DevOps sans imposer de choix ou présupposés technologiques ou de produits.

https://ditrit.io/gandalf/


## Table of content
- [Schema](#Schema)
- [Architecture](#Architecture)
- [Installation](#Installation)
- [Build](#Build)
- [Documentation](#Documentation)
- [Getting started](#Getting-started)


## Schema
![alt text](images/schemagandalf.png "gandalf schéma")


## Architecture

### Cluster
Le cluster Gandalf trace et fait transiter les commandes et les événements.
### Aggregator
Les agrégateurs Gandalf cloisonnent et simplifient l’architecture réseau.
### Connector
Les connecteurs Gandalf assurent la communication avec les briques du SI.   


## Installation

```
# Cloner ce repository
git clone https://github.com/ditrit/gandalf
cd gandalf

# Installer les dependences go
go get
```

## Build

```bash
go build -o gandalf
```

## Documentation
[Wiki](https://github.com/ditrit/gandalf/wiki).

## Getting started

### Cluster : 

**Initialisation Cluster**
```bash
./gandalf start cluster --offset 1 -l Cluster 
```
**Authentification a la CLI**
```bash
./gandalf cli -e http://localhost:9200 login <login> <password>
```
**Creation administrateur** 
```bash
./gandalf cli -e http://localhost:9200 create user <username> <email> <password> -t <token>
```
**Declaration cluster 2**
```bash
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
```
**Demarage cluster 2** 
```bash
./gandalf start cluster --offset 2 -l Cluster --join 127.0.0.1:9100 --secret <secret>
```
**Declaration cluster 3**
```bash
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
```
**Demarage cluster 3**
```bash
./gandalf start cluster --offset 3 -l Cluster --join 127.0.0.1:9100 --secret <secret>
```

### Tenant : 

**Creation tenant**
```bash
./gandalf cli -e http://localhost:9200 create tenant <tenant> -t <token>
```
BLALBLA TENANT


### Aggregateur : 
**Authentification a la CLI**
```bash
./gandalf cli -e http://localhost:9203 login <login> <password>
```
**Declaration aggregateur** 
```bash
./gandalf cli -e http://localhost:9203 declare aggregator member <tenant> <name> -t <token>
```
**Demarage aggregateur** 
```bash
./gandalf start aggregator --offset 4 -l <name> -t <tenant> --cluster 127.0.0.1:9100 --secret <secret>
```

### Connecteur :
**Creation connecteur** 
```bash
./gandalf cli -e http://localhost:9203 declare connector name <tenant> <name> -t <token>
```

**Declaration connecteur** 
```bash
./gandalf cli -e http://localhost:9203 declare connector member <tenant> <name> -t <token>
```
**Demarage connecteur** 
```bash
./gandalf start connector --offset 5 -l <name> --aggregator 127.0.0.1:9103 --secret <secret> --class <class> --product <product>
```
