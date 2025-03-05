package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nbursa/whistlechain-backend/blockchain"
	"github.com/nbursa/whistlechain-backend/storage"
)

// Report struct with explicit database column mapping
type Report struct {
	ID             int    `db:"id" json:"id"`
	CompanyID      int    `db:"company_id" json:"company_id"`
	EncryptedData  string `db:"encrypted_data" json:"encrypted_data"`
	BlockchainHash string `db:"blockchain_hash" json:"blockchain_hash"`
	CreatedAt      string `db:"created_at" json:"created_at"`
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Connect to Database
	storage.ConnectDB()

	// Ensure tables exist
	ensureTables()

	// Initialize Fiber app
	app := fiber.New()

	// API Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ WhistleChain Backend is running!")
	})	
	app.Post("/report", SubmitReport)
	app.Get("/reports", GetAllReports)
	app.Get("/report/:id", GetReportByID)

	// Fix: Use Heroku-assigned PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default for local testing
	}

	log.Println("🚀 Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}

// ensureTables creates the reports table if it doesn't exist
func ensureTables() {
	query := `
	CREATE TABLE IF NOT EXISTS reports (
		id SERIAL PRIMARY KEY,
		company_id INT NOT NULL,
		encrypted_data TEXT NOT NULL,
		blockchain_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := storage.DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Failed to create reports table: %v", err)
	}
	log.Println("✅ Tables ensured to exist")
}

// SubmitReport handles report submissions
func SubmitReport(c *fiber.Ctx) error {
	type Request struct {
		CompanyID           int    `json:"company_id"`
		EncryptedDescription string `json:"encryptedDescription"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Ensure the database is connected
	if storage.DB == nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection not initialized"})
	}

	// Store encrypted report in PostgreSQL
	_, err := storage.DB.Exec(`
        INSERT INTO reports (company_id, encrypted_data, blockchain_hash)
        VALUES ($1, $2, $3)
    `, req.CompanyID, req.EncryptedDescription, blockchain.GenerateHash(req.EncryptedDescription))

	if err != nil {
		log.Printf("Database error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database insert failed"})
	}

	return c.JSON(fiber.Map{
		"message": "Report submitted successfully",
	})
}

// GetAllReports returns all reports from PostgreSQL
func GetAllReports(c *fiber.Ctx) error {
	var reports []Report

	err := storage.DB.Select(&reports, "SELECT * FROM reports")
	if err != nil {
		log.Printf("Database query error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch reports"})
	}

	return c.JSON(reports)
}

// GetReportByID returns a specific report from PostgreSQL
func GetReportByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var report Report

	err := storage.DB.Get(&report, "SELECT * FROM reports WHERE id = $1", id)
	if err != nil {
		log.Printf("Database query error: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "Report not found"})
	}

	return c.JSON(report)
}
