CREATE TABLE IF NOT EXISTS cloud_pockets (
    "id" SERIAL PRIMARY KEY,
    "balance" float8 NOT NULL DEFAULT 0,
    "name" TEXT NOT NULL,
    "category" TEXT NOT NULL,
<<<<<<< HEAD
    "currency" TEXT NOT NULL DEFAULT "THB"
    -- PRIMARY KEY ("id")
=======
    "currency" TEXT NOT NULL 
        -- PRIMARY KEY ("id")
>>>>>>> f852357add66335a098c330ad6b9c780ec95e09c
);