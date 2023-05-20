#!/bin/bash

jq -n \
    --arg dbtype "$DB_TYPE" \
    --arg dbuser "$DB_USER" \
    --arg dbpass "$DB_PASS" \
    --arg dbname "$DB_NAME" \
    --arg serverport "$SERVER_PORT" \
    --arg tokensecret "$TOKEN_SECRET" \
    --arg replica1host "$REPLICA1_HOST" \
    --arg replica1port "$REPLICA1_PORT" \
    --arg replica2host "$REPLICA2_HOST" \
    --arg replica2port "$REPLICA2_PORT" \
    --arg replica3host "$REPLICA3_HOST" \
    --arg replica3port "$REPLICA3_PORT" \
    --arg replicaname "$REPLICA_NAME" \
    '{
        "ShowDocs": true,
        "Database": {
            "DbType": $dbtype,
            "DbUser": $dbuser,
            "DbPass": $dbpass,
            "DbName": $dbname,
            "ReplicaName": $replicaname,
            "Replicas" : [
                {
                    "Host": $replica1host,
                    "Port": $replica1port 
                },
                {
                    "Host": $replica2host,
                    "Port": $replica2port
                },
                {
                    "Host": $replica3host,
                    "Port": $replica3port 
                }
            ]
        },
        "Server": {
            "Port": $serverport
        },
        "Token": {
            "Secret": $tokensecret
        },
        "RateLimiter": {
            "SecondsWindow": 60,
            "MaxReqPerIP": 1000
        }
    }'  > config.json
