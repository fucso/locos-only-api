
-- +migrate Up
CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE events;
