CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL ,
    password varchar(255) NOT NULL,
    created_at TIMESTAMP NULL,
    PRIMARY KEY (id)
)