#!/bin/bash
set -e


#!Set global variables
CONTAINER_NAME="postgres-test"
IMAGE_NAME="postgres"
SQL_HOST=6432
SQL_PORT=5432


#! Check if the container already exists
if docker container inspect "$CONTAINER_NAME" >/dev/null 2>&1; then
    echo "Container $CONTAINER_NAME already exists. Skipping..."
    winpty docker start $CONTAINER_NAME
else
    # Run the container
    winpty docker run --name "$CONTAINER_NAME" -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123321 -p "$SQL_HOST":"$SQL_PORT" -d "$IMAGE_NAME"
fi
sleep 3


#!If you are using windows this command probably doesn't work' => This command check health postgress
if [ "$IMAGE_NAME" = "postgres" ]; then
    echo "Waiting for postgres..."

    while ! netstat -ano | findstr LISTENING | findstr ":$SQL_PORT "; do
        sleep 0.1
    done

    echo "PostgreSQL started..."
fi
sleep 3


#!Create databse
winpty docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE eventapp"
sleep 3
echo "Database for eventapp created"


#!Create table for created database
winpty docker exec -it postgres-test psql -U postgres -d eventapp -c "
    CREATE TABLE IF NOT EXISTS event (
        id bigserial not null primary key,
        name varchar(255) not null,
        description varchar(255) not null,
        location varchar(255) not null,
        dateTime timestamp not null,
        user_id integer
    );
"
sleep 3
echo "Events table created successfully"

# docker exec -it postgres-test psql -U postgres -d eventapp -c "SELECT * FROM event"
# docker exec -it postgres-test psql -U postgres -d eventapp -c "TRUNCATE event RESTART IDENTITY"
# docker ps --filter "status=exited"