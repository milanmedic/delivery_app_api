CREATE TABLE IF NOT EXISTS article_basket (
    "article" INTEGER NOT NULL,
    "basket" INTEGER NOT NULL,
    FOREIGN KEY("article")
        REFERENCES article(id)
    FOREIGN KEY("basket")
        REFERENCES basket(id)
);