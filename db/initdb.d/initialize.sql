-- CREATE DATABASE IF NOT EXISTS test;
USE test;

CREATE TABLE IF NOT EXISTS test (created_at datetime);
INSERT INTO test (created_at) VALUES(NOW());
