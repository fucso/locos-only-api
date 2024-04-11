
-- +migrate Up
CREATE TABLE "event_dates" (
  "id" SERIAL PRIMARY KEY,
  "event_id" integer NOT NULL,
  "date" date NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_dates" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

-- +migrate Down
DROP TABLE "event_dates";
