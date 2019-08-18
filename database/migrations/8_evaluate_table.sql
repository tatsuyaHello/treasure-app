-- +goose Up
CREATE TABLE evaluates (
  id varchar(255) NOT NULL DEFAULT "AAAA",
  method varchar(255),
  comment text,
  percentage varchar(255),
  PRIMARY KEY(id, method)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE evaluates;
