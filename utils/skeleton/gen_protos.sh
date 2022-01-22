#!/bin/bash
my_dir=`dirname $0`


## install dev dependencies
#GO111MODULE=off go get -u -v github.com/gogo/protobuf/proto
#GO111MODULE=off go get -u -v github.com/gogo/protobuf/jsonpb
#GO111MODULE=off go get -u -v github.com/gogo/protobuf/protoc-gen-gogofaster
#GO111MODULE=off go get -u -v github.com/gogo/protobuf/gogoproto
#GO111MODULE=off go get -u -v google.golang.org/grpc


# you want to be in /some/directory/ to get the "correct import path for generated code
cd ../../../../ &&
# current working directory now in /some/directory/

## generate models
#protoc \
#    -I=$GOPATH/src                                    ` #include gogoproto package in $GOPATH/src                                ` \
#    -I=$(pwd)                                         ` #include  /some/directory/                                               ` \
#    --gogofaster_out=$(pwd)                           ` #output in /some/directory/                                                ` \
#    github.com/pinezapple/LibraryProject20201/skeleton/model/*.proto ` #input file in /some/directory/github.com/cicdata-io/smartcic-core/model `

## generate rpc
# protoc \
#     -I=$GOPATH/src                                   ` #include gogoproto package in $GOPATH/src                              ` \
#     -I=$(pwd)                                        ` #include  /some/directory/                                             ` \
#     --gogofast_out=plugins=grpc:$(pwd)             ` #output in /some/director                                              ` \
#     -I=$GOPATH/src/github.com/gogo/protobuf/protobuf ` #runs fine without this line                                           ` \
#     github.com/pinezapple/LibraryProject20201/skeleton/model/*.proto  ` #input file in /some/directory/github.com/cicdata-io/smartcic-core/rpc `

protoc \
    -I=$GOPATH/src \
    --gogofaster_out=plugins=grpc:. \
    -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
    github.com/pinezapple/LibraryProject20201/skeleton/model/docmanager.proto &&

# protoc \
#     -I=$GOPATH/src \
#     -I=$(pwd) \
#     --gogofaster_out=$(pwd) \
#     $GOPATH/src/github.com/pinezapple/LibraryProject20201/skeleton/model/*.proto

echo "DONE: GEN PROTO"