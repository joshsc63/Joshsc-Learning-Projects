# Substituted from Makefile logic. Here for reference
# FROM golang:1.18-alpine as builder
# 
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# 
# # CGO_ENBALED - just use standard library, not C
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
# 
# RUN chmod +x /app/brokerApp

# build tiny docker image
FROM alpine:latest

RUN mkdir /app

# Replaced by referencing app made from makefile
#COPY --from=builder /app/brokerApp /app
#CMD [ "/app/brokerApp" ]

COPY brokerApp /app

CMD [ "/app/brokerApp"]