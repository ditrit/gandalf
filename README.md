# Gandalf Project

Gandalf (Gandalf is A Natural Devops Application Life-cycle Framework), is a tool to allow progressive DevOps adoption.

Gandalf allows each of your teams, at any time and in a few minutes, to integrate and evolve its tools and workflows.

https://ditrit.io/

## Core 
The Gandalf cluster, aggregator, connector infrastructure: [Readme](core/README.md)

## Connectors
Interfaces between Gandalf's pivot languages and exisiting APIs: [Readme](connectors/README.md)

## Libraries
Tooling to write connectors: [Readme](libraries/README.md)

## Verdeter
Tooling to write configuration: [Readme](verdeter/README.md)



## Docker

This section show how to start use gandalf with docker

### Build gandalf on docker

Run:

```bash
docker build . -t gandalf
```

Option to build:

- `COCKROACH_VERSION`, to select the version of cockroach to use, default is `v20.1.6`.

### How to use the image

The image of gandalf is not build with a default `CMD`, to let you choose how you have to start gandalf.

So, in the `docker-compose.yml`, you can see, that we use a cluster and an aggregator by providing the `CMD` in each services.

### How to test gandalf

You can find a `docker-compose.yml` that give you the ability to test a simple gandalf networking.

You should obtain a *"hello world"* response in your browser using the address *"http://127.0.0.1:9199/ditrit/Gandalf/1.0.0/"*.

