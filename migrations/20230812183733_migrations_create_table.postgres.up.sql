CREATE TABLE links (
   id serial PRIMARY KEY,
   link varchar(255) NOT NULL,
   alias varchar(50) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL
);