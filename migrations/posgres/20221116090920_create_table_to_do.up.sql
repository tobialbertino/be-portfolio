CREATE TABLE to_do (
  id BIGSERIAL NOT NULL,
  title TEXT NOT NULL,
  status BOOLEAN NOT NULL,
  created_at BIGINT,
  updated_at BIGINT,

  PRIMARY KEY (id)
);
