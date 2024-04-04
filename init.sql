USE go-challenge_development;

-- Create tables
CREATE TABLE locations
  (
     id VARCHAR(36) NOT NULL,
     name VARCHAR(255),
     address VARCHAR(45) NOT NULL,
     coordinates POINT,
     PRIMARY KEY (id)
  );

CREATE TABLE evses
  (
     uid VARCHAR(36) NOT NULL,
     locationId VARCHAR(36) NOT NULL,
     status INT(2) NOT NULL,
     PRIMARY KEY (uid)
  );

-- import data from CSV files
LOAD DATA INFILE '/var/lib/mysql-files/locations.csv'
INTO TABLE locations
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(id, name, address, @latitude, @longitude)
SET coordinates = POINT(@longitude, @latitude);

LOAD DATA INFILE '/var/lib/mysql-files/evses.csv'
INTO TABLE evses
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(locationId, uid, status);
