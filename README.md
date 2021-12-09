# Steedu Main Server

<img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<img src="https://img.shields.io/badge/postgres-14.1-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="postgres">&nbsp;<img src="https://img.shields.io/badge/License-GPLv3-blue?style=for-the-badge&logo=none" alt="license" />

# ⚡️ About The Project

스터디 관리 Web Application을 위한 API Server를 구현합니다. 
[fiber-go-template](https://github.com/create-go-app/fiber-go-template) 템플릿을 사용합니다.

# 📦 Built With
- go 1.17+
- [Fiber v2](https://github.com/gofiber/fiber)
- postgresql 14.1
- [ent](https://github.com/ent/ent)

# Getting Start

## Prerequisites

- [Docker](https://www.docker.com/get-started)

- Go tools
  - [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
  - [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
  - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
  - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

## Run Server
```
$ make compose.up
```

## Structure
```
├── app
│   ├── controllers
│   └── models
├── pkg
│   ├── configs
│   ├── middleware
│   ├── repository
│   ├── routes
│   └── utils
├── platform
│   ├── database
│   └── ent
├── rest.http
└── main.go
```
### ./app
비지니스 로직만을 다루는 폴더 입니다. 
어떤 데이터베이스 드라이버를 쓰는지, cache 전략을 갖고 있는지 혹은 third-party 등을 다루지 않습니다.
- `./app/controllers` router에서 사용되는 함수 단위의 contoller를 정의합니다. Fiber의 Handler에 해당합니다.
- `./app/models` 비지니스 모델 및 메소드를 정의 합니다. request 및, response에 해당합니다.

### ./docs
API 문서를 포함합니다. swager가 자동으로 생성합니다.

### ./pkg

프로젝트에 특정지어지는 기능들을 포함합니다. configs, middleware, routes, utils와 같이 비지니스에 사용되는 코드들이 해당합니다.

- `./pkg/configs` configuration가 정의됩니다. i.e. readTimeout
- `./pkg/middleware` middleware에 해당합니다. i.e. logger, cors
- `./pkg/repository` const 값이 정의됩니다. i.e. enum values
- `./pkg/routes` handler들을 routing합니다.
- `./pkg/utils`  utility를 포합합니다. i.e. server starter, error checker, etc


### .platform
platform level의 로직을 포함합니다. 실제로 프로젝트를 서비스하기 위해 필요한 기반 로직을 다룹니다. database나 cache등이 있습니다.

- `./platform/database` database를 setup합니다. ent와 PostgreSQL가 사용됩니다.
- `./platform/database` database schema를 정의하고 query를 수행하는 코드를 포합합니다.

## Roadmap
- [ ] Deployment
- [ ] Cover(Test Code)
- [ ] gosec
- [ ] goreport
- [ ] Add Change Log


## ⚠️  License
Distributed under the GPL License. See `LICENSE` for more information.