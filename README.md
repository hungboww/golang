### golang
JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.
It is easy to identify web application users through sessions, however, when your web apps API is interacting with say an Android or IOS client,
sessions becomes unusable because of the stateless nature of the http request. With JWT, we can create a unique token for each authenticated user,
this token would be included in the header of the subsequent request made to the API server, this method allow us to identify every
users that make calls to our API.

### build docker-compose and run server:
```bash
docker-compose up --build

```
### run migrate
--> up
```bash
docker run -v folder/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://root:root@hostport/db?sslmode=disable" up

```
--> down
```bash
docker run -v folder/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://root:root@hostport/db?sslmode=disable" down

```


