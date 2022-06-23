Use test_db;

CREATE TABLE users (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
name VARCHAR(255) NOT NULL,
created DATETIME NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

Insert into test_tb (id, name) values (1, 'test-user');