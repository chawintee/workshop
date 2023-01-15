CREATE SEQUENCE IF NOT EXISTS account_id;

CREATE TABLE IF NOT EXISTS accounts (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "transactions" (
    "transaction_id" SERIAL PRIMARY KEY,
    "source_cloud_pocket_id" int4 NOT NULL,
    "destination_cloud_pocket_id" int4 NOT NULL,
    "amount" float8 NOT NULL,
    "description" text,
    "status" text,
    "created_at" timestamp DEFAULT now()
);