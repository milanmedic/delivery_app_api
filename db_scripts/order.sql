CREATE TABLE IF NOT EXISTS customer_order (
    "id" INTEGER,
    "comment" TEXT NOT NULL,
    "address" INTEGER NOT NULL,
    "basket" TEXT NOT NULL,
    "accepted" BOOLEAN DEFAULT FALSE,
    "deliverer" TEXT,
    "status" TEXT DEFAULT "PENDING",
    "customer" TEXT NOT NULL,
    PRIMARY KEY("id" autoincrement),
    FOREIGN KEY("address")
        REFERENCES address(id)
    FOREIGN KEY("basket")
        REFERENCES basket(id)
    FOREIGN KEY("deliverer")
        REFERENCES deliverer(id)
    FOREIGN KEY("customer")
        REFERENCES customer(id)
);