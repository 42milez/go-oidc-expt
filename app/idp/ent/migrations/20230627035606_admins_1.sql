-- Create "admins" table
CREATE TABLE `admins` (`id` char(26) NOT NULL, `name` varchar(30) NOT NULL, `password_hash` varchar(736) NOT NULL, `totp_secret` varchar(160) NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
