# schemalex

Generate difference sql of two mysql schema.

https://github.com/schemalex/schemalex

This allows us to see and automate the SQL Migrations once we take a SQL Dump from the current DB, and compare it against the SQL dump in the previous DB Migration folder.


git schemalex -schema path/to/schema -dsn "$root:$passowrd@/$database" -workspace /path/to/git/repository -deploy

