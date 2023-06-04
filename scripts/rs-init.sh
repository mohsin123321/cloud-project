#!/bin/bash

DELAY=20

# configuration for the replica set
mongo <<EOF
var config = {
    "_id": "db-replication",
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "mongo1:27017",
            "priority": 3
        },
        {
            "_id": 1,
            "host": "mongo2:27017",
            "priority": 2
        },
        {
            "_id": 2,
            "host": "mongo3:27017",
            "priority": 1
        }
    ]
};
rs.initiate(config, { force: true });
rs.status();
EOF

echo "****** Waiting for ${DELAY} seconds for replicaset configuration to be applied ******"

sleep $DELAY

# create user if it does'nt exist
mongo <<EOF
rs.status();
use testdb;
if (!db.getUser('test')) {
    db.createUser({ user: 'test', pwd: 'test', roles: [ { role: 'dbOwner', db: 'testdb' } ]});
}
EOF
