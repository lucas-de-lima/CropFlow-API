# Aula Técnica — Projeto Cropflow API

## Objetivo

Apresentar a arquitetura, principais componentes e fluxos da API `cropflow-api`, com exemplos e exercícios práticos.

## Público-alvo

Desenvolvedores backend com conhecimentos básicos de Go e arquitetura de serviços.

## Pré-requisitos

- Go 1.20+
- MySQL (ou Docker)
- Conhecimentos básicos de REST APIs e JWT

## Sumário

1. Visão Geral
2. Arquitetura do Projeto
3. Estrutura de Pastas
4. Configuração e Execução Local
5. Pontos-chave do Código
6. Fluxos da API
7. Trechos de Código
8. Exemplos (cURL)
9. Exercícios Práticos
10. Referências

---

## 1. Visão Geral

O `cropflow-api` é uma API REST escrita em Go que modela entidades do domínio agrícola (pessoas, fazendas, culturas, fertilizantes). O projeto segue princípios de arquitetura limpa, separando domínio, casos de uso, adaptadores e infraestrutura.

## 2. Arquitetura do Projeto

- Domínio: `internal/domain` — entidades e contratos.
- Usecases: `internal/usecases` — orquestração da lógica de negócio.
- Adapters: `internal/adapters` — handlers HTTP, DTOs, repositórios DB.
- Infraestrutura: `internal/infrastructure` — persistência, segurança, configuração.
- Entrypoint: `cmd/api/main.go` — inicialização do servidor.

## 3. Estrutura de Pastas (resumo)

- `cmd/api/main.go`
- `config/config.go`
- `internal/adapters/http/handlers`
- `internal/adapters/http/dto`
- `internal/adapters/database/mysql`
- `internal/domain`
- `internal/infrastructure/persistence/models.go`
- `internal/infrastructure/security`

## 4. Configuração e Execução Local

1. Levantar dependências (ex.: Docker Compose):

```bash
docker-compose up -d
```

2. Ajustar variáveis de ambiente (ex.: `DB_HOST`, `DB_USER`, `DB_PASS`, `JWT_SECRET`).

3. Executar a API:

```bash
go run cmd/api/main.go
```

4. Rodar testes:

```bash
go test ./... -v
```

## 5. Pontos-chave do Código

- Handlers (`internal/adapters/http/handlers`): convertem requests em chamadas aos usecases.
- Usecases (`internal/usecases`): aplicam regras de negócio e coordenam repositórios/serviços.
- Repositórios (`internal/adapters/database/mysql`): acesso ao MySQL e mapeamento para entidades.
- Segurança (`internal/infrastructure/security`): `password_service.go` e `jwt_service.go`.

## 6. Fluxos da API (sequências)

- Fluxo: Criar Pessoa (cadastro)
  1. POST `/persons` com JSON (nome, email, senha).
  2. `person_handler.Create` valida e converte payload.
  3. `personUsecase.Create` valida regras e chama `password_service.Hash`.
  4. `personRepository.Save` persiste a entidade.
  5. Retorna `201 Created`.

- Fluxo: Login / Obter JWT
  1. POST `/auth/login` com email e senha.
  2. `authUsecase.Authenticate` busca usuário e verifica senha.
  3. `jwt_service.Generate` cria token com claims.
  4. Retorna `200 OK` com `{ "token": "..." }`.

- Fluxo: Criar Fazenda (autenticado)
  1. POST `/farms` com `Authorization: Bearer <token>`.
  2. Middleware valida token e injeta `userID` no contexto.
  3. `farmUsecase.Create` persiste a fazenda com `ownerID`.
  4. Retorna `201 Created`.

- Fluxo: Listar Fazendas por Dono
  1. GET `/farms/owner/{owner_id}`.
  2. `farmUsecase.ListByOwner` chama `farmRepository.FindByOwner`.
  3. Retorna `200 OK` com lista de fazendas.

## 7. Trechos de Código

Exemplo simplificado de `person_handler` (criar e autenticar):

```go
// internal/adapters/http/handlers/person_handler.go
package handlers

import (
   "net/http"
   "github.com/gin-gonic/gin"
)

type PersonHandler struct {
   personUsecase PersonUsecase
}

func (h *PersonHandler) Create(c *gin.Context) {
   var dto CreatePersonDTO
   if err := c.ShouldBindJSON(&dto); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
   }

   ctx := c.Request.Context()
   if err := h.personUsecase.Create(ctx, dto.ToDomain()); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
   }

   c.Status(http.StatusCreated)
}

func (h *PersonHandler) Login(c *gin.Context) {
   var dto LoginDTO
   if err := c.ShouldBindJSON(&dto); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
   }

   token, err := h.personUsecase.Authenticate(c.Request.Context(), dto.Email, dto.Password)
   if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
      return
   }

   c.JSON(http.StatusOK, gin.H{"token": token})
}
```

Exemplo simplificado de `auth_usecase`:

```go
// internal/usecases/auth_usecase.go
package usecases

import (
   "context"
)

type AuthUsecase struct {
   repo PersonRepository
   pwdService PasswordService
   jwtService JWTService
}

func (u *AuthUsecase) Authenticate(ctx context.Context, email, password string) (string, error) {
   person, err := u.repo.FindByEmail(ctx, email)
   if err != nil { return "", err }

   if !u.pwdService.Verify(password, person.HashedPassword) {
      return "", ErrInvalidCredentials
   }

   token, err := u.jwtService.Generate(person.ID, person.Role)
   if err != nil { return "", err }

   return token, nil
}
```

Exemplo de interface de repositório (`FindByOwner`):

```go
// internal/domain/farm/repository.go
package farm

import "context"

type Repository interface {
   Save(ctx context.Context, f *Farm) error
   FindByOwner(ctx context.Context, ownerID string) ([]*Farm, error)
}
```

Esboço de implementação MySQL:

```go
// internal/adapters/database/mysql/farm_repository.go
func (r *mysqlFarmRepository) FindByOwner(ctx context.Context, ownerID string) ([]*domain.Farm, error) {
   rows, err := r.db.QueryContext(ctx, "SELECT id, name, area FROM farms WHERE owner_id = ?", ownerID)
   if err != nil { return nil, err }
   defer rows.Close()
   var farms []*domain.Farm
   for rows.Next() {
      var f domain.Farm
      rows.Scan(&f.ID, &f.Name, &f.Area)
      farms = append(farms, &f)
   }
   return farms, nil
}
```

## 8. Exemplos (cURL)

```bash
# Cadastro
curl -X POST http://localhost:8080/persons -H "Content-Type: application/json" -d '{"name":"Ana","email":"a@x.com","password":"secret"}'

# Login
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d '{"email":"a@x.com","password":"secret"}'

# Usar token para criar fazenda
curl -X POST http://localhost:8080/farms -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"name":"Fazenda A","area":100}'
```

## 9. Exercícios Práticos

1. Criar endpoint `GET /farms/owner/{owner_id}`: atualizar interface e implementação do repositório; adicionar handler e rota; escrever testes.
2. Implementar refresh tokens: endpoint `/auth/refresh` e armazenamento seguro de tokens.
3. Escrever testes unitários para `auth_usecase` usando mocks.
4. Adicionar validações na entidade `crop` (área/produção).

## 10. Referências

- Arquitetura limpa e padrões hexagonais
- Documentação do Go Modules e `testing`
- Boas práticas com JWT e gerenciamento de senhas

---

Se quiser, posso gerar slides em Markdown (Reveal) ou adicionar um teste unitário para `auth_usecase`.
