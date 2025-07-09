FROM --platform=linux ubuntu:22.04 as base
ARG BUILDARCH

# Change your versions here
ENV GO_VERSION=1.21.1
ENV IGNITE_VERSION=28.0.0
ENV NODE_VERSION=18.x
ENV MOCKGEN_VERSION=1.6.0
ENV PROTOC_VERSION=21.7

FROM base AS platform-amd64
ENV PROTOC_PLATFORM=x86_64

FROM base AS platform-arm64
ENV PROTOC_PLATFORM=aarch_64

FROM platform-${BUILDARCH}

ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

ENV PACKAGES curl gcc jq make unzip
RUN apt-get update
RUN apt-get install -y $PACKAGES

# Install Go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-$BUILDARCH.tar.gz | tar -C $LOCAL -xzf -

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

# Install Node
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs

# Install Mockgen
RUN go install github.com/golang/mock/mockgen@v${MOCKGEN_VERSION}

# Install ProtoC
RUN mkdir -p /usr/lib/protoc
WORKDIR /usr/lib/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${PROTOC_PLATFORM}.zip -o protoc.zip
RUN unzip -o protoc.zip
RUN rm protoc.zip
RUN ln -s /usr/lib/protoc/bin/protoc ${LOCAL}/bin/protoc

RUN apt-get clean

EXPOSE 1317 4500 5000 26657 5173 9090 9091 26656 6060

WORKDIR /cifer

COPY go.mod /cifer/go.mod
RUN go mod download
RUN rm /cifer/go.mod

# Create the image
# $ docker build -f Dockerfile . -t cifer_i
# To test only 1 command
# $ docker run --rm -it -v $(pwd):/cifer -w /cifer cifer_i go test github.com/b9lab/cifer/x/cifer/keeper
# To build container
# $ docker create --name cifer -i -v $(pwd):/cifer -w /cifer -p 6060:6060 -p 26656:26656 -p 9091:9091  -p 9090:9090 -p 26659:26659 -p 443:443 -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 -p 5173:5173 cifer_i
# $ docker start cifer
# To run server on it
# $ docker exec -it cifer ignite chain serve --reset-once
# In other shell, to query it
# $ docker exec -it cifer npm run dev -- --host
# cosmos1l950395d9s8fytxkgjsthv9h3sp9pzjscm8a66
# cosmos1l950395d9s8fytxkgjsthv9h3sp9pzjscm8a66
# cosmos1l950395d9s8fytxkgjsthv9h3sp9pzjscm8a66
# npm run dev -- --host
# http://0.0.0.0:26659
#  wscat ws://127.0.0.1:26657/websocket

