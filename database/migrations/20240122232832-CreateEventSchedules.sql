
-- +migrate Up
CREATE TYPE "event_schedule_types" AS ENUM (
  'open',
  'start',
  'finish',
  'close',
  'timetable',
  'other'
);

CREATE TABLE "event_schedules" (
  "id" SERIAL PRIMARY KEY,
  "event_date_id" integer NOT NULL,
  "type" event_schedule_types NOT NULL,
  "start_at" timestamp NOT NULL,
  "end_at" timestamp,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_schedules" ADD FOREIGN KEY ("event_date_id") REFERENCES "event_dates" ("id");

-- +migrate Down
DROP TABLE "event_schedules";
DROP TYPE "event_schedule_types";
