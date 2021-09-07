ARG gandalf_uid
ARG gandalf_gid

FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image 
# GO111MODULE=on
ENV GO111MODULE=auto  \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
#COPY go.mod .
#COPY go.sum .
#RUN go mod download

# Copy the code into the container
#COPY . .
COPY go.mod .
COPY main.go .
COPY workflow.go .
COPY Dockerfile .

#
RUN go get


# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine

# Env vars
ENV env_identity=""
ENV env_timeout=""
ENV env_adresses=""

# Add Bash
RUN apk add --no-cache bash
#ENV env_identity=$identity
#ENV env_timeout=$timeout
#ENV env_adresses=$adresses

# Copy
COPY setup.sh .
COPY --from=builder /dist/main /

#
RUN bash /setup.sh "${gandalf_uid}" "${gandalf_gid}"

# Command to run
ENTRYPOINT "/main" $env_identity $env_timeout $env_adresses
#ENTRYPOINT touch /var/run/sockets/test.txt

#docker build --build-arg gandalf_uid=$(id -u gandalf) --build-arg gandalf_gid=$(id -g gandalf)
#docker run -e env_identity={identity} -e env_timeout={timeout} -e env_adresses={path_to_socket}/{socket} --volume "{path_to_socket}:/var/run/sockets/:rw" {id}