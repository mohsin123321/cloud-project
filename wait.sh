# Wait for MongoDB to be ready
./wait.sh mongodb:27017 -t 30

# Start the server
./server --port 8080 --db-type mongodb --db-addr mongodb --db-port 27017--db-user root --db-pass root
