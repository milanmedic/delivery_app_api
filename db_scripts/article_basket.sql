CREATE TABLE IF NOT EXISTS article_basket (
    "id" INTEGER,
    "article" INTEGER NOT NULL,
    "basket" TEXT NOT NULL,
    "article_quantity" INTEGER NOT NULL,
    PRIMARY KEY("id" autoincrement),
    FOREIGN KEY("article")
        REFERENCES article(id)
    FOREIGN KEY("basket")
        REFERENCES basket(id)
);