\c "books"
DROP TABLE IF EXISTS "author";
CREATE TABLE "author" (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    birthdate DATE NOT NULL ,
    death_date DATE NOT NULL
);
DROP TABLE IF EXISTS "book";
CREATE TABLE "book" (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    numberOfPages INT NOT NULL,
    description VARCHAR,
    authorID VARCHAR,
    FOREIGN KEY (authorID) REFERENCES "author"(id)
);