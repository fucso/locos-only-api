
-- +migrate Up
CREATE TABLE "performers" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

-- +migrate Down
Drop table "performers";
