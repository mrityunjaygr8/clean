CREATE TABLE "abstract_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
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
  "internal_id" serial NOT NULL,
  "user_id" varchar,
  "admin" bool DEFAULT false
);

CREATE TABLE "orgs" (
  "internal_id" serial NOT NULL,
  "id" varchar PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "org_users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "user_id" varchar,
  "admin" bool DEFAULT false,
  "org_id" varchar
);

CREATE TABLE "teams" (
  "id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "org_id" varchar,
  "members" varchar
);

CREATE TABLE "assets" (
  "unique_id" varchar PRIMARY KEY NOT NULL,
  "internal_id" serial NOT NULL,
  "host" varchar NOT NULL,
  "team" varchar
);

ALTER TABLE "tokens" ADD FOREIGN KEY ("user") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "admin_users" ADD FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("user_id") REFERENCES "abstract_users" ("internal_id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("internal_id");

ALTER TABLE "teams" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("internal_id");

CREATE TABLE "org_users_teams" (
  "org_users_internal_id" serial,
  "teams_members" varchar,
  PRIMARY KEY ("org_users_internal_id", "teams_members")
);

ALTER TABLE "org_users_teams" ADD FOREIGN KEY ("org_users_internal_id") REFERENCES "org_users" ("internal_id");

ALTER TABLE "org_users_teams" ADD FOREIGN KEY ("teams_members") REFERENCES "teams" ("members");


ALTER TABLE "assets" ADD FOREIGN KEY ("team") REFERENCES "teams" ("internal_id");
