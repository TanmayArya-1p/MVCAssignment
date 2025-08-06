
CREATE DATABASE IF NOT EXISTS mvc;

USE mvc;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM ('admin', 'chef', 'customer') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    issued_by INT NOT NULL,
    issued_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status ENUM (
        'pending',
        'preparing',
        'served',
        'billed',
        'paid'
    ) NOT NULL DEFAULT 'pending',
    billable_amount FLOAT,
    table_no INT,
    waiter INT NULL,
    paid_at TIMESTAMP NULL,
    tip FLOAT NULL,
    FOREIGN KEY (issued_by) REFERENCES users (id),
    FOREIGN KEY (waiter) REFERENCES users (id)
);

SET GLOBAL wait_timeout = 31536000;
SET GLOBAL interactive_timeout = 31536000;
