-- Modify "consents" table
ALTER TABLE `consents` DROP COLUMN `relying_party_id`, ADD COLUMN `client_id` varchar(255) NOT NULL, DROP INDEX `consent_user_consents_relying_party_id`, ADD INDEX `consent_user_consents_client_id` (`user_consents`, `client_id`);
