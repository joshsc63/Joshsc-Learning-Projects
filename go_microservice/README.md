# Golang Microservices w/ Postgres, MongoDB
The building blocks to Golang Microservices that includes 

![ui](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/ui.png "ui")

![design](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/design.png "design")


- UI site to connect the front end to the golang backend
- A broker service to handle & route requests
- authentication service using PostGres for user data
- A logger service that sends messages to MongoDB
- Mail Service for emails
- Listener service for Rabbit MQ messages
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

### Kubernetes Cluster
Make will initialize the kubernetes cluster on the local cluster
See [https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/project/docker-compose.yml](docker-compose.yml) file for pods & services

## Go Packages

- [https://github.com/xhit/go-simple-mail](simple-go-mail) : send email
- [https://github.com/vanng822/go-premailer](go-premailer) : http styling email
- [https://github.com/go-chi/chi](go-chi) : go router for HTTP services
- [https://github.com/rabbitmq/amqp091-go](rabbitmq-go) : rabbit MQ client