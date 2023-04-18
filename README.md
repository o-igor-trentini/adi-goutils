# GOUTILS - Atlassian Development Indicators #

- Repo principal: https://github.com/o-igor-trentini/atlassian-development-indicators

### Testes ###

Gerar relatório geral de cobertura de testes:
```bash
go test -coverprofile cover.out
```

Abrir relatório de cobertura de testes em uma página HTML:
```bash
go tool cover -html=cover.out
```
