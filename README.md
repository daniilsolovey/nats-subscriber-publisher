# Nats-subscriber-publisher

<img alt="" src="https://i.imgur.com/qyGhLjz.gif"/>

## Run docker container with NATS:
```
docker run --name your_name_nats_container -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```

## Build app:
```
go build
```

## Run publisher:
```
go run publisher/publisher.go
```

## Run subscriber:
```
go run subscriber/subscriber.go
```