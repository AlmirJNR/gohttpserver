# HTTP Server in Go

Este é um projeto de servidor HTTP básico em Go, criado para fins educacionais.
O servidor possui dois endpoints: `/hello` e `/healthcheck`.

### Endpoints

`GET /hello` Retorna uma mensagem de saudação.

`GET /healthcheck` Retorna a situação de saúde do servidor.

### Executando o Projeto

Para executar o projeto localmente, siga os passos abaixo:

- Certifique-se de ter o Go instalado em sua máquina
- Você pode baixar e instalar o Go a partir de [golang.org](https://golang.org/)
- ```bash
    # Clone este repositório para o seu ambiente local
    git clone https://github.com/AlmirJNR/gohttpserver.git
    cd gohttpserver

    # Execute o comando abaixo para iniciar o servidor.
    go run cmd/httpserver/main.go
    # O servidor estará disponível em http://localhost:8080
    # Caso a porta 8080 esteja ocupada, será incrementada 1 até obter sucesso
    ```

### Estrutura de Arquivos

`cmd/httpserver/main.go` Este é o ponto de entrada do servidor.
Configura os endpoints e inicia o servidor HTTP.

#### internal/controllers

Contém a lógica de controle dos endpoints.

`healthcheck/healthcheck.go`: Implementa o endpoint GET /healthcheck.

`hello/hello.go`: Implementa o endpoint POST /hello.

#### internal/models

Define os modelos de dados para requisições e respostas.

`requests/hello.go`: Modelo de requisição para o endpoint /hello.

`responses/healthcheck.go`: Modelo de resposta para o endpoint /healthcheck.

`responses/hello.go`: Modelo de resposta para o endpoint /hello.

#### pkg/utils

Contém funções utilitárias reutilizáveis.

`writejson.go`: Função para escrever respostas JSON.

`httpjsondecoder.go`: Função para decodificar requisições JSON.

`httperrors.go`: Função para tratar erros HTTP.

`guard/strings.go`: Funções para validação de strings.

`sanitizer/strings.go`: Funções para sanitização de strings.

### Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.
Licença

Este projeto não tem licença, faça o que quiser

Se você tiver alguma dúvida, sinta-se à vontade para entrar em contato
