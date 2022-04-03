CREATE TABLE IF NOT EXISTS address (
    "id" INTEGER,
    "city" TEXT NOT NULL,
    "street" TEXT NOT NULL,
    "street_num" INTEGER NOT NULL,
    "postfix" TEXT NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
);