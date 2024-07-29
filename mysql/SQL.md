# Section 54: MySQL Setup master slave | SQL

docker network create my_master_slave_mysql

docker run -d --name mysql-master --network my_master_slave_mysql -p 8811:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql/mysql-server:latest

docker run -d --name mysql-slave --network my_master_slave_mysql -p 8822:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql/mysql-server:latest

docker exec -it mysql-master mysql -uroot -p123456

# Section 104: Thao tác cần biết khắc phục LỖI MYSQL trực tuyến | MySQL Series

- show processlist;

* ID: Giá trị của thread hiện tại đang chạy
* User: Tên người dùng kết nối
* Host: Địa chỉ của client duy trì mqh kết nối của luồng hiện tại
* DB: Tên cơ sở dữ liệu mà luồng hiện tại đang sử dụng
* Command: Lệnh SQL đang thực thi [sleep, query, connect, binlog dump, close, kill, init db, query cache, table dump, refresh, change user, shutdown, statistics]
* Time: Thời gian luồng hiện tại đã chạy
* State: Trạng thái của luồng hiện tại [Updating, Sleep, Query, Locked, Sending, ...]
* Info: Thông tin về lệnh SQL đang thực thi

- Connection pool

# Section 106: LÀM CHỦ MYSQL - HIỂU NHẦM VỀ INDEX PRIMARY, UNIQUE, FULLTEXT?

- các cú pháp cơ bản tạo index trong MySQL:
  CREATE INDEX index_name ON table_name(column_name(length)) [ASC | DESC];

  ALTER TABLE table_name ADD INDEX index_name(column_name(length) [ASC | DESC]);

- EXP:
  CREATE TABLE `member_vip`(
  `member_id` INT(11) NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(255) NOT NULL,
  PRIMARY KEY(`member_id`) USING BTREE,
  UNIQUE KEY `email`(`email`)
  ) ENGINE=InnoDB COLLATE='utf8_general_ci';

  SHOW INDEX FROM `member_vip`;
  DROP INDEX `email` ON `member_vip`;

  CREATE UNIQUE INDEX index_name ON table_name(column_name(length)) [ASC | DESC];
  ALTER TABLE table_name ADD UNIQUE index_name(column_name(length) [ASC | DESC]);

  CREATE FULLTEXT INDEX index_name ON table_name(column_name(length));
  ALTER TABLE table_name ADD FULLTEXT index_name(column_name(length));
