# Documentation - Connectors Architecture

Connectors are used to executed commands from tools.

## Connectors Setup
Connectors are built in 2 levels:

+ Connector router receives and put to use messages into waiting queue.
+ Workers access and emit messages from router by using GRPC.
    + Routing needs a router.
    + A router (thus a connector) have an unique shoset.

## Communication

### Message Type
See [Messages](../architecture/messages.md)

### Connectivity

A router is directly connected to one or several aggregator with a same logical name (allows the disponibility and resilience at aggregator level).

## Routing

Received messages (either of event type, nor of command type) are simply stocked in waiting queue in order to be consume by a worker. No routing to set.
Sent messages (by a worker and thus through a GRPC call) are either commands, or events. 

### Sending of a type event message

+ Connector Shoset is only connected to one agregator instances.
+ The message is sent in same time on each shoset connections;
+ The message is alson stocked in local waiting queue in order to be able to be consumed by connector workers if needed. 

## Receiving a type event message

+ Message is received
+ Message is stocked into router waiting queue

---- 

## Sending of type command message

+ Shoset is connected to several instances of one aggregator
+ Identification of connexions according to these instances: this is a globality of the shoset connexions.
+ We make iterations (in random order) onto these connexions.
+ Sending of command onto the connexion and waiting (while a time: timeout/number of connexions) from a response as an event.
+ If no response in delay, trying the next connexion. 

The waiting is done from a waiting function call into the events waiting queue, within a filter onto the uuid, from the referenced command.

## Receiving of type command message

+ Message is received.
+ Sending a receipt message to the source connector.
+ Message is stockeed into waiting queue into router waiting queue. 

---- 

## Sending of type config message

+ The shoset is connected to several instances of an unique aggregator.
+ Identification of all connexions corresponding to these instances: this is all connexions of the shoset.
+ Sending of config onto a connexion

## Receiving a config message

+ Message is received. Response payload recovering.
+ Update if configuration and to pivot languages.
 
---- 

## Sending a configuration type message

+ Shoset is connected to differents instances of one aggregator. 
+ Identification connexions according to these instances: this is all connexions of the shoset.
+ Sending of configuration onto a connexion


## Receiving a configuration message

+ Message is received. Response payload recovering.
+ Update of logical configuration.
 
---- 

## Sending a secret type message

+ Shoset is connected to differents instances of one aggregator. 
+ Identification connexions according to these instances: this is all connexions of the shoset.
+ Sending of secret onto a connexion


## Receiving a configuration message

+ Message is received.
+ Recovering of response payload to validate or not the secret
 

# Initialization

A connector is initialized by several steps: 

## Communication
A connector sends a command to its shoset, its GRPC server, and connect itself to aggregators instances.

### Secret
The connector sends a command to the cluster by its aggregator to validate its secret.

### Logical Configuration
The connector sends a command to the cluster by its aggregtor to save and update its logical configuration. 

### Admin Worker
The connector starts its admin worker.

### Versions
A connector have one or several worker versions. These versions influence the validaion of the configuration from a connector, also used on the command waiting queues functions.
