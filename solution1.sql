-- -- creating dabatase
-- CREATE DATABASE sb_coding_test;

-- -- creating table USER
-- CREATE TABLE IF NOT EXISTS users (
--     id serial PRIMARY KEY,
--     username VARCHAR (5) NOT NULL,
--     parent INT NOT NULL
-- );

-- -- inserting table data
-- INSERT INTO users VALUES
--     (1, 'Ali', 2),
--     (2, 'Budi', 0),
--     (3, 'Cecep', 1);

-- display all data with "parentusername" based on user id on parent values 
SELECT u.id, u.username, p.username as parentusername 
FROM users u
LEFT JOIN users p ON u.parent = p.id;