-- Create "relying_parties" table
CREATE TABLE `relying_parties` (`id` bigint NOT NULL AUTO_INCREMENT, `client_id` varchar(255) NOT NULL, `client_secret` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `client_id` (`client_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "auth_codes" table
CREATE TABLE `auth_codes` (`id` bigint NOT NULL AUTO_INCREMENT, `code` char(10) NOT NULL, `expire_at` timestamp NOT NULL, `created_at` timestamp NOT NULL, `used_at` timestamp NOT NULL, `relying_party_id` bigint NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `authcode_relying_party_id_code` (`relying_party_id`, `code`), CONSTRAINT `auth_codes_relying_parties_auth_codes` FOREIGN KEY (`relying_party_id`) REFERENCES `relying_parties` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint unsigned NOT NULL, `name` varchar(255) NOT NULL, `password` varchar(284) NOT NULL, `totp_secret` char(160) NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "consents" table
CREATE TABLE `consents` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` bigint unsigned NOT NULL, `client_id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `consent_user_id_client_id` (`user_id`, `client_id`), CONSTRAINT `consents_users_consents` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "redirect_uris" table
CREATE TABLE `redirect_uris` (`id` bigint NOT NULL AUTO_INCREMENT, `uri` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, `relying_party_id` bigint NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `redirecturi_relying_party_id_uri` (`relying_party_id`, `uri`), CONSTRAINT `redirect_uris_relying_parties_redirect_uris` FOREIGN KEY (`relying_party_id`) REFERENCES `relying_parties` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
