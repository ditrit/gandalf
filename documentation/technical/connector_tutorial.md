# Tutorial to create a Gandalf connector.
## Structure

A connector is separated into two parts:

+ The Router located in gandalf/core/connector takes care of communication with the aggregator as well as the management and communication of Grpc workers.
+ The workers locate in gandalf/connectors/ takes care of the interaction with the tools

Note: The router part is the same for all Gandalf connectors
Schema:
## Creation of a new connector

+ Step 1: Creation of a new project in gandalf/connectors/
+ Step 2: Creation of the worker
+ Step 3: Create the pivot language in the gandalf-worker/configurations//<version_major>_<version_minor>_pivot.yaml repository

```go
model:
name: vcs
major: 1
minor: 0
commandtypes:
eventtypes:
resourcetypes: 
- model: 
  name: repository
keys:
- model:
  name: update_mode    
  defaultValue: 
  type: string
  shortname:  
  mandatory: true
- model:
  name: update_time
  defaultValue: 
  type: string
  shortname:  
  mandatory: true 
  ```
  
  + Step 4: Creation of connector product into the gandalf-worker/configurations////<version_major>_<version_minor>_product_connector.yaml repository 

```go
model:
name: VCSGithub1.0
major: 1
minor: 0
product:
  name: github
commandtypes:
- model:
  name: CREATE_REPOSITORY
  schema: '{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/CreateRepositoryPayload","definitions":{"CreateRepositoryPayload":{"required":["Token","Name","Description","Private"],"properties":{"Token":{"type":"string"},"Name":{"type":"string"},"Description":{"type":"string"},"Private":{"type":"boolean"}},"additionalProperties":false,"type":"object"}}}'
- model:
  name: CREATE_REPOSITORY_FROM_TEMPLATE
  schema: '{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/CreateRepositoryFromTemplatePayload","definitions":{"CreateRepositoryFromTemplatePayload":{"required":["Token","TemplateOwner","TemplateRepo","Name","Owner","Description","Private"],"properties":{"Token":{"type":"string"},"TemplateOwner":{"type":"string"},"TemplateRepo":{"type":"string"},"Name":{"type":"string"},"Owner":{"type":"string"},"Description":{"type":"string"},"Private":{"type":"boolean"}},"additionalProperties":false,"type":"object"}}}'
- model:
  name: DELETE_REPOSITORY
  schema: '{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/DeleteRepositoryPayload","definitions":{"DeleteRepositoryPayload":{"required":["Token","Owner","Repository"],"properties":{"Token":{"type":"string"},"Owner":{"type":"string"},"Repository":{"type":"string"}},"additionalProperties":false,"type":"object"}}}'
- model:
  name: CREATE_ISSUE
  schema: '{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/CreateIssuePayload","definitions":{"CreateIssuePayload":{"required":["Token","Owner","Repository","Title","Body"],"properties":{"Token":{"type":"string"},"Owner":{"type":"string"},"Repository":{"type":"string"},"Title":{"type":"string"},"Body":{"type":"string"}},"additionalProperties":false,"type":"object"}}}'
- model:
  name: CREATE_PULL
  schema: '{"$schema":"http://json-schema.org/draft-04/schema#","$ref":"#/definitions/CreatePullPayload","definitions":{"CreatePullPayload":{"required":["Token","Owner","Repository","Title","Body","Head","Base"],"properties":{"Token":{"type":"string"},"Owner":{"type":"string"},"Repository":{"type":"string"},"Title":{"type":"string"},"Body":{"type":"string"},"Head":{"type":"string"},"Base":{"type":"string"}},"additionalProperties":false,"type":"object"}}}'
eventtypes:
- model:
  name: COMMIT
  schema: '{"type":"string"}'
- model:
  name: PULL
  schema: '{"type":"string"}'
resourcetypes: 
keys:
- model:
  name: username     
  defaultValue:     
  type: string
  shortname:     
  mandatory: true
- model:
  name: password     
  defaultValue:    
  type: string
  shortname:     
  mandatory: true
```


+ Step 5: Build the worker with the go build -o worker command
+ Step 6: Packaging the worker in a worker.zip archive
+ Step 7: Upload the archive to the gandalf-worker/workers///<version_major>/<version_minor>/worker.zip repository
+ Step 8: Update / Create the file gandalf-worker/workers///versions.yaml

## Example of creating a worker (Step 2) with the Github connector

+ Step 1: Initializing the worker base and recovering the router inputs
```go
	var major = int64(1)
	var minor = int64(0)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)
```
+ Step 2: Connect to the tool and save the connection in the context on worker base
```go
var inputPayload InputPayload
err := json.Unmarshal([]byte(input.Text()), &inputPayload)
if err == nil {
	if inputPayload.Token != "" {
		clientGithub := client.Oauth2Authentification(inputPayload.Token)
		worker.Context["client"] = clientGithub
	} else if inputPayload.Username != "" && inputPayload.Password != "" {
		clientGithub := client.BasicAuthentification(inputPayload.Username, inputPayload.Password)
		worker.Context["client"] = clientGithub
	}

```

> input payload

```go
type InputPayload struct {
	Username         string
	Password         string
	Token            string
	EventTypeToPolls []models.EventTypeToPoll
	//....
}
```
+ Step 3: Registering worker functions Register function:

```go
worker.RegisterCommandsFuncs("CREATE_REPOSITORY", CreateRepository)
```

CreateRepository Function

```go
func CreateRepository(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var createRepositoryPayload repository.CreateRepositoryPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &createRepositoryPayload)
	if err == nil {
		var clientGithub *github.Client
		if createRepositoryPayload.Token != "" {
			clientGithub = client.Oauth2Authentification(createRepositoryPayload.Token)
		} else {
			clientGithub = context["client"].(*github.Client)
		}

		err = repository.CreateRepository(clientGithub, createRepositoryPayload.Name, createRepositoryPayload.Description, createRepositoryPayload.Private)
		if err == nil {

			return 0
		}
	}
	return 1
}
```
> The signature of the function must always be the same

```go
type CreateRepositoryPayload struct {
	Token       string
	Name        string
	Description string
	Private     bool
}

func CreateRepository(client *github.Client, name, description string, private bool) (err error) {
	ctx := context.Background()
	r := &github.Repository{Name: &name, Private: &private, Description: &description}
	_, _, err = client.Repositories.Create(ctx, "", r)

	return
}
```
+ Step 5: Run the worker

```go
worker.Run()
```