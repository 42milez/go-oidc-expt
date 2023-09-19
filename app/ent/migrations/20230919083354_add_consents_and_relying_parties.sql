-- Create "relying_parties" table
CREATE TABLE `relying_parties` (`id` bigint NOT NULL AUTO_INCREMENT, `client_id` varchar(255) NOT NULL, `client_secret` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `client_id` (`client_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "auth_codes" table
ALTER TABLE `auth_codes` ADD COLUMN `used_at` timestamp NOT NULL, ADD COLUMN `client_id` bigint NULL, ADD INDEX `auth_codes_relying_parties_auth_codes` (`client_id`), DROP FOREIGN KEY `auth_codes_users_auth_codes`, ADD CONSTRAINT `auth_codes_relying_parties_auth_codes` FOREIGN KEY (`client_id`) REFERENCES `relying_parties` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
-- Create "consents" table
CREATE TABLE `consents` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` bigint unsigned NOT NULL, `client_id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `consent_user_id_client_id` (`user_id`, `client_id`), CONSTRAINT `consents_users_consents` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
