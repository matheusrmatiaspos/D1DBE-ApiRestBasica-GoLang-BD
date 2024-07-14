# 🌐 Testes - API Rest básicas (GO Lang) + Banco de Dados

A atividade consiste na criação de duas APIs Rest (cada uma em uma linguagem de programação diferente) e conectar uma delas a um Banco de Dados relacional.

## 🚧 Proposta

- Criação de duas APIs Rest (cada uma em uma linguagem de programação diferente)

- Cada API deve conter 5 rotas/endpoints, além do padrão ("/"). Cada rota deve responder às requisições GET e POST.

- Os GETs deve retornar um JSON com no mínimo 3 registros com 5 campos cada. (Ex: produtos, livros, clientes, carros, etc...)

- Os POSTs devem apenas logar no console do servidor os dados enviados pelo cliente, uma vez que ainda não estamos trabalhando com banco de dados.

- As linguagens são apenas exemplos. O(A) aluno(a) é livre para escolher a linguagem desejada.

- Usar qualquer ferramenta para testes da API (Curl, Postman, Insomnia, Insomnium, etc.). Se desejar visualizar o resultado em um navegador, utilizar alguma extensão para facilitar a visualização do JSon.

### 👉🏻 Sobre a API

- Tema: Cinema 🎬

#### ⚡ Os Endpoints
```
GET      - / 
GET/POST - /atores 
GET      - /ator/{id}
GET/POST - /diretores 
GET      - /diretor/{id}
GET/POST - /generos 
GET      - /genero/{id}
GET/POST - /filmes 
GET      - /filme/{id}
GET/POST - /analises 
GET      - /analise/{id}
```

### 🗃️ Iniciando o Banco de Dados
Execute os seguintes comandos no pgAdmin:

- Criar Banco de Dados:
```
CREATE DATABASE Cinema
```

- Criando as Tabelas:
```
CREATE TABLE actor (id serial, name varchar(50), birth_year integer, nationality varchar(50), movies_starred integer,PRIMARY KEY(id));

CREATE TABLE director (id serial, name varchar(50), birth_year integer, nationality varchar(50), movies_directed integer,PRIMARY KEY(id));

CREATE TABLE genre (id serial, name varchar(50), description varchar(50), movie_count integer, created_at varchar(50),PRIMARY KEY(id));

CREATE TABLE movie (id serial, title varchar(50),  director integer NOT NULL, genre integer NOT NULL, release_year integer,PRIMARY KEY(id),FOREIGN KEY (director) REFERENCES director(id),FOREIGN KEY (genre) REFERENCES genre(id));

CREATE TABLE review (id serial,movie integer, reviewer varchar(50), rating integer, comment varchar(100),PRIMARY KEY(id),FOREIGN KEY(movie) REFERENCES movie(id))
```

- Inserindo alguns registros:
```
INSERT INTO actor (name,birth_year,nationality,movies_starred) VALUES
('Morgan Freeman',1937,'Americano',100),
('Marlon Brando',1924,'Americano',50),
('Christian Bale',1974,'Inglês',45);

INSERT INTO director(name,birth_year,nationality,movies_directed) VALUES
('Frank Darabont',1959,'Americano',10),
('Francis Ford Coppola',1939,'Americano',30),
('Christopher Nolan',1970,'Inglês',12);

INSERT INTO genre (name, description,movie_count,created_at) VALUES
('Drama','Filmes de Drama',500,'2020-05-10'),
('Crime','Filmes de crime',200,'2020-06-15'),
('Ação','Filmes de Ação',300,'2020-07-20');

INSERT INTO movie (title,director,genre,release_year) VALUES
('Som de Liberdade',1,1,1994),
('O Padrinho',2,2,1972),
('Batman: O Cavaleiro das Trevas',3,3,2008);

INSERT INTO review (movie,reviewer,rating,comment) VALUES
(1,'John Doe',5, 'Um filme incrível com uma história comovente.'),
(2,'Jane Smith',5,'Uma obra-prima clássica com performances excepcionais.'),
(3,'Alice Johnson',4,'Um filme emocionante com excelente direção.')
```

### 📁 Inciando as Variáveis de Ambiente
Crie um arquivo ".env" na raiz do projeto e coloque suas informações do bando de dados:
```
USER= //Seu Usuário
DB_NAME= Cinema
PASSWORD= // Sua Senha
HOST= // Seu Host
```

### 🏁 Iniciando a Api
Para esta API optei por utilizar o Mux para facilitar o gerenciamento de rotas e métodos (GET e POST) sendo assim sua instalação é crucial para que a aplicação rode corretamente. Para isso utilize o comando:
```bash
go get -u github.com/gorilla/mux
```
Para realizar a conexão com o banco de dados (Postgresql) é necessário obter as seguintes bibliotecas:
```
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
```
Para utilização de variaveis de ambiente é necessário instalar:
```
go get github.com/joho/godotenv
```

Após isso pode executar a aplicação normalmente com o comando:
```bash
go run main.go
```

### 🎟️ Acessando EndPoints com Curl
- /
```bash
curl -X GET http://localhost:8080/

//Bem-Vindo a TestApiGo
```

- /atores [GET]
```bash
curl -X GET http://localhost:8080/atores
// [{"id":"1","name":"Morgan Freeman","birth_year":1937,"nationality":"American","movies_starred":100},{"id":"2","name":"Marlon Brando","birth_year":1924,"nationality":"American","movies_starred":50},{"id":"3","name":"Christian Bale","birth_year":1974,"nationality":"British","movies_starred":45}]
```

- /atores [POST]
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"id": "4","name": "Tom Hanks","birth_year": 1956,"nationality": "Americano","movies_starred": 80}' http://localhost:8080/atores
// {"id":"4","name":"Tom Hanks","birth_year":1956,"nationality":"Americano","movies_starred":80}
```
- /ator/1
```bash
curl -X GET http://localhost:8080/ator/1
// {"id":"1","name":"Morgan Freeman","birth_year":1937,"nationality":"American","movies_starred":100}
```

### 📒 Disciplina
D1DBE - Desenvolvimento Back-End I

### 🚩 Outra API (NodeJS)
Clique [aqui](https://github.com/matheusrmatiaspos/D1DBE-ApiRestBasica-NodeJS) para acessar a outra api desenvolvida com essa mesma proposta, porém com um tema e linguagem diferentes.