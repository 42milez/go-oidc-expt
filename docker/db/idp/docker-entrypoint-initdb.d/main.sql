CREATE DATABASE IF NOT EXISTS idp;
CREATE DATABASE IF NOT EXISTS idp_test;
CREATE DATABASE IF NOT EXISTS atlas;

CREATE USER IF NOT EXISTS `idp`@`%` IDENTIFIED BY 'idp';
CREATE USER IF NOT EXISTS `idp_test`@`%` IDENTIFIED BY 'idp_test';
CREATE USER IF NOT EXISTS `atlas`@`%` IDENTIFIED BY 'atlas';

-- --------------------------------------------------
--  Permissions for 'idp'
-- --------------------------------------------------

GRANT
  DELETE,
  INSERT,
  SELECT,
  UPDATE
ON idp.* TO 'idp'@'%';

-- --------------------------------------------------
--  Permissions for 'idp_test'
-- --------------------------------------------------

GRANT
  DELETE,
  INSERT,
  SELECT,
  UPDATE
ON idp_test.* TO 'idp_test'@'%';

-- --------------------------------------------------
--  Permissions for 'atlas'
-- --------------------------------------------------

GRANT
  CREATE,
  DROP,
  REFERENCES,
  INDEX,
  ALTER
ON atlas.* TO 'atlas'@'%';

GRANT
  CREATE,
  DROP,
  REFERENCES,
  INDEX,
  ALTER
ON `atlas\_dev\_%`.* TO 'atlas'@'%';

GRANT
  CREATE,
  DROP,
  REFERENCES,
  INDEX,
  ALTER
ON idp.* TO 'atlas'@'%';

GRANT
  SELECT,
  INSERT,
  UPDATE,
  CREATE
ON idp.atlas_schema_revisions TO 'atlas'@'%';

GRANT
  CREATE,
  DROP,
  REFERENCES,
  INDEX,
  ALTER,
  SELECT
ON idp_test.* TO 'atlas'@'%';

GRANT
  SELECT,
  INSERT,
  UPDATE,
  CREATE
ON idp_test.atlas_schema_revisions TO 'atlas'@'%';
