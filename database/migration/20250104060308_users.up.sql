CREATE TABLE users
(
    id         INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    username   VARCHAR(50)                    NOT NULL UNIQUE,
    full_name  VARCHAR(255)                   NOT NULL,
    password   VARCHAR(255)                   NOT NULL,
    role       ENUM ('admin', 'chaser', 'customer') DEFAULT 'admin',
    created_at TIMESTAMP                            DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
