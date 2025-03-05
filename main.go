package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nbursa/whistlechain-backend/blockchain"
	"github.com/nbursa/whistlechain-backend/storage"
)

// Global Report Store
var reportStore = storage.NewReportStore()

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
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
		EncryptedDescription string `json:"encryptedDescription"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Store encrypted report (do NOT decrypt on the backend)
	encryptedReport := req.EncryptedDescription

	// Generate a hash for blockchain verification
	reportHash := blockchain.GenerateHash(encryptedReport)

	// TODO: Store `encryptedReport` in the database or IPFS
	// TODO: Store `reportHash` on the blockchain

	return c.JSON(fiber.Map{
		"message": "Report submitted successfully",
		"hash":    reportHash,
	})
}


// GetAllReports returns all reports
func GetAllReports(c *fiber.Ctx) error {
	reports := reportStore.GetAllReports()
	return c.JSON(reports)
}

// GetReportByID returns a specific report
func GetReportByID(c *fiber.Ctx) error {
	id := c.Params("id")
	report, exists := reportStore.GetReportByID(id)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	return c.JSON(report)
}
