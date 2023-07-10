-- Modify "admins" table
ALTER TABLE `admins` MODIFY COLUMN `password_hash` varchar(1000) NOT NULL;
