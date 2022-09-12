
# Criando o Projeto:
go mod init github.com/nivaldohydalgo/go-api-rest

# Compilar e executar o Projeto:
go run main.go

# Subir os containers Docker:
docker-compose up

# Acessar o PgAdmin4:
localhost:54321

# Pegar o hostname (endereço IP) de um serviço Docker:
docker-compose exec postgres sh
### Dentro da maquina Docker executar:
hostname -i    (retorna o IP)
### Control + D para sair da maquina Docker


