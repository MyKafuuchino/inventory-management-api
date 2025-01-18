CREATE TABLE orders
(
    id           INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id      INT NOT NULL,
    order_status ENUM ('pending', 'processed', 'completed', 'canceled') DEFAULT 'pending',
    total_price  INT NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);