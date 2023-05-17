Migrate Up
command : make migrateup
migrate -path database/migration -database "postgresql://postgres:yourpassword@localhost:5432/fitcare?sslmode=disable" -verbose up 

Migrate Down
command : make migratedown


root - username
secret - password
fitcare - database name