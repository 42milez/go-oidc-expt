# go-oidc-server

`go-oidc-server` is an experimental implementation of [OpenID Connect](https://openid.net/connect/) with Go.

## Features

TBD

## TRY FOR NOW :zap:

TBD ( Use the image upload to Docker Hub )

## Setup for development

### Install required packages

```
brew install ariga/tap/atlas docker-buildx docker-compose golangci-lint lima openssl@3
```

### Enable docker-buildx

```
mkdir -p ~/.docker/cli-plugins
ln -sfn /opt/homebrew/opt/docker-buildx/bin/docker-buildx ~/.docker/cli-plugins/docker-buildx
```

### Generate key pair for signing access token

```
mkdir -p app/idp/cert
openssl ecparam -genkey -name prime256v1 -noout -out app/idp/auth/cert/secret.pem
openssl ec -in app/idp/auth/cert/secret.pem -pubout -out app/idp/auth/cert/public.pem
```

References:

- [Generating an Elliptic Curve keys](https://cloud.google.com/iot/docs/how-tos/credentials/keys#generating_an_elliptic_curve_keys)

### Create and start the virtual machine

```
limactl start --name=go-oidc-server lima.yml
```

The virtual machine:

  - can be stopped with `limactl stop go-oidc-server`
  - can be deleted with `limactl delete go-oidc-server`

### Create and switch docker context:

```
docker context create lima-go-oidc-server --docker "host=unix:///${HOME}/.lima/go-oidc-server/sock/docker.sock"
docker context use lima-go-oidc-server
```

### Run containers

```
make up
```

The containers:

  - can be stopped with `make stop`
    - The stopped containers can be started with `make start`
  - can be stopped and removed with `make down`

### Apply migrations

```
./script/atlas/migrate-apply.sh
```

## Commands and scripts

The commands described later require the following parameters:

| Parameter      | Detail                                                                                                                 |
|----------------|------------------------------------------------------------------------------------------------------------------------|
| APP_NAME       | Application name ( `idp` or `auth` )                                                                                   |
| MIGRATION_NAME | A part of migration file name. The filename is determined according to the format `%Y%m%d%H%i%S_<MIGRATION_NAME>.sql`. |
| N_LATEST       | The number of latest migration files to be analyzed. `migrate-list.sh` runs analysis on them.                          |

### Generating assets

```
go generate ./...
```

### Generating versioned migration files

```
./script/atlas/migrate-diff.sh <APP_NAME> <MIGRATION_NAME>
```

### Verifying and linting migrations

```
./script/atlas/migrate-lint.sh <APP_NAME> [<N_LATEST>]
```

If `N_LATEST` isn't specified, the diff between `main` branch and the current one is selected as the changeset.

## References

- OpenID Connect
  - OpenID Connect Core 1.0 incorporating errata set 1
    - [EN](https://openid.net/specs/openid-connect-core-1_0.html)
    - [JA](https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html)
- OTP
  - [RFC4226: An HMAC-Based One-Time Password Algorithm](https://www.rfc-editor.org/rfc/rfc4226)
  - [RFC6238: Time-Based One-Time Password Algorithm](https://www.rfc-editor.org/rfc/rfc6238)
  - [Key Uri Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
- Password
  - [Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- ULID
  - [Universally Unique Lexicographically Sortable Identifier](https://github.com/ulid/spec)

## Notes

- [TOTP Base32 vs Base64](https://stackoverflow.com/questions/50082075/totp-base32-vs-base64)
