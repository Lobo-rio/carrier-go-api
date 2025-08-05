# carrier-go-api

## Descrição

A carrier-go-api é uma API desenvolvida em Go para gerenciar operações de transportadoras, clientes, produtos e pedidos, integrando funcionalidades essenciais para sistemas logísticos e de gestão comercial. O projeto foi estruturado para garantir escalabilidade, segurança e facilidade de manutenção, utilizando o GORM para persistência em banco de dados relacional e o framework Chi para rotas HTTP. A API permite o cadastro, atualização, consulta e exclusão de entidades como transportadoras, clientes, produtos e pedidos, além de gerenciar relacionamentos entre elas, como endereços, e-mails e itens de pedido. O objetivo é fornecer uma solução robusta para empresas que precisam controlar fluxos logísticos, pedidos e cadastros de forma centralizada, com validações automáticas e integração facilitada. O projeto é ideal para quem busca uma base sólida para sistemas de logística, ERP ou e-commerce, podendo ser facilmente adaptado para diferentes cenários de negócio.

---

## Instalação

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

## Configuração de Ambiente

- Variáveis de ambiente podem ser usadas para configurar o banco de dados, porta e outros parâmetros.
- Exemplo de configuração no arquivo `.env`:
  ```env
  DB_HOST=
  DB_USER=
  DB_PASSWORD=
  DB_NAME=
  DB_PORT=
  ```

---

## Como Usar a API

- Acesse os endpoints via HTTP (exemplo: `http://localhost:3000`)
- Principais rotas:
  - `/carriers` (GET, POST, PUT, DELETE)
  - `/clients` (GET, POST, PUT, DELETE)
  - `/products` (GET, POST, PUT, DELETE)
  - `/orders` (GET, POST, PUT, DELETE)
- Envie requisições JSON conforme os contratos definidos em `internal/contracts/`
- Utilize ferramentas como Postman ou Insomnia para testar os endpoints

---

## Estrutura do Projeto

- `cmd/api/main.go`: Ponto de entrada da aplicação
- `internal/domain/`: Models e regras de negócio
- `internal/contracts/`: Contratos de entrada/saída
- `internal/infra/database/`: Conexão e repositórios
- `internal/endpoints/`: Handlers HTTP
- `docs/`: Documentação

---

## Documentação

Consulte o diagrama e exemplos de uso em `/docs/diagram.md` e demais arquivos da pasta `/docs`.
