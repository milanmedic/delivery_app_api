CREATE TABLE IF NOT EXISTS all_orders (
    "id" INT,
    "customer" TEXT NOT NULL,
    "customer_order" INTEGER NOT NULL,
    "deliverer" TEXT NOT NULL,
    "status" TEXT DEFAULT "NOT_STARTED",
    FOREIGN KEY("deliverer")
        REFERENCES user ("id")
    FOREIGN KEY("customer")
        REFERENCES user ("id")
    FOREIGN KEY("customer_order")
        REFERENCES customer_order ("id")
);