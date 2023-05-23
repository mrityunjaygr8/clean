CREATE TABLE "abstract_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "tokens" (
  "type" varchar NOT NULL,
  "token" varchar NOT NULL,
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "user" varchar
);

CREATE TABLE "admin_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "user_id" varchar,
  "admin" bool DEFAULT false
);

CREATE TABLE "orgs" (
  "internal_id" serial NOT NULL UNIQUE,
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "org_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "user_id" serial,
  "admin" bool DEFAULT false,
  "org_id" varchar
);

CREATE TABLE "teams" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "org_id" varchar,
  "members" varchar
);

CREATE TABLE "assets" (
  "unique_id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL UNIQUE,
  "host" varchar NOT NULL,
  "team" varchar
);

ALTER TABLE "tokens" ADD CONSTRAINT "fk_token_abstract_user" FOREIGN KEY ("user") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "admin_users" ADD CONSTRAINT "fk_admin_users_abstract_users" FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "org_users" ADD CONSTRAINT "fk_org_users_abstract_users" FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "org_users" ADD CONSTRAINT "fk_org_users_orgs" FOREIGN KEY ("org_id") REFERENCES "orgs" ("internal_id");

ALTER TABLE "teams" ADD CONSTRAINT "fk_teams_orgs" FOREIGN KEY ("org_id") REFERENCES "orgs" ("internal_id");

CREATE TABLE "org_users_teams" (
  "org_users_internal_id" serial,
  "teams_members" varchar,
  PRIMARY KEY ("org_users_internal_id", "teams_members")
);

ALTER TABLE "org_users_teams" ADD CONSTRAINT "fk_org_users_teams_org_users" FOREIGN KEY ("org_users_internal_id") REFERENCES "org_users" ("internal_id");

ALTER TABLE "org_users_teams" ADD CONSTRAINT "fk_org_users_teams_teams" ADD FOREIGN KEY ("teams_members") REFERENCES "teams" ("members");


ALTER TABLE "assets" ADD CONSTRAINT "fk_assets_teams" FOREIGN KEY ("team") REFERENCES "teams" ("internal_id");
