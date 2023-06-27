-- Modify "admins" table
ALTER TABLE `admins` MODIFY COLUMN `totp_secret` char(160) NULL;
