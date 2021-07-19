-- +migrate Up
CREATE TABLE pilots (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE jets (
  id SERIAL PRIMARY KEY,
  pilot_id BIGINT UNSIGNED NOT NULL,
  age SMALLINT NOT NULL,
  name TEXT NOT NULL,
  color TEXT NOT NULL
);

ALTER TABLE jets ADD CONSTRAINT jet_pilots_fkey FOREIGN KEY (pilot_id) REFERENCES pilots(id);

-- +migrate Down
DROP TABLE jets;
DROP TABLE pilots;