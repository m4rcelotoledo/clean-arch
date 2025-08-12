#!/bin/sh

# Script de entrada para aguardar serviços e executar a aplicação

echo "🚀 Iniciando aplicação de pedidos..."

# Função para aguardar o MySQL estar pronto
wait_for_mysql() {
    echo "⏳ Aguardando MySQL estar pronto..."

    # Aguardar até que o MySQL esteja aceitando conexões
    while ! nc -z mysql 3306 2>/dev/null; do
        echo "MySQL ainda não está pronto, aguardando..."
        sleep 2
    done

    # Aguardar um pouco mais para garantir que o MySQL esteja totalmente inicializado
    echo "MySQL está aceitando conexões, aguardando inicialização completa..."
    sleep 5

    echo "✅ MySQL está pronto!"
}

# Função para aguardar o RabbitMQ estar pronto
wait_for_rabbitmq() {
    echo "⏳ Aguardando RabbitMQ estar pronto..."

    # Aguardar até que o RabbitMQ esteja aceitando conexões
    while ! nc -z rabbitmq 5672 2>/dev/null; do
        echo "RabbitMQ ainda não está pronto, aguardando..."
        sleep 2
    done

    echo "✅ RabbitMQ está pronto!"
}

# Aguardar serviços estarem prontos
wait_for_mysql
wait_for_rabbitmq

echo "🎯 Todos os serviços estão prontos! Iniciando aplicação..."
echo "📍 Portas expostas:"
echo "   - REST API: 8080"
echo "   - gRPC: 50051"
echo "   - GraphQL: 8081"

# Executar a aplicação
exec ./app
