CREATE TABLE key_info(
  time DATETIME NOT NULL,
  state ENUM('put','leave') NOT NULL,
  key_info_id INT NOT NULL
);
