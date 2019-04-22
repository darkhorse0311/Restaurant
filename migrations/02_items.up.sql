CREATE TABLE items (
  ID            SERIAL       PRIMARY KEY,
  name          varchar(240) NOT NULL,
  type          varchar(240) NOT NULL,
  protein       REAL         NOT NULL,
  carbs         REAL         NOT NULL,
  fats          REAL         NOT NULL,
  calories      REAL         NOT NULL,
  calsperpro    DECIMAL      NOT NULL,
  sodium        REAL         NOT NULL,
  r_id          INTEGER      REFERENCES items(id) ON DELETE CASCADE ON UPDATE CASCADE
);