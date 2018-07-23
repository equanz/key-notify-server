CREATE TYPE state AS ENUM('ON','OFF'); -- state(ENUM)型定義

CREATE TABLE key_info(
  TIME TIMESTAMP WITH TIME ZONE NOT NULL, --アクション時刻
  state state NOT NULL, --鍵の状態
  key_info_id SERIAL NOT NULL PRIMARY KEY --ID
);

CREATE TABLE app_id(
  id SERIAL NOT NULL PRIMARY KEY, -- ID
  app_id VARCHAR(48) NOT NULL -- app_id
);

