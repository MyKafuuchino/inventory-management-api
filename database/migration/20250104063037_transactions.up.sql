CREATE TABLE transactions
(
    id             CHAR(36) PRIMARY KEY,
    order_id       CHAR(36)       NOT NULL,
    user_id        CHAR(36)       NOT NULL,
    total_price    DECIMAL(10, 2) NOT NULL,
    payment_method VARCHAR(50)    NOT NULL,
    transaction_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);