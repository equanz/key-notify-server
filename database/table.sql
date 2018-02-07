CREATE TABLE key_info(
  time DATETIME NOT NULL, --アクション時刻
  state ENUM('ON','OFF') NOT NULL, --鍵の状態
  key_info_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY --ID
);

CREATE TABLE app_id(
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, -- ID
  app_id VARCHAR(48) NOT NULL -- app_id
);

