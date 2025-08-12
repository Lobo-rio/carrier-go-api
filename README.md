# carrier-go-api 🚀

## Descrição

A **carrier-go-api** é uma API desenvolvida em Go para gerenciar operações de transportadoras, clientes, produtos e pedidos. Ideal para sistemas logísticos e de gestão comercial, o projeto utiliza tecnologias modernas como GORM e Chi para garantir escalabilidade, segurança e facilidade de manutenção.

---

## Estrutura do Projeto 📂

```plaintext
carrier-go-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── contracts/
│   │   ├── carrier/
│   │   │   └── CreateCarrier.go
│   │   ├── clients/
│   │   │   └── CreateClient.go
│   │   ├── products/
│   │   │   └── CreateProduct.go
│   ├── domain/
│   │   ├── carrier/
│   │   │   ├── carrier.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── clients/
│   │   │   ├── clients.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── products/
│   │   │   ├── products.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   ├── internal-errors/
│   │   ├── error.go
│   │   └── validator.go
├── go.mod
├── go.sum
└── README.md
```

---

## Destaques do Projeto 🌟

![API](https://via.placeholder.com/600x200?text=Carrier+API+Logo)

- **Tecnologias Utilizadas:**

  - GORM para ORM
  - Chi para rotas HTTP
  - Testify para testes
  - Godotenv para variáveis de ambiente
  - Gomail para envio de e-mails

- **Funcionalidades:**
  - Cadastro, consulta, atualização e exclusão de transportadoras, clientes, produtos e pedidos
  - Gerenciamento de relacionamentos entre entidades
  - Validações automáticas

---

## Instalação 🛠️

1. **Pré-requisitos:**

   - Go 1.20 ou superior
   - PostgreSQL
   - Git

2. **Clone o repositório:**

   ```sh
   git clone https://github.com/Lobo-rio/carrier-go-api.git
   cd carrier-go-api
   ```

3. **Configure o banco de dados:**

   - Crie um banco chamado `carrier_api` no PostgreSQL
   - Ajuste o arquivo de conexão em `internal/infra/database/connection.go` com usuário, senha e host corretos

4. **Instale as dependências:**

   ```sh
   go mod tidy
   ```

5. **Execute as migrações:**

   - As tabelas serão criadas automaticamente ao iniciar a aplicação

6. **Inicie a API:**

   ```sh
   go run cmd/api/main.go
   ```

---

## Configuração de Ambiente 🌍

- Variáveis de ambiente podem ser usadas para configurar o banco de dados, porta e outros parâmetros.
- Exemplo de configuração no arquivo `.env`:

  ```env
  DB_HOST=localhost
  DB_USER=postgres
  DB_PASSWORD=senha
  DB_NAME=carrier_api

  GMAIL_PASSWORD=senha
  GMAIL_USER=email@gmail.com
  GMAIL_SMTP=smtp.gmail.com
  ```

---

## Como Usar a API 📡

- Acesse os endpoints via HTTP (exemplo: `http://localhost:3000`)
- Principais rotas:
  - `/carriers` (GET, POST, PATCH, DELETE)
  - `/clients` (GET, POST, PATCH, DELETE)
  - `/products` (GET, POST, PATCH, DELETE)
  - `/orders` (GET, POST, PATCH, DELETE)
- Envie requisições JSON conforme os contratos definidos em `internal/contracts/`
- Utilize ferramentas como Postman ou Insomnia para testar os endpoints

---

## Documentação 📖

Consulte o diagrama e exemplos de uso em `/docs` para envio de requisições à API.

---

## Contribuição 🤝

Contribuições são bem-vindas! Siga os passos abaixo:

1. Faça um fork do repositório
2. Crie uma branch para sua feature (`git checkout -b minha-feature`)
3. Commit suas alterações (`git commit -m 'Adicionei minha feature'`)
4. Envie para o repositório (`git push origin minha-feature`)
5. Abra um Pull Request

---

## Licença 📜

Este projeto está licenciado sob a [MIT License](LICENSE).
