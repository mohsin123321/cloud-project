#!/bin/bash

jq -n \
    --arg dbtype "$DB_TYPE" \
    --arg dbaddress "$DB_ADDRESS" \
    --arg dbport "$DB_PORT" \
    --arg dbuser "$DB_USER" \
    --arg dbpass "$DB_PASS" \
    --arg dbname "$DB_NAME" \
    --arg serverport "$SERVER_PORT" \
    --arg tokensecret "$TOKEN_SECRET" \
    '{
        "ShowDocs": true,
        "Database": {
            "DbType": $dbtype,
            "DbAddr": $dbaddress,
            "DbPort": $dbport,
            "DbUser": $dbuser,
            "DbPass": $dbpass,
            "DbName": $dbname
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
