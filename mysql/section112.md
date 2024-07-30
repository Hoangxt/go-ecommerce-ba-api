## Section 112: LÀM CHỦ MYSQL - 4 mức độ cô lập trong Transaction

-- VALUE ('A','Viet Nam', 1000)

-- OPEN TRANSACTION

START TRANSACTION;
-- 1. Update name
UPDATE `test_player` SET `name` = 'B' WHERE `userId` = 2;

SAVEPOINT `update_name_point`;

-- 2: Delete userId = 3
DELETE FROM `test_player` WHERE `userId` = 3;

ROLLBACK TO `update_name_point`;

SELECT \* FROM `test_player`;

COMMIT;

## 1. READ UNCOMMITTED: đọc dữ liệu chưa được commit

SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
START TRANSACTION;

UPDATE `test_player` SET `name` = 'B' WHERE `userId` = 2;

-- Chưa commit, rollback

## 2. READ COMMITTED: đọc dữ liệu đã được commit

## 3. REPEATABLE READ: đọc dữ liệu đã được commit, không cho phép đọc dữ liệu mới chưa commit

## 4. SERIALIZABLE:i đọc dữ liệu đã được commit, không cho phép đọc dữ liệu mới chưa commit, không cho phép ghi dữ liệu mớ

-- Cơ chế cô lập trong transaction mặc định của MySQL là REPEATABLE READ
