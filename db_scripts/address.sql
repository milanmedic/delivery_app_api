CREATE TABLE IF NOT EXISTS address (
    "id" INTEGER,
    "city" TEXT NOT NULL,
    "street" TEXT NOT NULL,
    "street_num" TEXT NOT NULL,
    "postfix" TEXT NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
);