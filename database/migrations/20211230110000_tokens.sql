CREATE TABLE IF NOT EXISTS tokens (
    id INT NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    token LONGTEXT NOT NULL,
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP NULL,
    PRIMARY KEY (id)
)