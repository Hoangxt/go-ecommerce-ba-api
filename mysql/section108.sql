-- Section 108: LÀM CHỦ MYSQL - Nguyên tắc sử dụng 6 loại index và công thức tối ưu.

CREATE TABLE users_vip (
    id INT(11) NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
-- 1. Normal Index (Chỉ mục thông thường) or Single Column Index (Chỉ mục cột đơn)
CREATE INDEX index_username ON users_vip (username);
-- 2. Unique Index (Chỉ mục duy nhất)
CREATE UNIQUE INDEX index_email ON users_vip (email);
-- 3. Primary Key Index (Chỉ mục khóa chính)
PRIMARY KEY (id);
-- 4. Fulltext Index (Chỉ mục toàn văn)
CREATE FULLTEXT INDEX index_fulltext_username ON users_vip (username);
-- 5. Composite Index (Chỉ mục tổng hợp)
CREATE INDEX index_username_email ON users_vip (username, email);

-------
CREATE TABLE test_table_001 (
    id INT PRIMARY KEY,
    a INT,
    b INT,
    c INT,
    d INT,
    INDEX idx_abc (a, b, c) -- Composite Index
)

SHOW INDEX FROM test_table_001;

EXPLAIN SELECT * FROM test_table_001 WHERE a = 1;
-- OK
EXPLAIN
SELECT *
FROM test_table_001
WHERE
    a = 1
    AND b = 2 -- OK;
EXPLAIN
SELECT *
FROM test_table_001
WHERE
    b = 1
    AND c = 2 -- NOT OK;
    -----------
EXPLAIN
SELECT *
FROM test_table_001
WHERE
    a = 1
    AND c = 2;
-- sử dụng mỗi index a mà thôi

------------------
EXPLAIN SELECT a FROM test_table_001 WHERE b = 1 AND c = 2;
--ok
EXPLAIN SELECT * FROM test_table_001 WHERE b = 1 AND a = 2;
--ok

-- Nguyên tắc ngoài cùng bên trái

CREATE TABLE orders (
    orderNumber INT PRIMARY KEY,
    order_date DATE,
    required_date DATE,
    shipped_date DATE,
    status INT,
    comments VARCHAR(255),
)

CREATE INDEX idx_order_status ON orders (orderNumber, status);
-- đúng
CREATE INDEX idx_status_order ON orders (status, orderNumber);
-- sai
-- Các trường hơp kết quả giống nhau thì không sử dụng index (status)
-- còn các trường hơp kết quả khác nhau thì sử dụng index (orderNumber)

--------------------
-- Công thức chọn lựa index bên trái.
SELECT COUNT(DISTINCT orderNumber) / COUNT(1) as o, COUNT(DISTINCT status) / COUNT(1) as s
FROM orders;

--> bên trái là orderNumber (vì orderNumber có số lượng lớn hơn status) [1.0, 0.5]