# Golang Microservices w/ Postgres, MongoDB
The building blocks to Golang Microservices that includes 

1. UI site to connect the front end to the golang backend
2. authentication service
3. A broker service to handle & route requests
4. A logger service 
5. Postgres DB for user data
6. MongoDB for logging messages

Services will deploy on a docker container on the local kubernetes cluster. 

## To use

### Build & Run on K8S
Makefile performs go build of binary files & initializes docker containers to a kubernetes cluster

1. cd project dir
2. Run Make `make up_build` for backend services
3. Run Make `make start` to start front-end

if localhost port is in use... can verify & end process `lsof -i tcp:80`

To stop the cluster `make down`

### Navigate to UI
Navigate to `localhost` on browser

## Components

### PG Database
DB runs locally. See dir `/project/db-data` for PG generated files
- Credentials hardcoded into docker-compose

Connect to `localhost`

### Mongo DB
Hosts log events in the `logger-service`

- Client: MongoDB Compass to connect
- URI: `mongodb://admin:password@localhost:27017/log?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false`


### Kubernetes Cluster
Make will initialize the kubernetes cluster on the local cluster
See [https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/project/docker-compose.yml](docker-compose.yml) file for pods & services
