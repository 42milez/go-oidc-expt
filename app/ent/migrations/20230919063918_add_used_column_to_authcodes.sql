-- Modify "auth_codes" table
ALTER TABLE `auth_codes` ADD COLUMN `used` bool NOT NULL DEFAULT 0;
