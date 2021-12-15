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

#### Sending of a type event message 

Connector Shoset is only connected to one agregator instances.

The message is sent in same time on each shoset connections;
The message is alson stocked in local waiting queue in order to be able to be consumed by connector workers if needed. 


---- 

## Create a Gandalf Connector

Design Steps:

0. Declare Major/Minor (int64)
1. Pick up Input/Outputs
2. Declare new Worker instance part
3. OAuth and context configure
4. "RegisterCommandFunc()" to create according to commands you want
5. Create functions acccording to commands

### Other parts

**Major/Minor:** Update "Major.Minor" version
    
    v1.0

    Version Major.Minor

**Payload (from function):** Arguments from command

- - - 

## Described Stages

### Pick Up Input/Output

Get scan/input and display it, in Golang :

    input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())


