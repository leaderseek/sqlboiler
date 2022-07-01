# sqlboiler
Contains the SQL schema definition and associated boilerplate compiled for Go.

## Setup DB
- Run `docker compose up`
- Run `psql -U user1 -f .\sql\schema.sql`
- Enter password: pass1

## Setup SQLBoiler
- Run `go install github.com/volatiletech/sqlboiler/v4@latest`.
- Run `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`.
- Check `.\sqlboiler.toml` matches with your database. [See here for help.](https://github.com/volatiletech/sqlboiler#configuration)

## Update Schema
- Edit `.\sql\schema.sql`
- [Setup DB](#setup-db)
- [Setup SQLBoiler](#setup-sqlboiler)
- Delete `.\repository`
- Run `sqlboiler psql`
