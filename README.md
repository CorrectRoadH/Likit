<p align="center">
    <img src="https://github.com/CorrectRoadH/Likit/blob/main/img/logo.png?raw=true" height="128"/></a>
</p>

# Likit

[Demo](https://likit.zeabur.app)

## What is Likit
Likit is a backend as a service for like(vote, unlike, count) feature. It is very ease to deploy and use.

[Getting started](./docs/getting-started.md)

# Screenshots

![](./img/screenshot-1.png)
![](./img/screenshot-2.png)

# Deployment

## Zeabur
Deloyment on Zeabur by one click

[![Deploy on Zeabur](https://zeabur.com/button.svg)](https://zeabur.com/templates/KZOLHA?referralCode=CorrectRoadH)

## Docker Compose
```
wget https://raw.githubusercontent.com/CorrectRoadH/Likit/main/docker-compose.yaml

docker compose up -d
```

## K8S
Coming soon

# Documentation
[Document](./docs/getting-started.md)

## RESTful API

### Vote
POST `http(s)://<your likit ip>/api/v1/vote`

body

```
{
    "businessId": "businessId",
    "messageId": "messageId",
    "userId": "userId"
}
```

POST `http(s)://<your likit ip>/api/v1/unvote`

body

### Unvote
```
{
    "businessId": "businessId",
    "messageId": "messageId",
    "userId": "userId"
}
```

### Count the number of votes
GET `http(s)://<your likit ip>/api/v1/count/:businessId/:messageId`

### List the users who voted
GET `http(s)://<your likit ip>/api/v1/list/:businessId/:messageId`

### Check if the user has voted
GET `http(s)://<your likit ip>/api/v1/isVoted/:businessId/:messageId/:userId`


## gRPC API
Coming soon

## SDK

### Golang
[Likit Go](https://github.com/CorrectRoadH/likit-go)


### Java
Coming soon
