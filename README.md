# Automation Godog

Este projeto é uma automação de testes de aceitação utilizando BDD com Godog e Golang.

## Estrutura do Projeto

```
automation-godog/
├── bin/
├── cmd/
│   └── automation-godog/
│       └── main.go
├── configs/
│   └── config.yaml
├── docs/
│   └── README.md
├── features/
│   ├── user_authentication.feature
│   ├── data_processing.feature
├── internal/
│   ├── aws/
│   │   ├── s3.go
│   │   ├── rds.go
│   │   ├── sqs.go
│   │   └── kafka.go
│   ├── steps/
│   │   ├── user_authentication_steps.go
│   │   ├── data_processing_steps.go
│   └── utils/
│       ├── http_client.go
│       └── config_loader.go
├── scripts/
│   ├── build.sh
│   ├── deploy.sh
├── testdata/
│   └── mock_responses/
│       ├── user_authentication_response.json
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```


## Execução

Para rodar os testes, utilize o comando:
```
go test ./...
```
