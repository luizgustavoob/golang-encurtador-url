# Encurtador de URLS

API construída a partir do livro [Programando em Go](https://www.casadocodigo.com.br/products/livro-google-go), da casa do código, cujo objetivo é encurtar uma URL, tornando-a mais acessível.

## Pré-requisitos
* [Docker](https://www.docker.com/)

## Execução

Após baixar esse repositório, você deverá acessá-lo via terminal desde a pasta raiz. Na sequência, deverá executar o comando
```
docker-compose up
```
e a API será iniciada na porta 8888.

## Encurtar URL

Para utilizar a funcionalidade da API, deve-se realizar um POST para o endpoint `/api/encurtar`, enviando no corpo da requisição a URL que se deseja encurtar. Ex:
```
  curl --location --request POST 'http://localhost:8888/api/encurtar' \
    --header 'Content-Type: text/plain' \
    --data-raw 'https://globoesporte.globo.com/futebol/brasileirao-serie-a/'
```

Como resposta, você receberá a URL encurtada.

## Estatísticas

A cada acesso na URL encurtada, a API contabiliza um acesso e disponibiliza estatísitcas para visualização. Para isso, você deve enviar um GET para o endpoint `/api/stats/{id_url_encurtada}`. Ex:
```
  curl --location --request GET 'http://localhost:8888/api/stats/hash_aleatorio_gerado'
```

Como resposta, você verá as estatísticas de acesso a URL encurtada.