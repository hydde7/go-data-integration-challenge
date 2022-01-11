# Data integration challenge

O projeto foi feito na linguagem Go, utilizando os containers do Docker com o banco de dados PostgreSQL. A partir da execução do programa, será criado uma tabela companies no banco de dados, note que ela não conterá a coluna website, para que seja anexado a coluna websites faça requisições na API, conforme o enunciado do desafio: https://github.com/eventials/data-integration-challenge

Tenha certeza que você possua Go, Docker e PostgreSQL instalados em sua máquina

A porta que a API opera é a 8000, enquanto o Docker utiliza a porta 5432 para o banco

## Packages

O projeto está dividido em 4 packages:
- database: Utilizado para fazer a conexão com o banco de dados
  - yawoenDb: Responsável pela conexão com o banco e criação da tabela
  - companies: É aqui que a conversa com a tabela companies é feita
- models: Modelo utilizado para as transações em banco
- controllers: Utilizado para controlar as ações do sistema
  - readCSV: Possui as funções de leitura dos arquivos CSV
  - routes: Cria as rotas para requisições do tipo HTTP


## Testes

Os testes foram feitos nas funções de database (companies) para garantir que a conversa com o banco esteja sendo feita corretamente

## Rotas

A rota para as requisições da api é localhost:8000/yawoenapi possuindo métodos GET e PATCH
- PATCH: Utilizado para a API dar merge com o segundo arquivo CSV e criar a coluna website na tabela companies
- GET: Utilizado para retornar um JSON com as companhias encontradas, para isto, deve-se enviar um JSON com parte do nome e zip da companhia
  - Exemplo:
         ```
         {
            "name": "Yawoen",
            "zip":"10023"
         }
         ```
