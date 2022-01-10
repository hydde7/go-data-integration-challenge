# Data integration challenge

A porta que a API opera é a 8000

- Linguagem: Go

Tenha certeza que você possua Go, Docker e PostgreSQL instalados em sua máquina


## Packages

O projeto está dividido em 4 packages:
- database: Utilizado para fazer a conexão com o banco de dados
- models: Modelo utilizado para as transações em banco
- readCSV: Possui as funções de leitura dos arquivos CSV
- routes: Cria as rotas para requisições do tipo HTTP


## Testes

Os testes foram feitos nas funções de database (companies) para garantir que a conversa com o banco esteja sendo feita corretamente
