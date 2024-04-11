
-- +migrate Up
CREATE TABLE "event_venues" (
  "id" SERIAL PRIMARY KEY,
  "event_date_id" integer NOT NULL,
  "venue_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_venues" ADD FOREIGN KEY ("event_date_id") REFERENCES "event_dates" ("id");
ALTER TABLE "event_venues" ADD FOREIGN KEY ("venue_id") REFERENCES "venues" ("id");

-- +migrate Down
DROP TABLE "event_venues";
