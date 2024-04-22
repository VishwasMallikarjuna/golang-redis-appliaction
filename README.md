# golang-redis-appliaction

Here's a simple Go application that connects to PostgreSQL and Redis to perform basic CRUD operations on the users table:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

Make sure you have Redis installed and running. If not, you can download and install it from the Redis website.
