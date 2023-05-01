jq -n \
    --arg dbtype "$DB_TYPE" \
    --arg dbaddress "$DB_ADDRESS" \
    --arg dbport "$DB_PORT" \
    --arg dbuser "$DB_USER" \
    --arg dbpass "$DB_PASS" \
    --arg serverport "$SERVER_PORT" \
    --arg tokensecret "$TOKEN_SECRET" \
    '{
        "ShowDocs": true,
        "Database": {
            "DbType": $dbtype,
            "DbAddr": $dbaddress,
            "DbPort": $dbport,
            "DbUser": $dbuser,
            "DbPass": $dbpass
        },
        "Server": {
            "Port": "8080"
        },
        "Token": {
            "Secret": $tokensecret
        }
    }'  > config.json
