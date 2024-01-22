
-- +migrate Up
CREATE TABLE "venues" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

-- +migrate Down
DROP TABLE "venues";
