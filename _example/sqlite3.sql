DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL UNIQUE,
    "age" INTEGER NULL,
    "message" TEXT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NULL
);
