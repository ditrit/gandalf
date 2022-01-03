# Gandalf - Libraries

The library allows to communicate in gRPC with the router part of the connector.
Instantiation

## NewClientGandalf(identity, timeout string, clientConnections []string)

+ identity: Logical name of the connector
+ timeout: Maximum message sending timeout
+ clientConnections: bind addresses Grpc connectors

## Operation

```
SendCommand(connectorType, command, timeout, payload string)
```
Allows command type message to be sent
```
SendAdminCommand(connectorType, command, timeout, payload string)
```
Allows slimming message to be sent
```
SendStop(major, minor int64)
```
Allows message sending to stop a version of a connector
```
SendEvent(topic, event, referenceUUID, timeout, payload string)
```
Allows sending of event type message
```
SendCommandList(major, minor int64, commands []string)
```
Allows sending of the version and the list of commands of a worker
```
CreateIteratorCommand()
```
Allows creation of a command iterator
```
CreateIteratorEvent()
```
Allows the creation of an iterator of type event
```
WaitCommand(command, idIterator string, version int64)
```
Allows you to wait for a command type message according to a version
```
WaitEvent(topic, event, referenceUUID, idIterator string)
```
Allows waiting for an event type message
```
WaitTopic(topic, referenceUUID, idIterator string)
```
Allows you to wait for a message from a topic

## Creation
### Recovery of the protobuf project

Just clone the protobuff project

## Compilation

To compile the project simply read the documentation

Note: There is a compiled version of the project on gogrpc

## Implementation of functions

Once the protobuf files are recovered and compiled, it is enough to implement the functions one to one.

Note:

+ The SendCommand, SendReply and SendEvent functions will iterate on the connection list to send the message. If no acceptance validation is received on during the timeout period then a new connection is chosen.
+ The other functions iterates on the connections thanks to a getClientIndex function to not send all the messages to the same connector.
