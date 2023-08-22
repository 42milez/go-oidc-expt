-- Create "redirect_uris" table
CREATE TABLE `redirect_uris` (`id` char(26) NOT NULL, `name` varchar(30) NOT NULL, `password_hash` varchar(1000) NOT NULL, `totp_secret` char(160) NULL, `created_at` timestamp NOT NULL, `modified_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "redirect_ur_is" table
ALTER TABLE `redirect_ur_is` DROP FOREIGN KEY `redirect_ur_is_users_redirect_uris`, ADD CONSTRAINT `redirect_ur_is_redirect_uris_redirect_uris` FOREIGN KEY (`user_id`) REFERENCES `redirect_uris` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "auth_codes" table
ALTER TABLE `auth_codes` DROP FOREIGN KEY `auth_codes_users_auth_codes`, ADD CONSTRAINT `auth_codes_redirect_uris_auth_codes` FOREIGN KEY (`user_id`) REFERENCES `redirect_uris` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
