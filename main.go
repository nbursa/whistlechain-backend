package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nbursa/whistlechain-backend/blockchain"
	"github.com/nbursa/whistlechain-backend/storage"
)

// Main function
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// ✅ Initialize Database Connection
	storage.ConnectDB()

	// Ensure database is connected before starting the server
	if storage.DB == nil {
		log.Fatal("❌ Database connection failed")
	}

	// Initialize Fiber app
	app := fiber.New()

	// API Routes
	app.Post("/report", SubmitReport)
	app.Get("/reports", GetAllReports)
	app.Get("/report/:id", GetReportByID)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
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
	var reports []struct {
		ID                int    `json:"id"`
		CompanyID         int    `json:"company_id"`
		EncryptedData     string `json:"encrypted_data"`
		BlockchainHash    string `json:"blockchain_hash"`
		CreatedAt         string `json:"created_at"`
	}

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

	var report struct {
		ID                int    `json:"id"`
		CompanyID         int    `json:"company_id"`
		EncryptedData     string `json:"encrypted_data"`
		BlockchainHash    string `json:"blockchain_hash"`
		CreatedAt         string `json:"created_at"`
	}

	err := storage.DB.Get(&report, "SELECT * FROM reports WHERE id = $1", id)
	if err != nil {
		log.Printf("Database query error: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "Report not found"})
	}

	return c.JSON(report)
}
