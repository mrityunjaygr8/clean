-- This query will find out if the email provided for the abstract_users table
-- Is related to a row in the admin_users table or the org_users table
-- Provided by ChatGPT
SELECT 'admin_users' AS relation_table, admin_users.admin, admin_users.id, NULL AS org_id
FROM admin_users
WHERE user_id = <id>
UNION
SELECT 'org_users' AS relation_table, org_users.admin, org_users.id, org_users.org_id
FROM org_users
WHERE user_id = <id>;

