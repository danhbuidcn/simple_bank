

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "unique_owner_currency";

DROP TABLE IF EXISTS "users";
