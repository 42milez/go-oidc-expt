Here is a sequence of Authorization Code flow:

```mermaid
sequenceDiagram
  participant EndUser as End-User
  participant RP as Relying Party
  participant IdP as Identity Provider
  participant DB as Database
  participant Cache as Cache

  EndUser ->>+ RP: A request that needs credentials issued by Identity Provider
  RP -->>- EndUser: Redirect to authorization endpoint

  alt End-user is authenticated / Permission is grantted
    EndUser ->>+ IdP: Authorization Request: GET /authorization
    IdP -->>- EndUser: Redirect to callback endpoint
  else End-user isn't authenticated / Permission isn't grantted 
    EndUser ->>+ IdP: Authorization Request: GET /authorization
    IdP -->> EndUser: Redirect to authentication endpoint
    EndUser ->> IdP: Authentication Request: POST /authentication
    IdP -->> EndUser: Redirect to consent endpoint
    EndUser ->> IdP: Grant permission: POST /consent
    IdP -->> EndUser: Redirect to authorization endpoint
    EndUser ->> IdP: Authorization Request: GET /authorization
    IdP -->>- EndUser: Redirect to callback endpoint
  end

  EndUser ->>+ RP: Pass authorization code: GET /callback
  
  alt Get tokens
    RP ->>+ IdP: Token Request: POST /token with grant_type=authorization_code
    IdP -->>- RP: Respond tokens
  else Refresh access token
    RP ->>+ IdP: Token Request: POST /token with grant_type=refresh_token
    IdP -->>- RP: Respond tokens
  end

  RP -->>- EndUser: Respond tokens
```
