
-- +migrate Up
ALTER TABLE relationships
  DROP INDEX follower_name,
  DROP INDEX followee_name
;
-- +migrate Down
