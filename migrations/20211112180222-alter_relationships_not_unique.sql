
-- +migrate Up
ALTER TABLE relationships
  MODIFY COLUMN follower_name varchar(255) NOT NULL,
  MODIFY COLUMN followee_name varchar(255) NOT NULL
;
-- +migrate Down
