CREATE TABLE products
(
    id          CHAR(36) PRIMARY KEY,
    name        VARCHAR(100)          NOT NULL,
    description TEXT,
    price       INT                   NOT NULL,
    stock       INT       DEFAULT '0' NOT NULL,
    low_stock   INT       DEFAULT '0' NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);