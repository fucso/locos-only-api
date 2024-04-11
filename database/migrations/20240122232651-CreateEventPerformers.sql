
-- +migrate Up
CREATE TABLE "event_performers" (
  "id" SERIAL PRIMARY KEY,
  "event_id" integer NOT NULL,
  "performer_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_performers" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

-- +migrate Down
DROP TABLE "event_performers";
