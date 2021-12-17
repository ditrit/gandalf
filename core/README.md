# Gandalf Core
Gandalf (Gandalf is A Natural Devops Application Life-cycle Framework), a tool to allow progressive DevOps adoption.

https://ditrit.io


## Table of content
- [Schema](#Schema)
- [Architecture](#Architecture)
- [Installation](#Installation)
- [Build](#Build)
- [Documentation](#Documentation)
- [Docker](#Docker)
- [Getting started](#Getting-started)
- [Demo](#Demo)
- [TODO](#TODO)


## Schema
![alt text](images/schemagandalf.png "gandalf schéma")


## Architecture

### Cluster
The Gandalf cluster traces and routes commands and events.
### Aggregator
Gandalf aggregators silo and simplify the network architecture.
### Connector
Gandalf connectors ensure communication with the bricks of the IS.

## Installation

```
# Clone repository
git clone https://github.com/ditrit/gandalf
cd gandalf

# Installing go dependencies
go get
```

## Build

```bash
go build -o gandalf
```

## Documentation
[Wiki](https://github.com/ditrit/gandalf/wiki).


## Docker
This section show how to start a cluster node and a standard aggregator with docker 

### Howto test Gandalf API using docker
simply run 
```bash 
./prepare_docker.sh
```
It buids a ```gandalfdocker``` image.
You can run a container to use the API endpoint on localhost:9203 :
```bash
run docker run -p 127.0.0.1:9203:9203/tcp gandalfdocker
```
You should obtain a *"hello world"* response in your browser using the adreess *"http://127.0.0.1:9203/ditrit/Gandalf/1.0.0/"*.



## Getting started
This section show how to start 3 cluster node, 1 standard aggregator, 1 custom aggregator, and 1 connector


### Initialisation : 
```bash
./mkinstall.sh
```
```bash
./gandalf.sh cluster
```

### Change user :
```bash
sudo -su gandalf
```

### Cluster : 

**Cluster initialisation**
```bash
./gandalf start cluster --offset 1
```

**Standard Aggregator initilisation**
```bash
./gandalf start aggregator --offset 4 -l gandalf -t gandalf --cluster 127.0.0.1:9100 --secret <secret_output_cluster>
```
**CLI authentification**
```bash
./gandalf cli -e http://localhost:9203 login <login_output_cluster> <password_output_cluster>
```

**Create cluster2 secret** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Cluster 2 start** 
```bash
./gandalf start cluster --offset 2 -l Cluster --join 127.0.0.1:9100 --secret <secret>
```
**Create cluster3 secret** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Cluster 3 start** 
```bash
./gandalf start cluster --offset 3 -l Cluster --join 127.0.0.1:9100 --secret <secret>
```


### Tenant : 

**Create tenant**
```bash
./gandalf cli -e http://localhost:9203 -t <token> create tenant <name> <shortdescription> <description>
```

### Aggregator : 

**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create logicalcomponent <tenant> aggregator <path_to_configuration> 
```
> Aggregator configuration example : 
```yaml
model:
logicalname: aggregator1
type: aggregator
pivot:
  name: aggregator
  major: 1
  minor: 0
keyvalues:
- model:
  value: https://raw.githubusercontent.com/ditrit/gandalf-workers/master 
  key:
    name: repository_url
```
**Create Aggregator secret** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Aggregator start** 
```bash
./gandalf start aggregator --offset 5 -l <name> -t <tenant> --cluster 127.0.0.1:9100 --secret <secret>
```

### Connector :
**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create logicalcomponent <tenant> connector <path_to_configuration> 
```

> Connector configuration example : 
```yaml
model:
logicalname: connector2
type: connector
productconnector:
  name: workflow
  major: 1
  minor: 0
  product:
    name: docker
aggregator: gandalf
keyvalues:
- model:
  value: toto
  key:
    name: totokey
- model:
  value: tata
  key:
    name: tatakey
resources:
- model:
  name: toto
- model:
  name: tata
```


**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Connector start** 
```bash
./gandalf start connector --offset 6 -l <name> --aggregator 127.0.0.1:9103 --secret <secret> --class <class> --product <product>
```


## Demo
This section show how to start a cluster node, a standard aggregator and 2 connectors (github/workflow). Once the 2 connectors started u can upload the example workflow at http://localhost:8080/gandalf/upload to create a repository on Github  

### Initialisation : 
```bash
./mkinstall.sh
```
```bash
./gandalf.sh cluster
```

### Change user :
```bash
sudo -su gandalf
```
### Cluster : 

**Cluster initialisation**
```bash
./gandalf start cluster --offset 1
```
**Standard Aggregator initilisation**
```bash
./gandalf start aggregator --offset 4 -l gandalf -t gandalf --cluster 127.0.0.1:9100 --secret <secret_output_cluster>
```
**CLI authentification**
```bash
./gandalf cli -e http://localhost:9203 login <login_output_cluster> <password_output_cluster>
```

### ConnectorGithub :
**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create logicalcomponent gandalf connector <path_to_configuration> 
```

> Connector configuration example : 
```yaml
model:
logicalname: connectorgithub #Must be in lower case
type: connector
productconnector:
  name: vcs
  major: 1
  minor: 0
  product:
    name: github
aggregator: gandalf
keyvalues:
- model:
  value: <github_access_token>
  key:
    name: token
```


**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Connector start** 
```bash
./gandalf start connector --offset 5 -l connectorgithub --aggregator 127.0.0.1:9103 --secret <secret> --class vcs --product github
```

### ConnectorWorkflow :
**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create logicalcomponent gandalf connector <path_to_configuration> 
```

> Connector configuration example : 
```yaml
model:
logicalname: connectorworkflowdocker #Must be in lower case
type: connector
productconnector:
  name: workflow
  major: 1
  minor: 0
  product:
    name: docker
aggregator: gandalf

```
**Upload configuration** 
```bash
./gandalf cli -e http://localhost:9203 -t <token_output_login> create secret  
```
**Connector start** 
```bash
./gandalf start connector --offset 6 -l connectorworkflowdocker --aggregator 127.0.0.1:9103 --secret <secret> --class workflow --product docker
```

**Upload workflow**

http://localhost:8080/gandalf/upload

```go
package main

import (
	"fmt"

	"github.com/ditrit/gandalf/libraries/goclient"
)

func Workflow(clientGandalf *goclient.ClientGandalf) {

	fmt.Println("SEND COMMMAND CREATE_REPOSITORY")
	payload := `{"Token":"","Name":"TestConnector","Description":"TestConnector","Private":true}`
	commandMessageUUID := clientGandalf.SendCommand("vcs.CREATE_REPOSITORY", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

}
```

## TODO
- Review Authorization API
- Review Enforce
- Application context filter
- Tests
- [Connectors](https://github.com/ditrit/gandalf/wiki/Cr%C3%A9ation-Connecteur).
- Review auto update worker
- Communication between gandalf component
- Auto update gandalf component
- Refactoring code