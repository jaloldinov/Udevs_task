CREATE TABLE "authors" (
  "id" UUID PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book_category" (
  "id" UUID PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" UUID PRIMARY KEY,
  "book_name" varchar NOT NULL,
  "author_id" UUID NOT NULL,
  "category_id" UUID NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("category_id") REFERENCES "book_category" ("id");
