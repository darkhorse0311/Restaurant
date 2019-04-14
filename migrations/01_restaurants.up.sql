CREATE TABLE restaurants (
  id    serial      PRIMARY KEY,
  name  varchar(240) UNIQUE NOT NULL,
  logo  text        NOT NULL
);