# Golang Microservices w/ Postgres, MongoDB, RabbitMQ

![ui](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/ui.png "ui")

![design](https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/images/design.png "design")

The building blocks to Golang Microservices that includes 

- UI site to connect the front end to the golang backend using caddy as web server
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

### Front End Web Server
Uses caddy docker image as webserver
File: `/project/Caddyfile`

Modify `/etc/hosts` file to add backend route
```
  1 ##
  2 # Host Database
  3 #
  4 # localhost is used to configure the loopback interface
  5 # when the system is booting.  Do not change this entry.
  6 ##
  7 127.0.0.1   localhost backend
  8 255.255.255.255 broadcasthost
  9 ::1             localhost backend
```

### Web Hosting w/ Linode
Docker swarm instances are hosted using Linode

- Ubuntu 20.2
- Labels: `node-1` & `node-2` (two nodes)
- SSH Key: pubkey

#### Linode Setup
- ssh into node-1/2 
- Add user `adduser joshsc63` -> enter PW -> blank default vals
- `usermod -aG sudo joshsc63`
- Setup ubuntu basic firewall `ufw allow ssh | ufw allow http | ufw allow https`
- Open ports for services `ufw allow 2377/tcp | ufw allow 7946/tcp | ufw allow 7946/udp | ufw allow 4789/udp | ufw allow 8025/tcp` -> enable `ufw enable` -> status `ufw status`
- login as user `ssh joshsc63@45.79.XXX.X`

* Docker Install to Linode Node
- `sudo apt-get update`
```
	sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
```
```
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
```
```
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
- `sudo apt-get update` -> `sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin`

* Setting Hostname
- `sudo hostnamectl set-hostname node-1`
- Add entries to hostfile `sudo vi /etc/hosts`
```
45.79.XXX.X     node-1
45.79.XXX.XXX   node-2
```

* Adding DNS entry
Using Go-Daddy for domain name

* In DNS Records
- Add `Type A` | `Name node-1` | `Value IP` for both nodes 
- Add `Type A` | `Name swarm` | `Value IP` for both nodes
- Add `TYPE C` | `Name broker` fetch reqs | `Value swarm.DOMAIN.com`

- test record `ping swarm.DOMAIN.com`

* Initialize Docker Swarm to Linode Nodes
- `sudo docker swarm init --advertise-addr $IP-node-1` node1
- Run `docker swarm join --token XXX` command returned on Node2 to have it as worker node




### Logger Service
Send log event via JSON / RPC / gRPC to MongoDB

`broker-service` handler file has switch statement for events. Log contains method calls for different messaging type
```
	case "log":
		//app.logItem(w, requestPayload.Log)        // log to mongoDB
		//app.logEventRabbit(w, requestPayload.Log) // log via rabbitmq
		app.logItemRPC(w, requestPayload.Log)     // log via RPC
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

### Docker

#### Tag/Push Image to Docker Hub

- `docker login`
- `docker build -f logger-service.dockerfile -t joshsc63/logger-service:1.0.0 .`
- `docker push joshsc63/logger-service:1.0.0`

#### Docker Swarm
Used to host containers for a light weight option vs K8S. Docker Storm orchestrates container node's instances

File: `/project/swarm.yml`

CMDs
- `docker swarm init` init 1 worker
- `docker stack deploy -c swarm.yml myapp` deploy to docker swarm
- `docker service ls` show services running
- `docker service scale myapp_listener-service=3` scale to # of instances
- `docker stack rm myapp` remove swarm
- `docker swarm leave --force` leave swarm
 
 Updating Service
 - Build new tagged version of image `docker build -f logger-service.dockerfile -t joshsc63/logger-service:1.0.1 .` -> `docker push joshsc63/logger-service:1.0.1`
 - Update version in docker swarm `docker service update --image joshsc63/logger-service:1.0.1 myapp_logger-service`
 

#### Local Docker Cluster
`make up` will initialize the local cluster 
See [https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/project/docker-compose.yml](docker-compose.yml) file for containers & services

## Go Packages

- [https://github.com/xhit/go-simple-mail](simple-go-mail) : send email
- [https://github.com/vanng822/go-premailer](go-premailer) : http styling email
- [https://github.com/go-chi/chi](go-chi) : go router for HTTP services
- [https://github.com/rabbitmq/amqp091-go](rabbitmq-go) : rabbit MQ client
- [https://github.com/grpc/grpc-go](gRPC-go) : gRPC messages 
- [https://github.com/golang/protobuf](gRPC-go) : Protocol Buffer

