# API de Gerenciamento de Loja de Carros (Go Car Shop)

Neste projeto, eu desenvolvi uma API RESTful completa para o gerenciamento de uma loja de veículos, utilizando Golang e MongoDB.

A aplicação foi construída seguindo um conjunto rigoroso de requisitos técnicos (detalhados no arquivo `guia.md`), demonstrando uma arquitetura de microsserviço limpa, separação de responsabilidades (MVC) e validação de dados robusta.

## Stack de Tecnologias

  * **Linguagem:** Golang (v1.19)
  * **Banco de Dados:** MongoDB (executado via Docker)
  * **Roteamento:** `gorilla/mux`
  * **Validação:** `go-playground/validator`
  * **Ambiente:** Docker Compose e `godotenv`

## Arquitetura e Funcionalidades

Eu estruturei a API seguindo uma arquitetura em camadas (similar ao MVC/DDD) para garantir um código limpo, testável e de fácil manutenção:

  * **Entities:** Define as estruturas de dados, como a `TCar`, usando structs, tags BSON/JSON e regras de validação.
  * **Models:** Camada responsável pela interação direta com o MongoDB, tratando as operações de banco (ex: `FindOneAndUpdate`).
  * **Services:** Implementa a lógica de negócios e as regras definidas nos requisitos (como validação de dados de entrada antes de chamar o Model).
  * **Controllers:** Responsáveis por receber as requisições HTTP, decodificar o JSON do corpo (body), chamar a camada de Serviço e formatar a resposta HTTP (JSON ou erro).
  * **Routes:** Mapeia os endpoints da API (ex: `/cars/{id}`) para os métodos corretos do Controller usando `gorilla/mux`.
  * **Helpers:** Utilitários para tratamento padronizado de erros customizados na API.

## Como Executar o Projeto

### Pré-requisitos

  * Go (v1.19+)
  * Docker e Docker Compose

### Passos

1.  **Clone o repositório:**

    ```bash
    git clone https://github.com/jeffthedeveloper/golang-car.git
    cd golang-car
    ```

2.  **Variáveis de Ambiente:**
    Este projeto utiliza um arquivo `.env` para carregar a URI do MongoDB. O arquivo `.env` fornecido já está configurado para funcionar com o `docker-compose.yml`.

3.  **Inicie o Banco de Dados (MongoDB):**
    Execute o Docker Compose para iniciar o container do MongoDB em background.

    ```bash
    docker-compose up -d
    ```

4.  **Instale as Dependências do Go:**

    ```bash
    go mod tidy
    ```

5.  **Execute a Aplicação:**

    ```bash
    go run main.go
    ```

    O servidor será iniciado na porta `:8080`.

## Endpoints da API

A API gerencia a coleção `cars` no banco de dados `CarShop`.

  * `POST /cars`: Cria um novo veículo.
  * `GET /cars`: Lista todos os veículos.
  * `GET /cars/{id}`: Busca um veículo específico pelo ID.
  * `PUT /cars/{id}`: Atualiza um veículo específico pelo ID.
  * `DELETE /cars/{id}`: Deleta um veículo específico pelo ID.
