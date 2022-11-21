# Running the JSON validation server

A `docker-compose.yml` file is provided for ease of setting up the required database, therefore Docker is required. (The user could also set up their own MongoDB instance but a database would need to be initialized according to the parameters in `database/mongo-init.js`)

To start the database, from the base directory:
`docker compose up`

To stop the database and delete the container:
`docker compose down`

To build the Go application:
`go build cmd/main.go`

Or to build and run:
`go run cmd/main.go`

Please note that the server is running at `localhost:4060` and http requests will need to be directed there.
