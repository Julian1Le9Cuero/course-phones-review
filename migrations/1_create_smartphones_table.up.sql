CREATE TABLE smartphone(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150),
	`country_origin` VARCHAR(150) ,
    `os` VARCHAR(100),
    `price` int(10),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`Id`)
) ENGINE=InnoDB 
DEFAULT CHARSET=utf8;