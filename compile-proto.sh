#!/bin/bash

declare -a protos=(
    'collection.proto'
    'lesson.proto'
    'livelesson.proto'
    'curriculum.proto'
    'antidoteinfo.proto'
    'image.proto'
    'livesession.proto'
);

echo -n "Compiling protobufs: ("
for i in "${protos[@]}"
do
    echo -n "$i"
    protoc -I api/exp/definitions/ -I./api/exp/definitions \
        -I api/exp/definitions/ \
        api/exp/definitions/"$i" \
            -I$GOPATH/src/github.com/nre-learning/antidote-core/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:api/exp/generated/ \
        --grpc-gateway_out=logtostderr=true,allow_delete_body=true:api/exp/generated/ \
        --swagger_out=logtostderr=true,allow_delete_body=true:api/exp/definitions/
    echo -n ","
done

echo ")"
