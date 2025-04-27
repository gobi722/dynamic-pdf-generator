package main

import (
	"context"
	"log"
	"os"
	"pdf_generate/pdf_generate"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection details
const (
	numWorkers = 10 // Number of parallel workers
)

var client *mongo.Client

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file " + err.Error())
	}
	// Create output directory if not exists
	if _, err = os.Stat(os.Getenv("PDF_FOLDER")); os.IsNotExist(err) {
		err = os.Mkdir(os.Getenv("PDF_FOLDER"), os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	// Connect to MongoDB
	url := os.Getenv("MONGODB_CONNECTION_STRING")
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	app := fiber.New()

	// Setup route
	app.Post("/pdf", Pdfhandler)

	// Start server
	log.Fatal(app.Listen(":9900"))

}
func Pdfhandler(c *fiber.Ctx) error {
	collection := client.Database(os.Getenv("MONGODB_NAME")).Collection(os.Getenv("MONGODB_COLLECTION"))

	// Fetch all Invoices
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to fetch data from MongoDB: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Channel to send Invoices for PDF generation
	docChan := make(chan pdf_generate.Invoice, numWorkers)
	wg := &sync.WaitGroup{}

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go pdfWorker(docChan, wg)
	}

	// Send Invoices to workers
	for cursor.Next(context.TODO()) {
		var doc pdf_generate.Invoice
		if err := cursor.Decode(&doc); err != nil {
			log.Printf("Failed to decode Invoice: %v", err)
			continue
		}
		docChan <- doc
	}

	// Close the channel and wait for workers to finish
	close(docChan)
	wg.Wait()

	log.Println("PDF generation completed")
	return nil
}

// Worker function to generate PDFs
func pdfWorker(docChan <-chan pdf_generate.Invoice, wg *sync.WaitGroup) {
	defer wg.Done()

	for doc := range docChan {
		pdf_generate.GenerateInvoicePDF(doc)
		// err :=
		// if err != nil {
		// 	log.Printf("Failed to generate PDF for %s: %v", doc.ID, err)
		// } else {
		// 	log.Printf("PDF generated for %s", doc.ID)
		// }
	}
}
