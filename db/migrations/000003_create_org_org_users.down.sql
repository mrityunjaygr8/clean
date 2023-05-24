ALTER TABLE "org_users" DROP CONSTRAINT "fk_org_users_abstract_users";
ALTER TABLE "org_users" DROP CONSTRAINT "fk_org_users_orgs";
DROP TABLE IF EXISTS "orgs";
DROP TABLE IF EXISTS "org_users";
