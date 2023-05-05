# This is an example of web application written in Go

## Libraries

- https://github.com/alexedwards/scs/postgresstore - Postgres Store for session mananger
- https://github.com/alexedwards/scs/v2 - Session Manager
- https://github.com/go-playground/form/v4 - HTTP Form Binder
- https://github.com/golang-migrate/migrate/v4 - Database migration tool
- https://github.com/jmoiron/sqlx - Database util library for cleaner and easier database operations
- https://github.com/julienschmidt/httprouter - HTTP Routing
- https://github.com/justinas/alice - HTTP Middleware chaining
- https://github.com/lib/pq - Postgres database driver
- https://github.com/justinas/nosurf - CSRF Token management
- https://golang.org/x/crypto - Password hashing

## TLS Certificates

For local development you can generate the certificates using following command:

> go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

Include generated files into `tls/` folder of the root directory
