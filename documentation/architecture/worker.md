# Gandalf - Worker

Workers are the functional part of a connector.
## How workers work

Workers work with three components: A client, a worker base, and the worker.
## Gandalf Client

See [Library](./library.md)

## Worker base

The worker base allows you to abstain from some features of the worker, such as the initialization of the client, and the sending of commands managed by the worker.

It is a structure composed of:

``` type Worker struct {
	major             int64
	minor             int64
	identity          string
	timeout           string
	connections       []string
	clientGandalf     *goclient.ClientGandalf
	OngoingTreatments *gomodels.OngoingTreatments
	WorkerState       *gomodels.WorkerState
	CommandsFuncs     map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int
	EventsFuncs       map[gomodels.TopicEvent]func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int
	//Start             func() *goclient.ClientGandalf
	Stop         func(clientGandalf *goclient.ClientGandalf, major, minor int64, workerState *gomodels.WorkerState)
	SendCommands func(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string) bool
	//Execute      func()
}
```

### Worker

The worker contains the implementation of the desired functions for the connector.
## Creation of workers

To create workers, it is necessary to have a gandalf client and a worker base in the chosen programming language.
Gandalf Client

See [Library](./library.md)

### Worker base

See above

### Worker

To create a worker:

+ Initialize the major and minor version
```
var major = int64(1)
var minor = int64(0)
```
+ Retrieve the values of the standard input in json
```
input:= bufio.NewScanner(os.Stdin)
input. Scan()
```
+ The worker base must be initialized
```
worker:= worker.NewWorker(major, minor)
```
+ It is then enough to implement the desired functions for the connector.
```
{unction_name}(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command)
{unction_name}(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event)
```
+ And finally save the functions to the worker base

```
worker.RegisterCommandsFuncs("{Job Name}", {Job Name})
```
