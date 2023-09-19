-- Modify "auth_codes" table
ALTER TABLE `auth_codes` MODIFY COLUMN `code` char(30) NOT NULL;
-- Modify "consents" table
ALTER TABLE `consents` DROP COLUMN `client_id`, ADD COLUMN `relying_party_id` bigint NOT NULL, DROP INDEX `consent_user_id_client_id`, ADD INDEX `consent_user_id_relying_party_id` (`user_id`, `relying_party_id`);
