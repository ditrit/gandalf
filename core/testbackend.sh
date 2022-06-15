#!/bin/bash

cmd=$1
suffix=${2:-0}

if [ -z $cmd ]
then 
  echo "Usage : testbackend cmd [storage]"
  echo
  echo "Allow to easily test Gandalf as a backend for Leto using Docker"
  echo 
  echo "Commands:"
  echo "  build   build the gandalfdocker image tu be used"
  echo "  run     runs a gandalf system (cluster and aggregator) using the given storage"
  echo "  del     delete a storage"
  echo 
  echo "Options:"
  echo "  storage the name of the storage to use. If none, the name used is '0'" 
  echo
fi

volume_cockroach=cockroach_${suffix}
volume_gandalf=gandalf_${suffix}
if [ "$1" == "build" ]
then
  cat >Dockerfile <<EOF
FROM centos AS builder
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
WORKDIR /go/src/
RUN yum -y update
RUN yum -y install git golang wget
RUN git clone https://github.com/ditrit/gandalf.git
WORKDIR /go/src/gandalf
RUN git checkout api-development
WORKDIR /go/src/gandalf/core
RUN go get
RUN go build -o gandalf
RUN wget https://binaries.cockroachdb.com/cockroach-v20.1.6.linux-amd64.tgz
RUN tar -xf cockroach-v20.1.6.linux-amd64.tgz

FROM centos
WORKDIR /usr/local/bin/
COPY --from=builder /go/src/gandalf/core/gandalf .
COPY --from=builder /go/src/gandalf/core/cockroach-v20.1.6.linux-amd64/cockroach .
COPY --from=builder /go/src/gandalf/core/certs certs
COPY --from=builder /go/src/gandalf/core/rungandalf.sh .
RUN chmod 600 certs/node.key
RUN chmod 600 certs/client.root.key
RUN mkdir /etc/gandalf
RUN mkdir /var/lib/gandalf
RUN cp -r certs /etc/gandalf/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN yum -y update
RUN yum install tzdata ca-certificates
EXPOSE 9203
CMD rungandalf.sh
EOF
  sudo docker build . --no-cache -t gandalfdocker
  rm Dockerfile
fi
if [ "$1" == "run" ]
then
  sudo docker run -p 127.0.0.1:9203:9203/tcp -v ${volume_gandalf}:/var/lib/gandalf -v ${volume_cockroach}:/var/lib/cockroach gandalfdocker
fi
if [ "$1" == "del" ]
then
  sudo docker volume rm ${volume_gandalf} 
  sudo docker volume rm ${volume_cockroach}
fi


