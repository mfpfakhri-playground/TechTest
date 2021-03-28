CREATE TABLE products(
	id          serial PRIMARY KEY,
	title       VARCHAR ( 500 ) NOT NULL,
	description VARCHAR ( 500 ) NOT NULL,
	rating      SMALLINT NOT NULL,
	image       VARCHAR (1024) NOT NULL,
	created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP 
);