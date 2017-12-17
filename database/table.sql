CREATE TABLE key_info(
  time DATETIME NOT NULL, --アクション時刻
  state ENUM('ON','OFF') NOT NULL, --鍵の状態
  key_info_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY --ID
);
