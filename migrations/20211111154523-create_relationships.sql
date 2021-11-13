
-- +migrate Up
CREATE TABLE IF NOT EXISTS relationships (id int, follower_id int, followee_id int);
