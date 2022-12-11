# Golang Microservices w/ Postgres, MongoDB, RabbitMQ

![ui](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/ui.png "ui")

![design](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/design.png "design")

The building blocks to Golang Microservices that includes 

- UI site to connect the front end to the golang backend
- A broker service to handle & route requests
- authentication service using PostGres for user data
- A logger service that sends messages to MongoDB
- Logger & Broker sends messages via RPC to each other
- Mail Service for emails
- Listener service for Rabbit MQ messages to send to MongoDB
- Rabbit MQ for messages
- Postgres DB for user data
- MongoDB for logging messages

Services will deploy on a docker container on the local kubernetes cluster. 

## To use

### Build & Run on K8S
Makefile performs go build of binary files & initializes docker containers to a kubernetes cluster

1. cd project dir
2. Run Make `make up_build` for backend services
3. Run Make `make start` to start front-end

if localhost port is in use... can verify & end process `lsof -i tcp:80`

To stop the cluster `make down`

### Troubleshooting

If encountering error when stopping service `failed to remove network project_default: Error response from daemon: error while removing network: network project_default id ac674b25216b099c9fc70f9d1d46886f73a8f03f6a0e75bf97c01d7dc4f4155a has active endpoints`

Run cmd `docker-compose down --remove-orphans`


### Navigate to UI
Run `make start` 

Navigate to `localhost` on browser

to stop: `make stop`

## Components

### Logger Service
Send log event via JSON / RPC / gRPC to MongoDB

`broker-service` handler file has switch statement for events. Log contains method calls for different messaging type
```
	case "log":
		//app.logItem(w, requestPayload.Log)        // log to mongoDB
		//app.logEventRabbit(w, requestPayload.Log) // log via rabbitmq
		//app.logItemRPC(w, requestPayload.Log)     // log via RPC
		app.logItemgRPC(w, requestPayload.Log)      // log via gRPC
```

#### gRPC
gRPC Reqs:
- `brew install protobuf`
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

`logger-service/logs/logs.proto` proto3 file for gRPC

- protoc auto-gen pg.go file created via `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto`

### Mail Service
Uses MailHog to simulate a localhost mail box vs a valid web email. See `docker-compose` - Navigate to `localhost:8025`

![mailhog](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/mailhog.png "mailhog")

### PG Database
DB runs locally. See dir `/project/db-data` for PG generated files
- Credentials hardcoded into docker-compose

Connect to `localhost`

### Mongo DB
![mongo](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/mongo.png "mongo")

Hosts log events in the `logger-service`

- Client: MongoDB Compass to connect
- URI: `mongodb://admin:password@localhost:27017/log?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false`

### Rabbit MQ
Log events are sent to RabbitMQ which then creates an entry to MondoDB from its message

See listener-service

### Kubernetes Cluster
Make will initialize the kubernetes cluster on the local cluster
See [https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/project/docker-compose.yml](docker-compose.yml) file for pods & services

## Go Packages

- [https://github.com/xhit/go-simple-mail](simple-go-mail) : send email
- [https://github.com/vanng822/go-premailer](go-premailer) : http styling email
- [https://github.com/go-chi/chi](go-chi) : go router for HTTP services
- [https://github.com/rabbitmq/amqp091-go](rabbitmq-go) : rabbit MQ client
- [https://github.com/grpc/grpc-go](gRPC-go) : gRPC messages 
- [https://github.com/golang/protobuf](gRPC-go) : Protocol Buffer

