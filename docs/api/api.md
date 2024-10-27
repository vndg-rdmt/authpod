


### Sign in to service

`POST /api/sign-in`

Request 
```json
{
    "login": "string",
    "password": "string"
}
```

Response
`200`
```json
{
    "session_id": "string",
    "expires_at": "date formatted string"
}
```

`401` is returned when credentials are not valid


### Check session

`GET /api/ping`

request

```http
Authortization: <session_id>
```

response
`200`
```json
{
    "user_id": 8
}
```

`401` is returned when session is invalid

### Issue new api token for user

`PUT /api/tokens/:user_id`

response
`200`

```json
{
    "token": "string",
    "expires_at": "date formatted string",
}
```

### Check api token for user

`GET /api/tokens/:token`

response
`200`
```json
{
    "user_id": 8
}
```
