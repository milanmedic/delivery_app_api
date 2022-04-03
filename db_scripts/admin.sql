CREATE TABLE IF NOT EXISTS administrator (
    "id" TEXT,
    "username" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "surname" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "date_of_birth" TEXT,
    "address" INTEGER NOT NULL,
    "role" TEXT NOT NULL,
    PRIMARY KEY("id"),
    FOREIGN KEY("address")
        REFERENCES address(id)
);