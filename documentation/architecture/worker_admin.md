# Gandalf - Admin Worker

The worker admin allows the management of other workers.

## Communication

The worker admin communicates in gRPC with the router thanks to the gandalf client

## Initialization

The worker admin initializes in several steps.

## Configuration

The worker admin recovers its pivot language.
Starting the workers

The worker admin creates the workers according to the versions:

+ If no version, then it will look for the latest version of the connector
+ Otherwise recuperation of the requested versions.

The start-up of a worker is carried out in four steps:

**Retrieving the worker configuration**

+ The connector sends a command to the cluster to retrieve its stored database configuration as well as all the pivot languages used by the connectors.
+ If the connector configuration does not exist in the database, it retrieves its configuration and the configuration keys from the repository.
    + The worker configuration must be present on {url}/{connecterType}/{product}/{versionMajor}_configuration.yaml
        Keys for the connector type must be present on {url}/{TypeConnector}/keys.yaml
    + Product keys must be present on {url}/{Type connector}/{product}/keys.yaml
        Keys for the major version must be present on {url}/{Typeconnector}/{product}/{versionMajor}/keys.yaml
    + Keys for the minor version must be present on {url}/{Typeconnector}/{product}/{versionMajor}/{versionMinor}/keys.yaml

**Recovery of workers**

The connector recovers workers through a url (this is the same url as for retrieving keys): {url}/{connecterType}/{product}/{versionMajor}_{versionMinor}_worker.zip.
**Executions of workers**

    The connector starts the workers in the given directory and sends the values of the configuration keys via the standard input.
    Starting a worker causes a command to be sent to the connector, which contains the list of commands that the worker implements himself.

**Validation of the configuration**

The connector validates its configuration against the expected one, and starts if the configuration is valid. Otherwise, the connector stops.

## Auto Update

The worker admin implements the update according to the specified option:

+ If auto, the worker admin will look to update all minutes.
+ If manual, the worker admin will wait for update commands.
+ If planned, the worker admin will update at the specified time.

## Setting up of orders

The worker admin saves these commands.

### ADMIN_GET_WORKER

Allows you to recover a version of a worker

### ADMIN_START_WORKER

Allows you to start a version of a worker

### ADMIN_STOP_WORKER

Allows you to stop a version of a worker

### ADMIN_GET_LAST_VERSION_WORKER

Allows you to recover the latest version of a worker

### ADMIN_UPDATE

Allows a worker to be updated

## Rolling Update

The worker admin can use an ADMIN_UPDATE command to update a worker. This update takes place in several steps:

+ Starting the new worker version
+ If the worker starts correctly, the worker of the previous version is put on stand by, that is to say that he no longer accepts new Command or Event but continues to execute the current Commands.
+ At the end of the execution of the current Commands, the worker of the previous version stops.
