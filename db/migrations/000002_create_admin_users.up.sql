CREATE TABLE IF NOT EXISTS "admin_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "user_id" varchar,
  "admin" bool DEFAULT false
);

ALTER TABLE "admin_users" ADD CONSTRAINT "fk_admin_users_abstract_users" FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");
