CREATE TABLE links (
   id serial PRIMARY KEY,
   full_link varchar(255) NOT NULL,
   alias varchar(10) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL
);