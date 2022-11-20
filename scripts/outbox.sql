\connect "ledger";

CREATE TABLE "public"."outbox" (
   "id" uuid NOT NULL,
   "aggregate_type" character varying(255) NOT NULL,
   "aggregate_id" character varying(255) NOT NULL,
   "payload" jsonb
) WITH (oids = false);