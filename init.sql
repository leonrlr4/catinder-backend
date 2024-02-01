CREATE DATABASE IF NOT EXISTS catinder;

USE catinder;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    picture VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    oauth_provider VARCHAR(255) DEFAULT '',
    jwt_token VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
    users (username, email, password)
VALUES
    ('cccccs', 'john@gmail.com', '123456');

CREATE TABLE IF NOT EXISTS cats (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cat_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    breed VARCHAR(255) NOT NULL,
    owner_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

INSERT INTO
    cats (cat_name, age, breed, owner_id)
VALUES
    ('Garfield', 5, 'Persian', 1),
    ('Tom', 3, 'Siamese', 1),
    ('Felix', 4, 'Tabby', 1);