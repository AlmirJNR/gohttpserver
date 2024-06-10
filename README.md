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
- ```shell
    # Clone este repositório para o seu ambiente local
    git clone https://github.com/AlmirJNR/gohttpserver.git
    cd gohttpserver

    # Execute o comando abaixo para iniciar o servidor.
    go run cmd/httpserver/main.go
    # O servidor estará disponível em http://localhost:8080
    # Caso a porta 8080 esteja ocupada, será incrementada 1 até obter sucesso
    ```

### Construindo o Projeto

Para construir os binários, selecione sua plataforma e rode um dos seguintes comandos de acordo:
- Linux/amd64:
    ```shell
    GOOS="linux" GOARCH="amd64" go build -ldflags="-s -w" -o ./out/linux_amd64 ./cmd/httpserver/main.go
    ```
- Windows/amd64:
    ```shell
    GOOS="windows" GOARCH="amd64" go build -ldflags="-s -w" -o ./out/windows_amd64.exe ./cmd/httpserver/main.go
    ```
    - Obs: O Windows defender pode acusar o binário do go como um trojan:
      - Mais detalhes: [Go Windows Defender no reddit](https://www.reddit.com/r/golang/comments/1au4e98/updated_to_122_now_windows_security_thinks_go_is/)
      - Uma possível solução: [Recomendação no fórum Microsoft](https://answers.microsoft.com/en-us/windows/forum/all/got-trojan-notification-after-build-golang-from/7b2c6be1-a142-4865-8aff-cf6c7ebfb9ea)

Após rodar qualquer um dos comandos, o arquivo desejado vai estar localizado dentro da pasta [out](./out)

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
