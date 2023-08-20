-- Create "users" table
CREATE TABLE `users` (`id` char(26) NOT NULL, `name` varchar(30) NOT NULL, `password_hash` varchar(1000) NOT NULL, `totp_secret` char(160) NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "auth_codes" table
CREATE TABLE `auth_codes` (`id` bigint NOT NULL AUTO_INCREMENT, `code` char(20) NOT NULL, `user_id` char(26) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `authcode_user_id_code` (`user_id`, `code`), CONSTRAINT `auth_codes_users_auth_codes` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
