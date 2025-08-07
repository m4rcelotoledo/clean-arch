# Sistema de Pedidos - Clean Architecture

Este projeto implementa um sistema de pedidos seguindo os princípios da Clean Architecture em Go, com múltiplas interfaces de comunicação (REST API, gRPC, GraphQL) e integração com RabbitMQ para eventos.

## 🏗️ Arquitetura

O projeto segue os princípios da Clean Architecture com as seguintes camadas:

- **Entities**: Regras de negócio centrais (Order)
- **Use Cases**: Casos de uso da aplicação (CreateOrder)
- **Infrastructure**: Implementações concretas (Database, Web, gRPC, GraphQL)
- **Events**: Sistema de eventos com RabbitMQ

## 🚀 Tecnologias Utilizadas

- **Go 1.19**
- **MySQL 5.7** - Banco de dados
- **RabbitMQ 3.13.7** - Message broker para eventos
- **GraphQL** - Interface de consulta
- **gRPC** - Comunicação RPC
- **REST API** - Interface HTTP
- **Wire** - Injeção de dependência
- **Chi** - Router HTTP
- **GQLGen** - Geração de código GraphQL

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
├── docker-compose.yaml     # Serviços de infraestrutura
├── go.mod                  # Dependências Go
└── gqlgen.yml              # Configuração GraphQL
```

## 🛠️ Pré-requisitos

- Go 1.19 ou superior
- Docker e Docker Compose
- Make (opcional, para comandos de automação)

## 🚀 Como Executar

### 1. Iniciar Serviços de Infraestrutura

```bash
# Iniciar MySQL e RabbitMQ
docker-compose up -d
```

### 2. Aguardar a inicialização do banco de dados

O MySQL precisa de alguns segundos para inicializar completamente. Aguarde até que o container esteja pronto.

### 3. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```env
# Database
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders

# Servers
WEB_SERVER_PORT=8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080

# RabbitMQ
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
```

**Importante**: Certifique-se de que o arquivo `.env` está na raiz do projeto (mesmo nível do `go.mod`).

### 4. Instalar Dependências

```bash
go mod download
```

### 5. Gerar Código GraphQL (se necessário)

```bash
go run github.com/99designs/gqlgen generate
```

### 6. Gerar Código gRPC (se necessário)

Se você estiver usando ASDF, pode ser necessário usar o caminho completo das ferramentas:

```bash
# Gerar código protobuf
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       internal/infra/grpc/protofiles/order.proto

# Ou usar o caminho completo se as ferramentas não estiverem no PATH:
# ~/.asdf/installs/golang/1.24.5/bin/protoc --go_out=. --go_opt=paths=source_relative \
#        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#        internal/infra/grpc/protofiles/order.proto
```

### 7. Executar a Aplicação

```bash
go run cmd/ordersystem/main.go wire_gen.go
```

## 🌐 Interfaces Disponíveis

### Portas dos Serviços

- **REST API**: Porta 8000
- **GraphQL**: Porta 8080
- **gRPC**: Porta 50051
- **MySQL**: Porta 3306
- **RabbitMQ**: Porta 5672
- **RabbitMQ Management**: Porta 15672

### REST API (Porta 8000)

**Criar Pedido:**
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "order-123",
    "price": 150.50,
    "tax": 20.54
  }'
```

**Listar Pedidos:**
```bash
curl -X GET http://localhost:8000/order
```

### GraphQL (Porta 8080)

Acesse o playground: http://localhost:8080

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

## 📊 Monitoramento

- **RabbitMQ Management**: http://localhost:15672
  - Usuário: `guest`
  - Senha: `guest`

## 🧪 Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com coverage
go test -cover ./...
```

## 📝 Exemplo de Uso

### Via REST API

```bash
# Criar um pedido
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "order-001",
    "price": 100.00,
    "tax": 15.00
  }'
```

Resposta esperada:
```json
{
  "id": "order-001",
  "price": 100.00,
  "tax": 15.00,
  "final_price": 115.00
}
```

```bash
# Listar todos os pedidos
curl -X GET http://localhost:8000/order
```

Resposta esperada:
```json
[
  {
    "id": "order-001",
    "price": 100.00,
    "tax": 15.00,
    "final_price": 115.00
  }
]
```

### Via GraphQL

1. Acesse http://localhost:8080
2. Execute a mutation para criar:

```graphql
mutation CreateOrder {
  createOrder(input: {
    id: "order-002"
    Price: 200.00
    Tax: 30.00
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

3. Execute a query para listar:

```graphql
query ListOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

## 🔧 Desenvolvimento

### Gerar Código Wire (Injeção de Dependência)

```bash
go install github.com/google/wire/cmd/wire@latest
wire ./cmd/ordersystem
```

### Gerar Código GraphQL

```bash
go run github.com/99designs/gqlgen generate
```

### Estrutura de Dados

O sistema trabalha com a entidade `Order` que possui:
- `ID`: Identificador único do pedido
- `Price`: Preço do produto
- `Tax`: Taxa aplicada
- `FinalPrice`: Preço final (Price + Tax)

## 🐳 Docker

Para executar apenas os serviços de infraestrutura:

```bash
docker-compose up -d
```

Para parar os serviços:

```bash
docker-compose down
```

## 📋 Funcionalidades

- ✅ Criação de pedidos via REST API
- ✅ Listagem de pedidos via REST API
- ✅ Criação de pedidos via GraphQL
- ✅ Listagem de pedidos via GraphQL
- ✅ Criação de pedidos via gRPC
- ✅ Listagem de pedidos via gRPC
- ✅ Sistema de eventos com RabbitMQ
- ✅ Validação de dados
- ✅ Cálculo automático do preço final
- ✅ Persistência em MySQL
- ✅ Clean Architecture
- ✅ Injeção de dependência com Wire

## 🔍 Logs e Monitoramento

O sistema emite eventos quando um pedido é criado. Você pode monitorar esses eventos no RabbitMQ Management Console.

## 🚨 Troubleshooting

### Problemas Comuns

1. **Erro de conexão com MySQL**: Verifique se o container está rodando
2. **Erro de conexão com RabbitMQ**: Verifique se o container está rodando
3. **Porta já em uso**: Verifique se as portas 8000, 8080, 50051 estão livres

### Comandos Úteis

```bash
# Verificar status dos containers
docker-compose ps

# Ver logs dos containers
docker-compose logs

# Reiniciar serviços
docker-compose restart
```

## 📄 Licença

Este projeto é parte do curso Go Expert da FullCycle.
