
-- +migrate Up
ALTER TABLE relationships
  MODIFY COLUMN follower_name varchar(255) unique NOT NULL,
  MODIFY COLUMN followee_name varchar(255) unique NOT NULL
;
-- +migrate Down
