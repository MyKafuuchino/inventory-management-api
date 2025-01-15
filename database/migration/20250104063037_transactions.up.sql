CREATE TABLE transactions
(
    id             INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    order_id       INT         NOT NULL,
    user_id        INT         NOT NULL,
    total_price    INT         NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    transaction_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);