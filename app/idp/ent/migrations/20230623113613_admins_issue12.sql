-- Modify "admins" table
ALTER TABLE `admins` MODIFY COLUMN `id` char(36) NOT NULL, MODIFY COLUMN `totp_secret` varchar(160) NOT NULL;
