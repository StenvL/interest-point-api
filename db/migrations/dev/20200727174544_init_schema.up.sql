CREATE TABLE IF NOT EXISTS city (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS point_type (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS point (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(150) NOT NULL,
	description TEXT,
	type_id INT NOT NULL,
	city_id INT NOT NULL,
	lon INT,
	lat INT,

    INDEX (type_id),
    INDEX (city_id),
    FOREIGN KEY (type_id) REFERENCES point_type (id),
    FOREIGN KEY (city_id) REFERENCES city (id)
) ENGINE=INNODB;