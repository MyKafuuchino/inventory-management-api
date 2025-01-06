CREATE TABLE order_details
(
    id         CHAR(36) PRIMARY KEY,
    order_id   CHAR(36)       NOT NULL,
    product_id CHAR(36)       NOT NULL,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);