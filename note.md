# Session
Browser
   |
   | login
   v
Server
   |
   +--> simpan session di database/redis

# With JWT
Browser
   |
   | login
   v
Server
   |
   +--> generate token
   |
   +--> kirim token ke client

## Apa isi JWT

JWT terdiri dari:
```
header.payload.signature
value :
xxxxx.yyyyy.zzzzz
```
header jwt berisi : 
```.json
{
  "alg": "HS256", //-> algoritma sigin
  "typ": "JWT" //-> tipe
}
```
payload(claims)
berisi data user
```
{
  "user_id": "123",
  "type": "access",
  "exp": 1750000000
}
```
signature jwt terdiri dari : header + payload + secret
jwt lib membuat hash dengan :
```
HMACSHA256(
    header.payload,
    secret
)
```