
-- +migrate Up
ALTER TABLE relationships
  ADD PRIMARY KEY (id),
  MODIFY COLUMN id bigint(20) NOT NULL AUTO_INCREMENT
;
-- +migrate Down
