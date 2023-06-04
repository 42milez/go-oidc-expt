# go-oidc-server

`go-oidc-server` is an experimental implementation of [OpenID Connect](https://openid.net/connect/) with Go.

## Features

TBD

## TRY FOR NOW :zap:

### Setup

#### Install required packages:

```
brew install docker-buildx docker-compose golangci-lint lima
```

#### Enable `docker-buildx`:

```
mkdir -p ~/.docker/cli-plugins
ln -sfn /opt/homebrew/opt/docker-buildx/bin/docker-buildx ~/.docker/cli-plugins/docker-buildx
```

#### Create and start the virtual machine:

```
limactl start --tty=false --name=go-oidc-server lima.yml
```

- Note:
  - The VM:
    - can be stopped with `limactl stop go-oidc-server`
    - can be deleted with `limactl delete go-oidc-server`

#### Create and switch docker context:

```
docker context create lima-go-oidc-server --docker "host=unix:///${HOME}/.lima/go-oidc-server/sock/docker.sock"
docker context use lima-go-oidc-server
```

#### Create and run containers

```
make up
```

Note:

- The containers:
  - can be stopped with `make stop`
    - The stopped containers can be started with `make start`
  - can be stopped and removed with `make down`

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

## Notes

- [TOTP Base32 vs Base64](https://stackoverflow.com/questions/50082075/totp-base32-vs-base64)
