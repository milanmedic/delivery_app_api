CREATE TABLE IF NOT EXISTS customer_order (
    "id" INTEGER,
    "comment" TEXT NOT NULL,
    "quantity" INTEGER NOT NULL,
    "price" INTEGER NOT NULL,
    "address" INTEGER NOT NULL,
    PRIMARY KEY("id" autoincrement),
    FOREIGN KEY("address")
        REFERENCES address(id)
);