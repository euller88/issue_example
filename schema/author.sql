CREATE SCHEMA library;

CREATE TABLE library.authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);