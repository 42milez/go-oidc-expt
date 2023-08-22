-- Modify "auth_codes" table
ALTER TABLE `auth_codes` ADD COLUMN `expire_at` timestamp NOT NULL;
