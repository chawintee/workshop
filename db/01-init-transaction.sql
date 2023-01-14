CREATE TABLE IF NOT EXISTS "transactions" (
    "transaction_id" SERIAL PRIMARY KEY,
    "source_cloud_pocket_id" int4 NOT NULL,
    "destination_cloud_pocket_id" int4 NOT NULL,
    "amount" float8 NOT NULL,
    "description" text,
    "status" text,
    "created_at" timestamp DEFAULT now()
);