CREATE TABLE people (
    name VARCHAR(255) NOT NULL PRIMARY KEY,
    count INT NOT NULL
);

-- CREATE DATABASE IF NOT EXISTS hello;

-- USE hello;

-- CREATE TABLE availability (
--     weekday SMALLINT NOT NULL,
--     start_time TIME NOT NULL,
--     end_time TIME NOT NULL,
--     venue_id BIGINT NOT NULL,
--     PRIMARY KEY (weekday, venue_id),
--     FOREIGN KEY (venue_id) REFERENCES venue(venue_id)
-- );

-- CREATE TABLE booking (
--     id BIGINT AUTO_INCREMENT PRIMARY KEY,
--     start_time TIMESTAMP NOT NULL,
--     end_time TIMESTAMP NOT NULL,
--     venue_id BIGINT NOT NULL,
--     email TEXT NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     status ENUM('PENDING', 'APPROVED', 'REJECTED') NOT NULL,
--     reason TEXT,
--     approver_id BIGINT,
--     FOREIGN KEY (venue_id) REFERENCES venue(venue_id),
--     FOREIGN KEY (approver_id) REFERENCES approver(approver_id)
-- );

-- CREATE TABLE venue_type (
--     type_id BIGINT AUTO_INCREMENT PRIMARY KEY,
--     type_name TEXT NOT NULL UNIQUE
-- );

-- CREATE TABLE venue (
--     venue_id BIGINT AUTO_INCREMENT PRIMARY KEY,
--     name TEXT NOT NULL,
--     type_id BIGINT NOT NULL,
--     FOREIGN KEY (type_id) REFERENCES venue_type(type_id)
-- );

-- CREATE TABLE approver (
--     approver_id BIGINT AUTO_INCREMENT PRIMARY KEY,
--     email TEXT NOT NULL UNIQUE,
--     name TEXT NOT NULL,
--     hashed_password TEXT NOT NULL
-- );
