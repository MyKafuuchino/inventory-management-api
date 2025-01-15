CREATE TABLE products
(
    id          INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name        VARCHAR(100)          NOT NULL,
    description TEXT,
    price       INT                   NOT NULL,
    stock       INT       DEFAULT '0' NOT NULL,
    low_stock   INT       DEFAULT '0' NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);