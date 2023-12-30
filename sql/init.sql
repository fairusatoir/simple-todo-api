CREATE TABLE if not exists task (
    id INT PRIMARY KEY auto_increment,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date DATE,
    is_completed BOOLEAN NOT NULL DEFAULT false
)engine=InnoDB;