CREATE TABLE product (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  price FLOAT NOT NULL,
  description VARCHAR(500) NOT NULL
);

CREATE TABLE user_api
(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(3000) NOT NULL,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stockproduct
(
    id SERIAL PRIMARY KEY NOT NULL,
    productId integer NOT NULL,
    quantity integer NOT NULL,
    balance integer NOT NULL,
	FOREIGN KEY (productId) REFERENCES product(id)
);

CREATE TABLE queue_message_process
(
    id SERIAL PRIMARY KEY NOT NULL,
    message VARCHAR(250) NOT NULL,
    result VARCHAR(2) NOT NULL,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
