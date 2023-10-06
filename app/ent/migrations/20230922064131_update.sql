-- Modify "auth_codes" table
ALTER TABLE `auth_codes` ADD COLUMN `user_id` bigint unsigned NOT NULL;
