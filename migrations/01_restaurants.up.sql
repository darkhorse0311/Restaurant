CREATE TABLE restaurants (
  id    serial      PRIMARY KEY,
  name  varchar(40) UNIQUE NOT NULL,
  logo  text        NOTNULL
);