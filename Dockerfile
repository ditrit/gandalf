# Setup build stage for Gandalf
FROM golang:1.18-bullseye as setup-build-stage-gandalf
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y git
WORKDIR /app
COPY . .

# Build stage for Gandalf
FROM setup-build-stage-gandalf as build-stage-gandalf
WORKDIR /app/core
RUN go get
RUN go build -o gandalf

# Build stage for CockRoach
FROM ubuntu:21.04 as build-stage-coackroach
ENV COCKROACH_VERSION=v20.1.6
ENV COCKROACH_PKG=cockroach-$COCKROACH_VERSION.linux-amd64
ENV COCKROACH_TGZ=$COCKROACH_PKG.tgz
WORKDIR /app
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y wget
RUN wget https://binaries.cockroachdb.com/$COCKROACH_TGZ
RUN tar xf ./$COCKROACH_TGZ -C ./
RUN mv ./$COCKROACH_PKG/cockroach ./

## Setup stage for production
FROM ubuntu:21.04 as setup-stage-production
COPY --from=build-stage-gandalf /app/core/gandalf usr/bin/
COPY --from=build-stage-gandalf /app/core/certs /etc/gandalf/certs
COPY --from=build-stage-coackroach /app/cockroach /usr/local/bin/
WORKDIR /app
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y apt-utils
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates tzdata
RUN ln -fs /usr/share/zoneinfo/Europe/Paris /etc/localtime
RUN mkdir -p /var/lib/cockroach /var/lib/gandalf /var/log/gandalf
RUN chmod 600 /etc/gandalf/certs/node.key /etc/gandalf/certs/client.root.key

# Production stage
FROM setup-stage-production as stage-production
WORKDIR /etc/gandalf
