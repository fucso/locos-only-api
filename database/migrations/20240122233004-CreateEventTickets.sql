
-- +migrate Up
CREATE TYPE "event_ticket_types" AS ENUM (
  'adv',
  'door',
  'other'
);

CREATE TABLE "event_tickets" (
  "id" SERIAL PRIMARY KEY,
  "event_id" integer NOT NULL,
  "type" event_ticket_types NOT NULL,
  "name" VARCHAR(255),
  "price" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);
ALTER TABLE "event_tickets" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

-- +migrate Down
DROP TABLE "event_tickets";
DROP TYPE "event_ticket_types";
