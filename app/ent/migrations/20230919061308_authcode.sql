-- Modify "auth_codes" table
ALTER TABLE `auth_codes` ADD INDEX `auth_codes_users_auth_codes` (`user_id`), ADD UNIQUE INDEX `authcode_code` (`code`);
