CREATE TABLE cake_store
(
    id int NOT NULL AUTO_INCREMENT,
    title varchar(100) NOT NULL,
    description varchar(255) NOT NULL,
    rating float NOT NULL,
    image varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;