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
2. Run Make `make up_build`

To stop the cluster `make down`

### Navigate to UI
Navigate to `localhost` on browser

### Kubernetes Cluster
Make will initialize the kubernetes cluster on the local cluster
See [https://github.com/joshsc63/Joshsc-Learning-Projects/blob/main/go_microservice/project/docker-compose.yml](docker-compose.yml) file for pods & services
