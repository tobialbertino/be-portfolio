CREATE TABLE to_do (
  id BIGSERIAL NOT NULL,
  title TEXT NOT NULL,
  status BOOLEAN NOT NULL,
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,

  PRIMARY KEY (id)
);
