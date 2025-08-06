CREATE TABLE IF NOT EXISTS tags (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tag_rel (
    id INT PRIMARY KEY AUTO_INCREMENT,
    item_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (item_id) REFERENCES items (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);

CREATE INDEX tag_rel_item_id_idx ON tag_rel (item_id);

CREATE INDEX tag_rel_tag_id_idx ON tag_rel (tag_id);
