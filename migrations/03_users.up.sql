CREATE TABLE IF NOT EXISTS users (
  id        serial          PRIMARY KEY,
  username  varchar(40)     UNIQUE NOT NULL,
  email     varchar(40)     UNIQUE,
  password  varchar(240)    NOT NULL
);