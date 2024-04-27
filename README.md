# Simple Go Application with PostgreSQL and Redis

This is a simple application written in Go that demonstrates basic CRUD (Create, Read, Update, Delete) operations on a "users" table in PostgreSQL, with Redis used for caching.

## Table of Contents

- [Introduction](#introduction)
- [Setup](#setup)
  - [PostgreSQL Setup](#postgresql-setup)
  - [Redis Setup](#redis-setup)
- [Go Application](#go-application)
  - [Running the Application](#running-the-application)
  - [Configuration](#configuration)
- [Code Overview](#code-overview)
- [Operations](#operations)
- [Notes](#notes)

## Introduction

This application showcases how to build a simple Go application that interacts with PostgreSQL for data storage and Redis for caching. It provides basic CRUD operations on a "users" table, demonstrating how to create, read, update, and delete user records.

## Setup

### PostgreSQL Setup

1. **Install PostgreSQL:** If you haven't already, [install PostgreSQL](https://www.postgresql.org/download/).
2. **Create Database:**
   - Create a new PostgreSQL database named `simple_app_db`.
   - Run the following SQL query to create a `users` table:

   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100),
       email VARCHAR(100)
   );

### Redis Setup
