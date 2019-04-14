CREATE TABLE items (
  ID            SERIAL      PRIMARY KEY,
  name          varchar(40) NOT NULL,
  type          varchar(40) NOT NULL,
  protein       INTEGER     NOT NULL,
  carbs         INTEGER     NOT NULL,
  fats          INTEGER     NOT NULL,
  calories      INTEGER     NOT NULL,
  calsperpro    DECIMAL     NOT NULL,
  sodium        INTEGER     NOT NULL,
  r_id          INTEGER     REFERENCES items(id) ON DELETE CASCADE ON UPDATE CASCADE
);