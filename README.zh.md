<p align="center">
    <img src="https://github.com/CorrectRoadH/Likit/blob/main/img/logo.png?raw=true" height="128"/></a>
</p>

# Likit

[Demo](https://likit.zeabur.app)

[English](./README.md) | 简体中文

## What is Likit
Likit 是一个 Backend as a Service ，提供点赞（投票、点赞、计数）功能。 部署和使用非常容易。

Likit的目标是帮助您在10分钟内实现类似的功能。

Likit 适用于中小型应用程序。 您可以选择差异投票工具引擎来获得不同的功能和性能。

[Getting started](./docs/getting-started.md)

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

# Roadmap
- [ ] add Rank with limit and Rank from Message Id to Simple Vote Engine
- [ ] Vote to Queue. Async vote
- [ ] add Vote events record to Middle Vote Engine
- [ ] add output message queue to Middle Vote Engine
- [ ] add tag to message
- [ ] Rank by tag

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