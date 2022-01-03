# Gandalf - Cluster Architecture


Cluster members are used to coordinate and trace any exchange at destination or from aggregators. 

## Communication

### Message Type
See [Messages](../architecture/messages.md)


### Connectivity
Cluster members are access by each aggregator from a Gandalf solution. The cluster have a function to transit type event and commands messages from and to destination to aggregators from same tenant. 

### Routing

#### Receiving/Sending type event message

+ A message is transmited to an aggregator instance
+ The message is followed (stocked into database in cluster level)
+ The message is transmited to every tenant aggregators (except the original)


#### Receiving/Sending type event message

+ A message is transmited by an aggregator instance
+ The message is followed (stocked into database in cluster level)
+ The message is transmited to every tenant aggregators (except the original)


#### Receiving/Sending type command message

+ A message is transmited by an aggregator instance
+ The cluster recovers configuration from pivot languages from any connectors
+ The message is followed (stocked into database in cluster level)
+ The message is transmited to every tenant aggregators

#### Receiving/Sending type config message

+ A message is transmited by an aggregator instance
+ Cluster define destination to give: "What aggregator and what connector will manage this command (according to tenant)".
+ The connector destination is insert into message (Target)
+ The message is followed (stocked into database in cluster level)
+ The message is transmited to every tenant aggregators


#### Receiving/Sending type config message

+ A message is transmited by an aggregator instance
+ Cluster aggrees to secret 
+ The message is transmited to every source aggregators

## Initialization

A cluster is initiate in several steps:

### Communication

Cluster set its shoset.

### Secret

Cluster sends a command to another cluster to validate its secret (except first cluster)

### Logical Configuration

Cluster sends a command to cluster to save and upgrade its logical configuration (except first cluster).

### Database

Cluster creates and init the database if it did not exist (first cluster). Cluster starts a cockroach node.

### API

The cluster starts its API. 

## Database

The database is embedded in differents cluster nodes of gandalf.
For a gandalf node, there is only one cockroach node.

### Connexion Saving

Save cluster connexions to prevent a service interruption.  

### Data Sink

Commands and Events transiting into gandalf system are saved into data sink to manage a trackability. 
