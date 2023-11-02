-- +migrate up
CREATE DATABASE elastic_oj;
USE elastic_oj;

-- +migrate down
DROP DATABASE IF EXISTS elastic_oj;