CREATE TABLE IF NOT EXISTS parents (
  id uuid NOT NULL,
  firstname VARCHAR(40) NOT NULL,
  lastname VARCHAR(40) NOT NULL,
  email VARCHAR(80) NOT NULL,
  phone VARCHAR(15) NOT NULL,
  PRIMARY KEY(id)
);

--
-- A Child should be bound to atleast one parent
--
CREATE TABLE IF NOT EXISTS children (
  id uuid NOT NULL,
  firstname VARCHAR(40) NOT NULL,
  lastname VARCHAR(40) NOT NULL,
  age integer NOT NULL,
  parent1 uuid NOT NULL,
  parent2 uuid DEFAULT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_parent1 FOREIGN KEY(parent1) REFERENCES parents(id),
  CONSTRAINT fk_parent2 FOREIGN KEY(parent2) REFERENCES parents(id)
);