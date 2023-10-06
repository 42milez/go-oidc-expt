-- Modify "relying_parties" table
ALTER TABLE `relying_parties` ADD UNIQUE INDEX `relyingparty_client_id_client_secret` (`client_id`, `client_secret`);
