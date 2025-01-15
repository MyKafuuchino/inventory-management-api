CREATE TABLE orders
(
    id          INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id     INT NOT NULL,
    total_price INT NOT NULL,
    status      ENUM ('pending', 'processed', 'completed', 'canceled') DEFAULT 'pending',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);