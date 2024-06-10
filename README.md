# go-oidc-server

`go-oidc-server` is an experimental implementation of [OpenID Connect](https://openid.net/connect/) with Go.

## Features

### Supported authentication path

- [x] Authorization Code Flow
- [ ] Implicit Flow
- [ ] Hybrid Flow

## Setup for development

### Install required packages

```
./scripts/bootstrap/brew.sh
```

### Generate key pair for signing access token

```
./scripts/bootstrap/keypair.sh
```

### Generate keys for encrypting/decrypting cookie

```
./scripts/bootstrap/key.sh
```

The script creates a key pair in `app/pkg/xjwt/cert`.

References:

- [Generating an Elliptic Curve keys](https://cloud.google.com/iot/docs/how-tos/credentials/keys#generating_an_elliptic_curve_keys)

### Generate certificates for load balancer

For the conformance test of OpenID connect. 

```
mkcert -install
mkcert host.docker.internal
cat host.docker.internal.pem > host.docker.internal-fullchain.pem
cat "$(mkcert -CAROOT)/rootCA.pem" >> host.docker.internal-fullchain.pem
openssl dhparam -out dhparam.pem 2048
mv *.pem docker/load-balancer/etc/nginx/ssl
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
make migrate-apply SERVICE=db DATABASE=idp
make migrate-apply SERVICE=db DATABASE=idp_test
```

## Commands and scripts

`make` supports the following commands:

| Command       | Detail                                  |
|---------------|-----------------------------------------|
| build         | Build a docker image to deploy          |
| build-local   | Build docker images                     |
| benchmark     | Run all benchmarks                      |
| cleanup-db    | Clean up database                       |
| cleanup-go    | Clean up caches                         |
| fmt           | Run formatter                           |
| gen           | Run generator                           |
| lint          | Run linters                             |
| migrate-apply | Apply migrations                        |
| migrate-diff  | Generate migrations                     |
| migrate-lint  | Run analysis on the migration directory |
| resolve       | Resolve dependencies                    |
| seed          | Seeding database                        |
| test          | Run all tests                           |
| up            | Create and start containers             |
| down          | Stop and remove containers              |
| start         | Start containers                        |
| stop          | Stop containers                         |
| destroy       | Delete all resources                    |

### Generating database schema

e.g. The following command generates `AuthCode` schema.

```
go run -mod=mod entgo.io/ent/cmd/ent new --target app/pkg/ent/schema AuthCode
```

### Generating versioned a migration file

```
make migrate-diff MIGRATION_NAME=<MIGRATION_NAME>
```

| Parameter      | Detail                                                                                                                 |
|----------------|------------------------------------------------------------------------------------------------------------------------|
| MIGRATION_NAME | A part of migration file name. The filename is determined according to the format `%Y%m%d%H%i%S_<MIGRATION_NAME>.sql`. |

### Verifying and linting migrations

```
make migrate-lint [N_LATEST=<N_LATEST>]
```

| Parameter      | Detail                                                                                                                 |
|----------------|------------------------------------------------------------------------------------------------------------------------|
| N_LATEST       | The number of latest migration files to be analyzed. `migrate-list.sh` runs analysis on them.                          |

If `N_LATEST` isn't specified, the diff between `main` branch and the current one is selected as the changeset.

### Applying migrations

```
make migrate-apply DB_NAMES=idp,idp_test
```

| Parameter | Detail                                          |
|-----------|-------------------------------------------------|
| DB_NAMES  | Database names that will be applied migrations. |

### Seeding database

```
make seed
```

## Documents

### Swagger

API specification (Swagger UI) is available on `http://localhost:8880`. Before accessing the URL, it needs to run the following command to start `swagger-ui` container.

```
docker-compose up -d swagger-ui
```

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
  - [TOTP Base32 vs Base64](https://stackoverflow.com/questions/50082075/totp-base32-vs-base64)
- Password
  - [Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- Session
  - [Session Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html)

## Miscellaneous

- Computer Science
  - Memory Model
    - [The Go Memory Model](https://go.dev/ref/mem)
    - [Memory Models](https://research.swtch.com/mm)
      - [Memory barrier](https://en.wikipedia.org/wiki/Memory_barrier)
      - [Out-of-order execution](https://en.wikipedia.org/wiki/Out-of-order_execution)
  - Synchronization
    - [Lock-free/Wait-free algorithm](https://ja.wikipedia.org/wiki/Lock-free%E3%81%A8Wait-free%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0)
    - [Semaphore](https://en.wikipedia.org/wiki/Semaphore_(programming))
- Others
  - [Go Style Decisions: Naming](https://google.github.io/styleguide/go/decisions#naming)
  - [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)
  - [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
