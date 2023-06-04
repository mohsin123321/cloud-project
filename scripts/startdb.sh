#!/bin/bash

DELAY=15

# spin up the mongo db replica set
docker-compose --file docker-compose.yml up -d mongo1 mongo2 mongo3

echo "****** Waiting for ${DELAY} seconds for containers to go up ******"
sleep $DELAY

# give execute permissions
chmod +x ./scripts/rs-init.sh

docker exec mongo1 /scripts/rs-init.sh

# build and spin up the go app service
docker-compose --file docker-compose.yml up -d --build go-app