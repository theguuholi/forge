# Forge

Projeto de estudo em Go para construir uma plataforma de publicação e administração de sites estáticos de clientes.

## Estado atual

A CLI exibe a versão e os comandos disponíveis. A interpretação dos argumentos `version` e `help` ainda não foi implementada: a saída atual apresenta as duas informações juntas.

Upload de sites, painel web com Go + HTMX + Tailwind e publicação são objetivos futuros.

## Pré-requisitos

- Go instalado. O `go.mod` declara Go 1.18; a verificação deste estágio foi feita com Go 1.27.4.
- Git para controle de versão.

Execute os comandos abaixo a partir da raiz do repositório.

## Executar durante o desenvolvimento

```bash
go run ./cmd/forge
```

Saída atual:

```text
Forge version: 0.0.1
Commandos disponiveis: version, help
```

## Testar

Executar todos os testes:

```bash
go test ./...
```

Exibir cada cenário e executar sem reutilizar resultados em cache:

```bash
go test -count=1 -v ./...
```

Executar apenas os testes da CLI:

```bash
go test -v ./internal/cli
```

Executar um subteste específico:

```bash
go test -v ./internal/cli -run '^TestRun$/^includes_version$'
```

Verificar acessos concorrentes durante os testes:

```bash
go test -race ./...
```

O race detector requer um ambiente compatível e, normalmente, um compilador C disponível.

## Cobertura

```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Abra `coverage.html` no navegador para inspecionar o relatório. Cobertura mede instruções executadas; testes também precisam verificar resultados e efeitos esperados.

## Formatar e analisar

```bash
go fmt ./...
go vet ./...
```

`go fmt` pode alterar a formatação dos arquivos. `go vet` analisa problemas potenciais e não substitui os testes.

## Compilar e executar

```bash
go build -o forge ./cmd/forge
./forge
```

O binário é gerado para o sistema operacional e arquitetura do ambiente de build por padrão.

## Verificação antes de concluir uma mudança

```bash
go fmt ./...
go vet ./...
go test -count=1 ./...
go build -o forge ./cmd/forge
```

Para alterações de comportamento, escrever primeiro um teste que falha pelo motivo esperado, implementar e verificar novamente. Acrescentar testes de integração quando houver dependências externas.

## Estrutura

```text
cmd/forge/main.go           Entrada do executável
internal/cli/cli.go         Composição da saída e versão
internal/cli/help.go        Texto de ajuda
internal/cli/cli_test.go    Testes do comportamento público da CLI
go.mod                     Identidade do módulo e versão de Go
```

Os testes protegem comportamentos. Não é necessário criar um arquivo de teste para cada arquivo de implementação quando o comportamento já está verificado.
