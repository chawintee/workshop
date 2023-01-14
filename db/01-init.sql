CREATE SEQUENCE IF NOT EXISTS account_id;

CREATE TABLE "accounts" (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS transaction_id;

CREATE TABLE "transactions" (
    "transaction_id" int4 NOT NULL DEFAULT nextval('transaction_id'::regclass),
    "source_cloud_pocket_id" int4 NOT NULL,
    "destination_cloud_pocket_id" int4 NOT NULL,
    "amount" float8 NOT NULL,
    "description" text,
    PRIMARY KEY ("transaction_id")
);