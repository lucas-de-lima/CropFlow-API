# CropFlow API

API REST para gestão de operações agrícolas desenvolvida em Go, implementando Clean Architecture e Domain-Driven Design.

## Visão Geral

CropFlow é uma API que permite o gerenciamento de propriedades rurais, culturas agrícolas e fertilizantes, com controle de acesso baseado em roles. A aplicação foi construída seguindo princípios de arquitetura limpa, garantindo separação de responsabilidades e testabilidade.

## Tecnologias

- **Linguagem**: Go 1.21
- **Framework HTTP**: Gin
- **Banco de Dados**: MySQL 8.0
- **ORM**: GORM
- **Autenticação**: JWT
- **Containerização**: Docker e Docker Compose

## Arquitetura

O projeto segue os princípios de Clean Architecture, organizando o código em camadas:

- **Domain**: Entidades de negócio e regras de domínio
- **Use Cases**: Lógica de aplicação e orquestração
- **Adapters**: Implementações de interfaces (HTTP handlers, repositórios MySQL)
- **Infrastructure**: Serviços de infraestrutura (JWT, criptografia de senhas)

<details>
<summary>Estrutura do Projeto</summary>

```
cropflow-api/
├── cmd/api/                    # Ponto de entrada da aplicação
├── config/                     # Gerenciamento de configuração
├── internal/
│   ├── adapters/
│   │   ├── database/mysql/     # Implementações de repositórios MySQL
│   │   └── http/
│   │       ├── handlers/       # Controllers HTTP
│   │       ├── dto/           # Data Transfer Objects
│   │       └── routes/        # Configuração de rotas
│   ├── domain/
│   │   ├── entities/          # Entidades de domínio
│   │   ├── repositories/      # Interfaces de repositórios
│   │   ├── farm/             # Domínio de fazendas
│   │   ├── crop/             # Domínio de culturas
│   │   ├── fertilizer/       # Domínio de fertilizantes
│   │   └── person/           # Domínio de usuários
│   ├── infrastructure/
│   │   └── security/         # Serviços de segurança (JWT, senhas)
│   └── usecases/             # Casos de uso da aplicação
├── docker-compose.yml
├── Dockerfile
└── docs/                      # Documentação adicional
```

</details>

## Pré-requisitos

- Docker e Docker Compose
- Go 1.21+ (apenas para desenvolvimento local)
- MySQL 8.0 (ou uso do Docker Compose)

## Instalação e Execução

### Usando Docker Compose (Recomendado)

```bash
# Clone o repositório
git clone <repository-url>
cd cropflow-api

# Inicie os serviços
docker compose up -d

# Verifique se os serviços estão rodando
docker ps

# A API estará disponível em http://localhost:8080
```

### Execução Local

```bash
# Instale as dependências
go mod download

# Configure as variáveis de ambiente (veja seção abaixo)

# Inicie o MySQL com Docker
docker run -d --name cropflow-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=cropflow \
  -p 3306:3306 \
  mysql:8.0

# Execute a aplicação
go run ./cmd/api
```

<details>
<summary>Variáveis de Ambiente</summary>

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `DB_HOST` | Host do banco de dados | `localhost` |
| `DB_PORT` | Porta do banco de dados | `3306` |
| `DB_USER` | Usuário do banco de dados | `root` |
| `DB_PASSWORD` | Senha do banco de dados | (vazio) |
| `DB_NAME` | Nome do banco de dados | `cropflow` |
| `JWT_SECRET` | Chave secreta para JWT | `your-secret-key` |
| `JWT_ISSUER` | Emissor do token JWT | `cropflow` |
| `PORT` | Porta do servidor HTTP | `8080` |

**Importante**: Altere `JWT_SECRET` em produção por uma chave segura.

</details>

## Autenticação e Autorização

A API utiliza JWT (JSON Web Tokens) para autenticação. Após criar um usuário via `POST /persons`, é necessário realizar login via `POST /auth/login` para obter um token.

O token deve ser enviado no header `Authorization` no formato:
```
Authorization: Bearer <token>
```

### Roles e Permissões

A aplicação possui três níveis de acesso:

| Role | Descrição | Permissões |
|------|-----------|------------|
| `USER` | Usuário comum | Visualizar fazendas |
| `MANAGER` | Gerente | Visualizar fazendas e listar todas as culturas |
| `ADMIN` | Administrador | Acesso completo, incluindo gestão de fertilizantes |

<details>
<summary>Matriz de Permissões Detalhada</summary>

| Endpoint | USER | MANAGER | ADMIN |
|----------|------|---------|-------|
| `POST /persons` | ✅ | ✅ | ✅ |
| `POST /auth/login` | ✅ | ✅ | ✅ |
| `POST /farms` | ✅ | ✅ | ✅ |
| `GET /farms` | ✅ | ✅ | ✅ |
| `GET /farms/:id` | ✅ | ✅ | ✅ |
| `POST /farms/:id/crops` | ✅ | ✅ | ✅ |
| `GET /farms/:id/crops` | ✅ | ✅ | ✅ |
| `GET /crops` | ❌ | ✅ | ✅ |
| `GET /crops/:id` | ✅ | ✅ | ✅ |
| `POST /fertilizers` | ✅ | ✅ | ✅ |
| `GET /fertilizers` | ❌ | ❌ | ✅ |
| `GET /fertilizers/:id` | ✅ | ✅ | ✅ |

</details>

## Endpoints da API

### Autenticação

- `POST /persons` - Criar novo usuário
- `POST /auth/login` - Autenticar e obter token JWT

### Fazendas

- `POST /farms` - Criar fazenda
- `GET /farms` - Listar fazendas (requer autenticação)
- `GET /farms/:id` - Obter detalhes de uma fazenda

### Culturas

- `POST /farms/:id/crops` - Criar cultura em uma fazenda
- `GET /farms/:id/crops` - Listar culturas de uma fazenda
- `GET /crops` - Listar todas as culturas (requer role MANAGER ou ADMIN)
- `GET /crops/:id` - Obter detalhes de uma cultura

### Fertilizantes

- `POST /fertilizers` - Criar fertilizante
- `GET /fertilizers` - Listar todos os fertilizantes (requer role ADMIN)
- `GET /fertilizers/:id` - Obter detalhes de um fertilizante

### Associações

- `POST /crop/:cropId/fertilizer/:fertilizerId` - Associar fertilizante a uma cultura
- `GET /crop/:cropId/fertilizers` - Listar fertilizantes de uma cultura

<details>
<summary>Exemplos de Requisições</summary>

#### Criar Usuário

```bash
curl -X POST http://localhost:8080/persons \
  -H "Content-Type: application/json" \
  -d '{
    "username": "usuario",
    "password": "senha123",
    "role": "USER"
  }'
```

#### Login

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "usuario",
    "password": "senha123"
  }'
```

#### Criar Fazenda

```bash
curl -X POST http://localhost:8080/farms \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Fazenda Exemplo",
    "size": 100.5
  }'
```

#### Listar Fazendas (com autenticação)

```bash
curl -X GET http://localhost:8080/farms \
  -H "Authorization: Bearer <seu-token-jwt>"
```

</details>

## Desenvolvimento

### Executar Testes

```bash
go test ./... -v
```

### Build

```bash
go build -o cropflow-api ./cmd/api
```

### Dependências

O gerenciamento de dependências é feito via Go Modules. Para atualizar dependências:

```bash
go mod tidy
go mod download
```

<details>
<summary>Adicionando Novas Funcionalidades</summary>

Seguindo a arquitetura do projeto:

1. **Defina as entidades de domínio** em `internal/domain/`
2. **Crie as interfaces de repositório** em `internal/domain/repositories/`
3. **Implemente a lógica de domínio** nas pastas específicas do domínio
4. **Implemente os repositórios** em `internal/adapters/database/mysql/`
5. **Crie os casos de uso** em `internal/usecases/`
6. **Implemente os handlers HTTP** em `internal/adapters/http/handlers/`
7. **Defina os DTOs** em `internal/adapters/http/dto/`
8. **Configure as rotas** em `internal/adapters/http/routes/routes.go`

</details>

## Monitoramento

### Logs

Com Docker Compose:

```bash
# Logs da API
docker logs cropflow-api

# Logs do banco de dados
docker logs cropflow-mysql

# Logs de ambos
docker compose logs -f
```

### Health Check

Para verificar se a API está respondendo:

```bash
curl http://localhost:8080/farms
```

## Contribuindo

1. Faça um fork do repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Faça commit das suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Faça push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a MIT License.

---

**CropFlow API** - Sistema de Gestão Agrícola
