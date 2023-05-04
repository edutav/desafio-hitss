CREATE TABLE IF NOT EXISTS clientes (
	id serial PRIMARY KEY,
	nome VARCHAR (50) NOT NULL,
	sobrenome VARCHAR (50) NOT NULL,
	contato VARCHAR (255) NOT NULL,
    endereco VARCHAR (255) NOT NULL,
    data_nascimento VARCHAR (255) NOT NULL,
    cpf VARCHAR (255) NOT NULL
);