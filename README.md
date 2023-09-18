# Golang API com Gin e GORM

Este é um projeto de API em Go (Golang) que utiliza o framework Gin para criar endpoints RESTful e o GORM para interagir com um banco de dados PostgreSQL. Ele fornece uma estrutura básica para criar, ler, atualizar e excluir recursos (CRUD) relacionados a aberturas de vagas (openings).

## Configuração

Antes de executar a aplicação, você precisa configurar o ambiente e o banco de dados.

### Pré-requisitos

- Go (Golang) instalado
- PostgreSQL instalado e configurado
- Dependências do Go instaladas: Gin, GORM, Swagger, etc.

### Configuração do Banco de Dados

1. Crie um banco de dados PostgreSQL.
2. Atualize as informações de conexão com o banco de dados no arquivo de configuração.

### Executando a Aplicação

  go run main.go
  
  A aplicação será executada na porta padrão 8080.
  
  Uso da API
  Rotas
  GET /api/v1/opening: Obter todas as aberturas de vagas.
  GET /api/v1/opening/:id: Obter uma abertura de vaga por ID.
  POST /api/v1/opening: Criar uma nova abertura de vaga.
  PUT /api/v1/opening/:id: Atualizar uma abertura de vaga existente.
  DELETE /api/v1/opening/:id: Excluir uma abertura de vaga por ID.
  
  Documentação Swagger
  Você pode acessar a documentação Swagger em http://localhost:8080/swagger/index.html para explorar e testar as rotas da API.
  
  Contribuição
  Sinta-se à vontade para contribuir com este projeto. Abra problemas (issues) e envie solicitações de pull (pull requests) para melhorias, correções de bugs ou novos recursos.
  

