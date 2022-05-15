# Dynamic SQL query with Go
This is a simple project to show how to use Go standard library to do dynamic (unknown column) SQL query.

All the explanations are in source code **main.go** file.

This example using CockroachDB as backend database,<br/>
which and be launched in **docker-compose** folder.

### Run the example
> Example below only support Linux
1. go to docker-compse folder and launch CockroachDB container:
  ```bash
  cd docker-compose
  docker compose up -d
  ```
2. under the same folder, run below script to create testing data:
  ```bash
  docker compose exec cockroachdb cockroach sql --insecure \
  --execute "$(cat init.sql)"
  ```
3. back to project root folder and test query:
  ```bash
  cd ..
  go run main.go "SELECT * FROM users"
  ```
    result will be:
  ```
  id                                   name   age married location                                             phone_numbers       creation_time
4572589b-6295-4880-b1b0-c45849524193 Nathan 25  true    {"city": "Salt Lake City", "state": "UT"}            {987-6666}          2022-05-15T14:27:22.956924Z
9a9ebdb5-cfba-4efa-824c-aad81fa7ad75 Tom    31  true    {"city": "LA", "state": "CA"}                        {000-0000,001-1234} 2022-05-15T14:27:22.956924Z
ec9c6e0b-5598-45b0-a0e6-819ffcb44036 Mary       false   {"city": "Phoenix", "state": "AZ", "street": "Main"}                     2022-05-15T14:27:22.956924Z
  ```
4. we can select specific columns and rename it:
  ```bash
  go run main.go \
  "SELECT name, location->>'city' AS \"CITY\" FROM users"
  ```
    result will be:
  ```
  name   CITY
Nathan Salt Lake City
Tom    LA
Mary   Phoenix
  ```
