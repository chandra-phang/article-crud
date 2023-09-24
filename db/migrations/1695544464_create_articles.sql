CREATE TABLE "articles" (
  "id" varchar(100) NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar(255),
  "content" text,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
)
