CREATE TABLE IF NOT EXISTS "orgs" (
  "internal_id" serial NOT NULL UNIQUE,
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "org_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "user_id" serial,
  "admin" bool DEFAULT false NOT NULL,
  "org_id" serial
);

ALTER TABLE "org_users" ADD CONSTRAINT "fk_org_users_abstract_users" FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "org_users" ADD CONSTRAINT "fk_org_users_orgs" FOREIGN KEY ("org_id") REFERENCES "orgs" ("internal_id");
