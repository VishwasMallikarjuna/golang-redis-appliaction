package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

// User represents a user in the database
type User struct {
	ID    int
	Name  string
	Email string
}

var db *sql.DB
var rdb *redis.Client
var ctx = context.Background() // Context variable

func init() {
	// Connect to PostgreSQL
	var err error
	db, err = sql.Open("postgres", "postgres://yourusername:yourpassword@localhost/simple_app_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	defer db.Close()
	defer rdb.Close()

	// Create a user
	userID, err := createUser("Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created user with ID: %d\n", userID)

	// Get a user by ID (with caching)
	user, err := getUserWithCache(userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved user: %+v\n", user)

	// Update user's email
	err = updateUserEmail(userID, "alice.updated@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user's email")

	// Delete the user
	err = deleteUser(userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user")
}

func getUserWithCache(userID int) (*User, error) {
	// Check if the user exists in the cache
	cacheKey := fmt.Sprintf("user:%d", userID)
	cachedUserJSON, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// User found in cache, unmarshal JSON into User object and return
		var user User
		if err := json.Unmarshal([]byte(cachedUserJSON), &user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	// User not found in cache, fetch from database
	user, err := getUser(userID)
	if err != nil {
		return nil, err
	}

	// Marshal User object to JSON and set in cache with expiration time
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err := rdb.Set(ctx, cacheKey, userJSON, 24*time.Hour).Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func createUser(name, email string) (int, error) {
	var userID int
	err := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", name, email).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func getUser(userID int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func updateUserEmail(userID int, newEmail string) error {
	_, err := db.Exec("UPDATE users SET email = $1 WHERE id = $2", newEmail, userID)
	if err != nil {
		return err
	}
	return nil
}

func deleteUser(userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
