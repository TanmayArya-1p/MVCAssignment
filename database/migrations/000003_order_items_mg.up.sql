CREATE TABLE IF NOT EXISTS order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    item_id INT NOT NULL,
    instructions VARCHAR(255),
    quantity INT NOT NULL,
    price FLOAT NOT NULL,
    issued_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status ENUM ('pending', 'preparing', 'served') NOT NULL DEFAULT 'pending',
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (item_id) REFERENCES items (id)
);

CREATE INDEX idx_order_item_order ON order_items (order_id);
