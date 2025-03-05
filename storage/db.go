package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// ConnectDB initializes PostgreSQL connection
func ConnectDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL not set. Make sure the Heroku database is configured correctly.")
	}

	// Ensure SSL Mode is required for Heroku
	sslDatabaseURL := fmt.Sprintf("%s?sslmode=require", databaseURL)

	log.Println("🔍 Connecting to PostgreSQL with SSL...")
	var err error
	DB, err = sqlx.Open("postgres", sslDatabaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to open DB connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL (SSL Enabled)")
}
