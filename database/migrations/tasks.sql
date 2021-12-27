CREATE TABLE tasks (
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    user_id int NOT NULL,
    title varchar(255) NOT NULL,
    content LONGTEXT NOT NULL,
    status ENUM('completed', 'uncompleted'),
    created_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)