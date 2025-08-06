CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    price FLOAT NOT NULL,
    image VARCHAR(255) NOT NULL DEFAULT "/images/placeholder.png"
);
