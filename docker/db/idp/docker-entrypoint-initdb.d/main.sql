-- --------------------------------------------------
--  Databases
-- --------------------------------------------------

CREATE DATABASE IF NOT EXISTS idp;
CREATE DATABASE IF NOT EXISTS idp_test;
CREATE DATABASE IF NOT EXISTS atlas;

-- --------------------------------------------------
--  Users
-- --------------------------------------------------

-- 'idp' user
CREATE USER IF NOT EXISTS `idp`@`%` IDENTIFIED BY 'idp';
GRANT DELETE, INSERT, SELECT, UPDATE ON idp.* TO 'idp'@'%';

-- 'idp_test' user
CREATE USER IF NOT EXISTS `idp_test`@`%` IDENTIFIED BY 'idp_test';
GRANT DELETE, INSERT, SELECT, UPDATE ON idp_test.* TO 'idp_test'@'%';

-- 'atlas' user
CREATE USER IF NOT EXISTS `atlas`@`%` IDENTIFIED BY 'atlas';
GRANT ALTER, CREATE, DROP, INDEX, REFERENCES ON atlas.* TO 'atlas'@'%';
GRANT ALTER, CREATE, DROP, INDEX, REFERENCES ON `atlas\_dev\_%`.* TO 'atlas'@'%';
GRANT ALTER, CREATE, DROP, INDEX, REFERENCES ON idp.* TO 'atlas'@'%';
GRANT CREATE, INSERT, SELECT, UPDATE ON idp.atlas_schema_revisions TO 'atlas'@'%';
GRANT ALTER, CREATE, DROP, INDEX, REFERENCES, SELECT ON idp_test.* TO 'atlas'@'%';
GRANT CREATE, INSERT, SELECT, UPDATE ON idp_test.atlas_schema_revisions TO 'atlas'@'%';

-- --------------------------------------------------
--  Initial Data
-- --------------------------------------------------

-- INSERT INTO users (id, name, password_hash, created_at, modified_at) VALUES ('', '', '', '', '');
-- INSERT INTO auth_codes (code, expire_at, created_at, user_id) VALUES ('', '', '', '');
-- INSERT INTO auth_codes (code, expire_at, created_at, user_id) VALUES ('', '', '', '');
-- INSERT INTO auth_codes (code, expire_at, created_at, user_id) VALUES ('', '', '', '');
-- INSERT INTO redirect_uris (uri, created_at, user_id) VALUES ('', '', '');
-- INSERT INTO redirect_uris (uri, created_at, user_id) VALUES ('', '', '');
-- INSERT INTO redirect_uris (uri, created_at, user_id) VALUES ('', '', '');
