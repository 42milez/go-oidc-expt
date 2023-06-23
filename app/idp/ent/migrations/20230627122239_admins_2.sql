-- Modify "admins" table
ALTER TABLE `admins` MODIFY COLUMN `password_hash` varchar(751) NOT NULL;
