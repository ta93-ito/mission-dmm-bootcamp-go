
-- +migrate Up
CREATE TABLE relationship (
    id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    follower_name varchar(255) NOT NULL,
    followee_name varchar(255) NOT NULL
);
-- +migrate Down
