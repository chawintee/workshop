CREATE TABLE IF NOT EXISTS cloud_pockets (
    "id" SERIAL PRIMARY KEY,
    "balance" float8 NOT NULL DEFAULT 0,
    "name" TEXT NOT NULL,
    "category" TEXT NOT NULL,
    "currency" TEXT NOT NULL DEFAULT "THB",
    -- PRIMARY KEY ("id")
);