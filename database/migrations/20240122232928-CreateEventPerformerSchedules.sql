
-- +migrate Up
CREATE TABLE "event_performer_schedules" (
  "id" SERIAL PRIMARY KEY,
  "event_schedule_id" integer NOT NULL,
  "event_performer_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_performer_schedules" ADD FOREIGN KEY ("event_schedule_id") REFERENCES "event_schedules" ("id");
ALTER TABLE "event_performer_schedules" ADD FOREIGN KEY ("event_performer_id") REFERENCES "event_performers" ("id");

-- +migrate Down
DROP TABLE "event_performer_schedules";
