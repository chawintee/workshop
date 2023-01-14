CREATE SEQUENCE IF NOT EXISTS account_id;

CREATE TABLE IF NOT EXISTS accounts (
    "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
    "balance" float8 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

