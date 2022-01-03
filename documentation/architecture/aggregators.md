# Gandalf - Aggregator Architecture

The function of the aggregator is to aggregate the different connectors according to the network topology and to partition the tenants.

## Communication

### Message Type
See [Messages](../architecture/messages.md)

## Connectivity

An aggregator is connected directly to members of a Gandalf cluster and is accessed by a set of connectors (via their routers).

## Routing

### Receiving/Sending an Event Type Message

The message and its origin (its original connection) are read on the tail of the events

+ If the tenant indicated in the event is not the one defined for the aggregator, the event is not transmitted. However, an event of type "error" is sent (only to the cluster?).

The message is sent simultaneously on each of the connections of the shoset, except on the one from which it comes.

### Receiving/Sending a command type message

The message, the type of its origin ("connector" or "cluster") and its origin (its original connection) are read on the tail of the commands.

+ If the tenant indicated in the event is not the one defined for the aggregator, the command is not transmitted. However, an event of type "error" is sent (only to the cluster?).
+ If the source of the command is "connector", the connections to the cluster are selected.
+ If the source of the command is of type "cluster", the destination (logical name of the connector to which the message is routed) is extracted from the message and connections to that destination are selected.

We iterate on the selected connections (in the same way as in the case of connectors):

+ Sending the order on the connection and waiting (how long?) for an event in response
+ If no response within the delay, try the next connection.

### Receiving/Sending a config type message

The message, the type of its origin ("connector" or "cluster") and its origin (its original connection) are read on the config queue.

+ If the tenant indicated in the event is not the one defined for the aggregator, the config is not transmitted. However, an event of type "error" is sent (only to the cluster?).
+ If the source of the config is of type "connector", the connections to the cluster are selected.
+ If the source of the config is of type "cluster", the destination (logical name of the connector to which the message is routed) is extracted from the message and connections to that destination are selected.
+ Sending the config to a connection.

### Receiving/Sending a Configuration Type Message

The message, the type of its origin ("connector" or "cluster") and its origin (its original connection) are read on the tail of the configurations.

+ If the source of the configuration is "connector", the connections to the cluster are selected.
+ If the source of the configuration is of type "cluster", the destination (logical name of the connector to which the message is routed) is extracted from the message and connections to that destination are selected.
+ Sending the config to a connection.

### Receiving/Sending a configuration Database Type Message

The aggregater receives the configuration and creates a connection to the database.

### Receiving/Sending a Secret Type Message

The message, the type of its origin ("connector" or "cluster") and its origin (its original connection) are read on the tail of the secrets.

+ If the source of the secret is of type "connector", the connections to the cluster are selected.
+ If the source of the secret is of type "cluster", the destination (logical name of the connector to which the message is routed) is extracted from the message and connections to that destination are selected.
+ Sending the config to a connection.


## Initialization

An aggregater is initialized in several steps:
### Communication

The aggregater sets up its Shoset, and connects to the cluster instances.
### Secret

The aggregater sends a command to the cluster to validate its secret.
### Logical Configuration

The aggregater sends a command to the cluster to save and update its logical configuration.
### Database

The aggregater sends a command to the cluster to retrieve the database connection information.
### API

The aggregater starts its API
