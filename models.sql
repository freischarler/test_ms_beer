CREATE TABLE IF NOT EXISTS beers (
    beer_id serial NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    brewery VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL,
    currency VARCHAR(100) NOT NULL, 
    CONSTRAINT pk_beer PRIMARY KEY(beer_id)
);


