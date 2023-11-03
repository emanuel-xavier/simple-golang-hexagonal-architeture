# This is a simple attempt to use hexagonal arichiteture in a golang api

This is a simple API that can create and list some users inside a database.

## Routes

Search for a user using its id.
| METHOD | ROUTE          | DESCRIPTION                    |
|--------|----------------|--------------------------------|
| GET    | <ip>/user/:id  | Search for a user using its ID |
| GET    | <ip>/user/list | List all users in the database |
| POST   | <ip>/user      | Create a new user              |
