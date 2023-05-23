CREATE TABLE IF NOT EXISTS "abstract_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

