# golang-api-clients

Esse projeto é para a prova de desenvolvedor backend para hitss.

## Pré-requisitos

Antes de executar o projeto, você precisa ter o Docker e o Docker Compose instalados em seu ambiente de desenvolvimento. Caso ainda não tenha instalado, você pode seguir as instruções abaixo:

- **Docker**: Para instalar o Docker, siga as instruções oficiais de instalação para o seu sistema operacional, disponíveis em: https://www.docker.com/get-docker

- **Docker Compose**: O Docker Compose normalmente é instalado junto com o Docker em muitos sistemas, mas se não estiver disponível, você pode seguir as instruções de instalação aqui: https://docs.docker.com/compose/install/

Certifique-se de que o Docker e o Docker Compose estejam corretamente instalados e configurados em seu ambiente antes de prosseguir com a execução do projeto.

## Instalação

Para executar o projeto, você pode seguir os passos abaixo usando o Docker e o Docker Compose:

1. Certifique-se de ter o Docker e o Docker Compose instalados em seu ambiente de desenvolvimento (consulte a seção de Pré-requisitos para mais detalhes).

2. Clone o repositório do projeto para o seu ambiente local.

3. Navegue até o diretório do projeto no terminal.

4. Execute o comando abaixo para construir e executar os contêineres Docker do projeto:
   - docker compose up -d
5. Após a conclusão da construção dos contêineres, o projeto estará em execução e você poderá acessá-lo através do navegador ou de outras ferramentas apropriadas.

   - Para ver o swagger do projeto basta acessar o link:
     - http://localhost:8080/swagger/index.html

6. Para parar e remover os contêineres Docker do projeto, você pode executar o comando abaixo no diretório do projeto:
   docker compose down

## Contato

endereço de e-mail: edutav@gmail.com