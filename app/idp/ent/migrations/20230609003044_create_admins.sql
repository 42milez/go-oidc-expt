-- Create "admins" table
CREATE TABLE `admins` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(30) NOT NULL, `password` varchar(100) NOT NULL, `totp_secret` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
