# Sistema de Pedidos - Clean Architecture

Este projeto implementa um sistema de pedidos seguindo os princípios da Clean Architecture em Go, com múltiplas interfaces de comunicação (REST API, gRPC, GraphQL) e integração com RabbitMQ para eventos.

**🚀 Setup completamente automatizado com Docker!** O comando `docker compose up` inicia todo o ambiente automaticamente, incluindo a aplicação Go.

## 🏗️ Arquitetura

O projeto segue os princípios da Clean Architecture com as seguintes camadas:

- **Entities**: Regras de negócio centrais (Order)
- **Use Cases**: Casos de uso da aplicação (CreateOrder)
- **Infrastructure**: Implementações concretas (Database, Web, gRPC, GraphQL)
- **Events**: Sistema de eventos com RabbitMQ

## 🚀 Tecnologias Utilizadas

- **Go 1.24.5**
- **MySQL 8.0** - Banco de dados
- **RabbitMQ 3.13.7** - Message broker para eventos
- **GraphQL** - Interface de consulta
- **gRPC** - Comunicação RPC
- **REST API** - Interface HTTP
- **Wire** - Injeção de dependência
- **Chi** - Router HTTP
- **GQLGen** - Geração de código GraphQL
- **Docker** - Containerização e orquestração

## 📁 Estrutura do Projeto

```
20-CleanArch/
├── api/                    # Arquivos de teste da API
├── cmd/                    # Ponto de entrada da aplicação
│   └── ordersystem/
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs/                # Configurações da aplicação
├── internal/               # Código interno da aplicação
│   ├── entity/             # Entidades de domínio
│   ├── event/              # Sistema de eventos
│   ├── infra/              # Infraestrutura
│   │   ├── database/       # Camada de dados
│   │   ├── graph/          # GraphQL
│   │   ├── grpc/           # gRPC
│   │   └── web/            # REST API
│   └── usecase/            # Casos de uso
├── pkg/                    # Pacotes compartilhados
│   └── events/             # Sistema de eventos
├── migrations/             # Migrações SQL automáticas
├── Dockerfile              # Imagem da aplicação Go
├── docker-compose.yaml     # Orquestração dos serviços
├── entrypoint.sh           # Script de inicialização da aplicação
├── go.mod                  # Dependências Go
└── gqlgen.yml              # Configuração GraphQL
```

## 🛠️ Pré-requisitos

- Docker e Docker Compose
- Go 1.24.5 (apenas para desenvolvimento local)

## 🚀 Como Executar

### 🐳 Setup Automatizado com Docker

```bash
docker compose up --build
```

**O que acontece automaticamente:**
1. **MySQL**: Inicia com healthcheck e executa migrações automaticamente
2. **RabbitMQ**: Inicia com healthcheck para verificação de disponibilidade
3. **Aplicação Go**: Só inicia após MySQL e RabbitMQ estarem prontos
4. **Migrações**: Executadas automaticamente pelo MySQL
5. **Todos os endpoints**: REST, gRPC e GraphQL ficam disponíveis

### 🔧 Desenvolvimento Local

```bash
# 1. Iniciar serviços de infraestrutura
docker-compose up -d mysql rabbitmq

# 2. Configurar variáveis de ambiente (.env)
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=8080
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8081
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest

# 3. Executar a aplicação
go run cmd/ordersystem/main.go wire_gen.go
```

## 🌐 Interfaces Disponíveis

### Portas dos Serviços
- **REST API**: Porta 8080
- **GraphQL**: Porta 8081
- **gRPC**: Porta 50051
- **MySQL**: Porta 3306
- **RabbitMQ**: Porta 5672
- **RabbitMQ Management**: Porta 15672

### REST API (Porta 8080)

**Criar Pedido:**
```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "order-123",
    "price": 150.50,
    "tax": 20.54
  }'
```

**Listar Pedidos:**
```bash
curl -X GET http://localhost:8080/order
```

### GraphQL (Porta 8081)

Acesse o playground: http://localhost:8081

**Mutation para criar pedido:**
```graphql
mutation {
  createOrder(input: {
    id: "order-123"
    Price: 150.50
    Tax: 20.54
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Query para listar pedidos:**
```graphql
query {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC (Porta 50051)

Use um cliente gRPC para conectar e chamar os métodos:
- `CreateOrder` - Para criar pedidos
- `ListOrders` - Para listar pedidos

## 🐳 Arquivos Docker

### Dockerfile
- **Multi-stage build** para otimização
- **Imagem base**: Go 1.24.5-alpine para build, Alpine para runtime
- **Segurança**: Executa como usuário não-root
- **Binário**: Compila `./cmd/ordersystem` em `./app`

### docker-compose.yaml
- **MySQL**: Com healthcheck e migrações automáticas
- **RabbitMQ**: Com healthcheck
- **Aplicação**: Depende dos serviços estarem saudáveis
- **Volumes**: Mapeia `./migrations:/docker-entrypoint-initdb.d`

### entrypoint.sh
- Aguarda MySQL e RabbitMQ estarem prontos
- Só executa a aplicação após dependências estarem disponíveis

## 🧪 Testando a Aplicação

### Verificar Status dos Serviços
```bash
docker compose ps
```

### Ver Logs da Aplicação
```bash
docker compose logs -f app
```

### Testar Endpoints
```bash
# REST API
curl http://localhost:8080/order

# GraphQL
curl -X POST http://localhost:8081/query \
  -H "Content-Type: application/json" \
  -d '{"query":"query { orders { id price tax final_price } }"}'
```

### Acessar RabbitMQ Management
- URL: http://localhost:15672
- Usuário: `guest`
- Senha: `guest`

## 🛠️ Comandos Docker Úteis

```bash
# Iniciar serviços
docker compose up -d

# Parar serviços
docker compose down

# Reconstruir e iniciar
docker compose up --build -d

# Ver logs
docker compose logs -f

# Acessar shell da aplicação
docker compose exec app sh

# Acessar MySQL
docker compose exec mysql mysql -u root -proot orders
```

## 🔍 Troubleshooting

### Problemas Comuns

1. **Aplicação não inicia**: Verifique logs com `docker compose logs app`
2. **Migrações não executaram**: Verifique logs do MySQL
3. **Porta já em uso**: Pare containers e verifique portas

### Comandos de Diagnóstico

```bash
# Verificar logs da aplicação
docker compose logs app

# Verificar se MySQL está rodando
docker compose exec mysql mysqladmin ping -u root -proot

# Verificar migrações
docker compose exec mysql ls -la /docker-entrypoint-initdb.d/
```

## 📋 Funcionalidades

- ✅ Criação e listagem de pedidos via REST API, GraphQL e gRPC
- ✅ Sistema de eventos com RabbitMQ
- ✅ Validação de dados e cálculo automático do preço final
- ✅ Persistência em MySQL com migrações automáticas
- ✅ Clean Architecture com injeção de dependência (Wire)
- ✅ **Setup completamente automatizado com Docker**
- ✅ **Healthchecks para todos os serviços**
- ✅ **Inicialização automática da aplicação Go**

## 🎯 Como Funciona o Setup Docker

1. **Inicialização**: `docker compose up` inicia todos os serviços
2. **MySQL**: Executa automaticamente as migrações da pasta `./migrations`
3. **Healthchecks**: Verificam se os serviços estão saudáveis
4. **Aplicação**: Só inicia após MySQL e RabbitMQ estarem prontos
5. **Conectividade**: Todos os serviços se comunicam via rede Docker interna

## 📄 Licença

Este projeto é parte do curso Go Expert da FullCycle.
