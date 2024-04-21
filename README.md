# golang-redis-appliaction

Here's a simple Go application that connects to PostgreSQL and Redis to perform basic CRUD operations on the users table:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
