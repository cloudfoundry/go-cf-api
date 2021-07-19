-- +migrate Up

CREATE TABLE languages (
  id SERIAL PRIMARY KEY,
  language TEXT NOT NULL
);

CREATE TABLE pilot_languages (
  pilot_id BIGINT UNSIGNED NOT NULL,
  language_id BIGINT UNSIGNED NOT NULL
);

ALTER TABLE pilot_languages ADD CONSTRAINT pilot_language_pkey PRIMARY KEY (pilot_id, language_id);
ALTER TABLE pilot_languages ADD CONSTRAINT pilot_language_pilots_fkey FOREIGN KEY (pilot_id) REFERENCES pilots(id);
ALTER TABLE pilot_languages ADD CONSTRAINT pilot_language_languages_fkey FOREIGN KEY (language_id) REFERENCES languages(id);

-- +migrate Down
DROP TABLE pilot_languages;
DROP TABLE languages;