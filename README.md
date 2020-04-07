# Gandalf Core
La solution Gandalf (Gandalf is A Natural Devops Application Life-cycle Framework) a pour unique objectif de faciliter l’adoption du DevOps sur tout le cycle de vie DevOps sans imposer de choix ou présupposés technologiques ou de produits.

https://ditrit.io/gandalf/

## Schema :
![alt text](images/schemagandalf.png "gandalf schéma")


## Build :

```
go build -tags libsqlite3
```

## Documentation

https://taiga.orness.com/project/xavier-namt/wiki/home

## Demo
```
Cluster :
./garcimore cluster init cluster 127.0.0.1:9000 &
./garcimore cluster join cluster 127.0.0.1:9001 127.0.0.1:9000 &
./garcimore cluster join cluster 127.0.0.1:9002 127.0.0.1:9000 &
```

```
Aggregator :
./garcimore aggregator init agg1 titi 127.0.0.1:8000 127.0.0.1:9000 &
./garcimore aggregator init agg2 titi 127.0.0.1:8100 127.0.0.1:9000 &
./garcimore aggregator join agg1 titi 127.0.0.1:8001 127.0.0.1:9000 127.0.0.1:8000 &
./garcimore aggregator join agg2 titi 127.0.0.1:8101 127.0.0.1:9000 127.0.0.1:8100 &
```

```
Connector :
./garcimore connector init con1 titi 127.0.0.1:7000 127.0.0.1:7010 127.0.0.1:8000 &
./garcimore connector init con2 titi 127.0.0.1:7100 127.0.0.1:7110 127.0.0.1:8100 &
./garcimore connector join con1 titi 127.0.0.1:7001 127.0.0.1:7011 127.0.0.1:8000 127.0.0.1:7000 &
./garcimore connector join con2 titi 127.0.0.1:7101 127.0.0.1:7111 127.0.0.1:8100 127.0.0.1:7100 &
```

```
Test Worker Command :
./garcimore test send cmd test test
./garcimore test receive cmd test
```

```
Test Worker Event :
./garcimore test send evt test test test
./garcimore test receive evt test test
```
## To Do

Test !!