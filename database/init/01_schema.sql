CREATE TABLE IF NOT EXISTS locations (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255) NOT NULL,
    latitude DECIMAL(10, 7) NOT NULL,
    longitude DECIMAL(11, 7) NOT NULL,
    last_updated DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS evses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    location_id INT NOT NULL,
    uid VARCHAR(36) NOT NULL,
    status INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES locations(id),
    INDEX idx_evses_location_id (location_id)
);
