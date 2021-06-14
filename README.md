# shops_categories_ms
Microservice which will get all categories, attributes and attributes values. This microservice will save it into structured DB.

## Changelog
**0.0.1**
- Added config.example.json file which can be used like a template
- Created Makefile
- Added CI script
- Created Docker-compose file

## Config
If you need get something from `config.json`, please use config package. Don't get directly from `config.json` file!
If you want to add some into `config.json`, please don't forget to adjust `config/config.go` file.

## Migrations
Pkg: https://github.com/rubenv/sql-migrate

For running migration, type `make migrate` and the script will migrate all files which is not exists in DB table `gorp_migrations`.
If you want to rollback your migrations, type `make rollback`. **Be careful!** this command will remove ALL your migrations which registered in migrations table. 