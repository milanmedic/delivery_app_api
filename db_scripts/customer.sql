CREATE TABLE IF NOT EXISTS customer (
    "id" TEXT,
    "username" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "surname" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT,
    "date_of_birth" TEXT,
    "address" INTEGER NOT NULL,
    "role" TEXT NOT NULL,
    "verification_status" TEXT DEFAULT "UNVERIFIED",
    PRIMARY KEY("id"),
    FOREIGN KEY("address")
        REFERENCES address(id)
);