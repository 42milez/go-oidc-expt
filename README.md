# go-oidc-server

`go-oidc-server` is an experimental implementation of [OpenID Connect](https://openid.net/connect/) with Go.

## Features

TBD

## TRY FOR NOW :zap:

TBD ( Use the image upload to Docker Hub )

## Setup for development

### Install required packages

```
./script/bootstrap/brew.sh
```

### Enable docker-buildx

```
./script/bootstrap/docker.sh
```

### Generate key pair for signing access token

```
./script/bootstrap/keypair.sh
```

The script creates a key pair in `app/pkg/xjwt/cert`.

References:

- [Generating an Elliptic Curve keys](https://cloud.google.com/iot/docs/how-tos/credentials/keys#generating_an_elliptic_curve_keys)

### Create and start the virtual machine thant runs docker containers

```
make lc-create
make lc-start
```

The virtual machine:
- can be stopped with `make lc-stop`
- can be deleted with `make lc-delete`

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
  - Stopped containers can be started with `make start`
- can be stopped and removed with `make down`
- and volumes can be deleted with `make destroy`

### Apply migrations

```
./script/atlas/migrate-apply.sh
```

## Commands and scripts

The commands described later require the following parameters:

| Parameter      | Detail                                                                                                                 |
|----------------|------------------------------------------------------------------------------------------------------------------------|
| MIGRATION_NAME | A part of migration file name. The filename is determined according to the format `%Y%m%d%H%i%S_<MIGRATION_NAME>.sql`. |
| N_LATEST       | The number of latest migration files to be analyzed. `migrate-list.sh` runs analysis on them.                          |

### Generating assets

```
make gen
```

### Generating database schema

e.g. The following command generates `AuthCode` schema.

```
go run -mod=mod entgo.io/ent/cmd/ent new --target app/ent/schema AuthCode
```

### Generating versioned a migration file

```
make migrate-diff MIGRATION_NAME=<MIGRATION_NAME>
```

### Verifying and linting migrations

```
make migrate-lint [N_LATEST=<N_LATEST>]
```

If `N_LATEST` isn't specified, the diff between `main` branch and the current one is selected as the changeset.

### Applying migrations

```
make migrate-apply
```

### Seeding database

```
make seed
```

## Documents

- API specification (Swagger UI) is available on `http://localhost:8080`.

## References

- Identifier
  - [Awesome Identifiers](https://github.com/adileo/awesome-identifiers)
  - [sonyflake](https://github.com/sony/sonyflake)
  - [Universally Unique Lexicographically Sortable Identifier](https://github.com/ulid/spec)
- OpenID Connect
  - OpenID Connect Core 1.0 incorporating errata set 1
    - [English](https://openid.net/specs/openid-connect-core-1_0.html)
    - [Japanese](https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html)
- OpenAPI
  - [OpenAPI.Tools](https://openapi.tools/)
- OTP
  - [RFC4226: An HMAC-Based One-Time Password Algorithm](https://www.rfc-editor.org/rfc/rfc4226)
  - [RFC6238: Time-Based One-Time Password Algorithm](https://www.rfc-editor.org/rfc/rfc6238)
  - [Key Uri Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
- Password
  - [Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- Session
  - [Session Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html)

## Notes

- [TOTP Base32 vs Base64](https://stackoverflow.com/questions/50082075/totp-base32-vs-base64)
