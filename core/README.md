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

```
# Cloner ce repository
git clone https://github.com/ditrit/gandalf
cd gandalf

# Installer les dependences go
go get
```

## Build :

```bash
go build -o gandalf
```

## Documentation


## CLI
L'ensemble d'une solution gandalf est piloté par un unique binaire **'gandalf'**.

gandalf mode command [options]
mode : connector|aggregator|cluster|cli

### Common options :


### Cluster mode usage :
usage:  


**Fichier de configuration gandalf en mode cluster (by exemple) :**

```bash
```

### Aggregator mode usage :
usage:  

**Fichier de configuration gandalf en mode aggregator (by exemple) :**

```bash
```

### Connector mode usage :
usage:  

**Fichier de configuration gandalf en mode connector (by exemple) :**

```bash
```

## Demo

### Cluster : 

**Initialisation Cluster :**
```bash
./gandalf cluster -l Cluster --offset 1
```
**Authentification a la CLI :**
```bash
./gandalf cli -e http://localhost:9200 login <login> <password>
```
**Creation administrateur :** 
```bash
./gandalf cli -e http://localhost:9200 create user <username> <email> <password> -t <token>
```
**Declaration cluster 2 :**
```bash
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
```
**Demarage cluster 2 :** 
```bash
./gandalf cluster -l Cluster --offset 2 --db_nodename node2 --join 127.0.0.1:9100 --secret <secret>
```
**Declaration cluster 3 :**
```bash
./gandalf cli -e http://localhost:9200 declare cluster member -t <token>
```
**Demarage cluster 3 :**
```bash
./gandalf cluster -l Cluster --offset 3 --db_nodename node3 --join 127.0.0.1:9100 --secret <secret>
```

### Tenant : 

**Creation tenant :**
```bash
./gandalf cli -e http://localhost:9200 create tenant <tenant> -t <token>
```

**Creation administrateur tenant :**
```bash
TODO
```

### Aggregateur : 
**Creation aggregateur :** 
```bash
./gandalf cli -e http://localhost:9200 declare aggregator name <tenant> <name> -t <token>
```
**Declaration aggregateur :** 
```bash
./gandalf cli -e http://localhost:9200 declare aggregator member <tenant> <name> -t <token>
```
**Demarage aggregateur :** 
```bash
./gandalf aggregator -l <name> -t <tenant> --port 10000 --cluster 127.0.0.1:9100 --secret <secret>
```

### Connecteur :

**Creation connecteur :** 
```bash
./gandalf cli -e http://localhost:9200 declare connector name <tenant> <name> -t <token>
```

**Declaration connecteur :** 
```bash
./gandalf cli -e http://localhost:9200 declare connector member <tenant> <name> -t <token>
```
**Demarage connecteur :** 
```bash
./gandalf connector -l <name> --port 10100 --aggregator 127.0.0.1:10000 --secret <secret> --class utils --product Custom
```



## To Do

Test !!