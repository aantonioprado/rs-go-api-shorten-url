# API Shorten URL

API desenvolvida durante o curso de **Go da Rocketseat**, com o objetivo de praticar rotas HTTP, integração com **Redis** e manipulação simples de dados para encurtar URLs.

O projeto implementa um encurtador de URLs: você envia uma URL original e recebe um código curto para acessá-la posteriormente.

## 🚀 Tecnologias

- **Go** (Golang)  
- **Redis** (noSQL)

## ⚙️ Funcionalidades

- Criar uma URL curta a partir de uma URL original  
- Redirecionar automaticamente para a URL original ao acessar o código curto

## 💡 Sobre o projeto

Este repositório foi criado durante o curso de Go da Rocketseat para praticar:
- Estrutura básica de uma aplicação Go
- Integração com Redis para armazenamento chave‑valor
- Criação de rotas HTTP
- Redirecionamento para URLs originais a partir de códigos curtos

## ▶️ Como executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/aantonioprado/rs-go-api-shorten-url
   cd rs-go-api-shorten-url
   ```

2. Suba o Redis com Docker:
   ```bash
   docker compose up -d
   ```

3. Execute a API:
   ```bash
   go run .
   ```

4. Acesse no navegador:
   ```
   http://localhost:8080
   ```
