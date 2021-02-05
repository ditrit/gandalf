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
```bash
Cluster :
./gandalf cluster -l Cluster
./gandalf cluster -l Cluster --offset 1 --db_nodename node2 --join 127.0.0.1:9099 --secret TUTU
./gandalf cluster -l Cluster --offset 2 --db_nodename node3 --join 127.0.0.1:9099 --secret TITI
```

```bash
Aggregator :
./gandalf aggregator -l Aggregator1 -t tenant1 --port 10000 --cluster 127.0.0.1:9099 --secret TATA
```

```bash
Connector :
./gandalf connector -l Connector1 --port 10100 --aggregator 127.0.0.1:10000 --secret TOTO --class utils --product Custom
```

## To Do

Test !!