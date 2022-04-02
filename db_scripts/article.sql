CREATE TABLE IF NOT EXISTS article (
    "id" INTEGER,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "price" INTEGER NOT NULL,
    PRIMARY KEY("id" autoincrement),
);