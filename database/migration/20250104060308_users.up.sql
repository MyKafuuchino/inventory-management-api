CREATE TABLE users
(
    id         CHAR(36) PRIMARY KEY                 NOT NULL,
    username   VARCHAR(50)                          NOT NULL UNIQUE,
    full_name  VARCHAR(255)                         NOT NULL,
    password   VARCHAR(255)                         NOT NULL,
    role       ENUM ('admin', 'chaser', 'customer') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);