# API de Biblioteca

Uma API RESTful simples para gerenciar livros, empréstimos e usuários em um sistema de biblioteca. Este projeto é desenvolvido em Go com o framework Gin, GORM para ORM e PostgreSQL como banco de dados.

## Funcionalidades

- **Livros**: Criar, ler, atualizar e excluir livros.
- **Empréstimos**: Criar, ler, atualizar e excluir empréstimos.
- **Usuários**: Criar, ler, atualizar, excluir usuários e gerenciar autenticação de usuários.

## Tecnologias Usadas

- **Go**: Linguagem de programação.
- **Gin**: Framework web para Go.
- **GORM**: Biblioteca ORM para Go.
- **PostgreSQL**: Banco de dados.
- **Docker**: Containerização.
- **Docker Compose**: Definir e executar aplicações Docker multi-contêiner.

## Estrutura do Projeto

- **`/adapter`**: Contém adaptadores para HTTP e repositórios.
  - **`/http`**: Camada HTTP, incluindo manipuladores e roteadores.
  - **`/repository`**: Interfaces e implementações para persistência de dados.
- **`/application`**: Lógica de negócios e casos de uso.
- **`/container`**: Injeção de dependência e configuração do contêiner.
- **`/domain`**: Modelos de domínio e validação.
- **`/postgress`**: Configuração de conexão com o banco de dados PostgreSQL.
- **`main.go`**: Ponto de entrada da aplicação.
- **`docker-compose.yml`**: Configuração do Docker Compose.
- **`.env`**: Configuração de variáveis de ambiente.

### Endpoints da API

- **Livros**
  - `POST /api/v1/books` - Criar um novo livro
  - `GET /api/v1/books/:id` - Obter um livro por ID
  - `PUT /api/v1/books/:id` - Atualizar um livro por ID
  - `DELETE /api/v1/books/:id` - Excluir um livro por ID

- **Empréstimos**
  - `POST /api/v1/loans` - Criar um novo empréstimo
  - `GET /api/v1/loans/:id` - Obter um empréstimo por ID
  - `PUT /api/v1/loans/:id` - Atualizar um empréstimo por ID
  - `DELETE /api/v1/loans/:id` - Excluir um empréstimo por ID

- **Usuários**
  - `POST /api/v1/users` - Criar um novo usuário
  - `POST /api/v1/users/login` - Login de usuário
  - `GET /api/v1/users/:id` - Obter um usuário por ID
  - `PUT /api/v1/users/:id` - Atualizar um usuário por ID
  - `DELETE /api/v1/users/:id` - Excluir um usuário por ID
