CREATE TABLE orders
(
    id          CHAR(36) PRIMARY KEY,
    user_id     CHAR(36)       NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status      ENUM ('pending', 'processed', 'completed', 'canceled') DEFAULT 'pending',
    created_at  TIMESTAMP                                              DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                                              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);