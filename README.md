<p align="center">
    <a href="https://github.com/CorrectRoadH/Likit">
        <img src="https://github.com/CorrectRoadH/Likit/blob/main/img/logo.png?raw=true" height="128"/>
    </a>
</p>

# Likit

[Demo](https://likit.zeabur.app)

English | [简体中文](./README.zh.md)

## What is Likit
Likit is a Backend as a Service for like(vote, unlike, count) feature. It is very ease to deploy and use.

The goal of Likit is help your implement like function within 10 minutes.

Likit is suitable for small and middle size application. You can select difference vote implement engine to get difference features and performance.

[Getting Started](./docs/getting-started.md)

| features | Simple Vote Engine | Middle Vote Engine(WIP) |
| -- | -- | -- |
| Vote | ✅ |  🚧 |
| Unvote | ✅ |  🚧 |
| Count | ✅ | 🚧 |
| List the voted users | ✅ | 🚧 | 
| Is user Voted | ✅ | 🚧 | 
| Vote events | ❌ | 🚧 |
| Message Queue | ❌ | ❌ | 
| Rank | 🚧 | 🚧 |
| Requires | redis |  redis,postgres |

# Screenshots

![](./img/screenshot-1.png)
![](./img/screenshot-2.png)

# Deployment

## Zeabur
Deloyment on Zeabur by one click

[![Deploy on Zeabur](https://zeabur.com/button.svg)](https://zeabur.com/templates/KZOLHA?referralCode=CorrectRoadH)

Note: In zeabur the gRPC port is 443.

## Docker Compose
```
wget https://raw.githubusercontent.com/CorrectRoadH/Likit/main/docker-compose.yaml

docker compose up -d
```
and visit `http://localhost:7789`
grpc host is `localhost:4778`

## K8S
Coming soon

# Documentation
[Document](./docs/getting-started.md)

## Usage

### gRPC API
[gRPC SDKs in Buf](https://buf.build/likit/likit/sdks/main)

### Golang
[Likit Go](https://github.com/CorrectRoadH/likit-go)

### Java
[Likit Java](https://github.com/LxiHaaa/Likit-client)

### RESTful API
OpenAPI will be supported in the future

[RESTful API](./docs/restful.md)

# Who is using Likit

<a href="https://github.com/Get-Tech-Stack/TechStack">
    <img src="https://avatars.githubusercontent.com/u/141936114?s=48&v=4" height="64"/>
    <div>TechStack</div>
</a>