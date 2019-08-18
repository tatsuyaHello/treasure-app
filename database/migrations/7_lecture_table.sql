-- +goose Up
CREATE TABLE lectures (
  id INTEGER NOT NULL AUTO_INCREMENT,
  year VARCHAR(255) NOT NULL,
  lecture_id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  sub_title VARCHAR(255),
  english_title VARCHAR(255),
  unit INTEGER,
  semester VARCHAR(255),
  location VARCHAR(255),
  lecture_style VARCHAR(255),
  teacher VARCHAR(255),
  overview TEXT,
  goal TEXT,
  evaluate_id varchar(255),
  textbook VARCHAR(255),
  reference_url VARCHAR(255),
  remarks TEXT,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE lectures;
