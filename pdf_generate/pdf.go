package pdf_generate

import (
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoicePDF(invoice Invoice) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)
	pdf.AddPage()

	// Add watermark
	watermark := "CONFIDENTIAL"
	pdf.SetFont("Arial", "B", 50)
	pdf.SetTextColor(200, 200, 200)
	pdf.TransformBegin()
	pdf.TransformRotate(45, 105, 150)
	pdf.Text(50, 150, watermark)
	pdf.TransformEnd()

	// Reset text color
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 12)

	// Add header
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "Invoice Statement for "+invoice.To, "0", 1, "C", false, 0, "")
	pdf.Ln(10)

	// Add sender and recipient details
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(100, 10, fmt.Sprintf("From: %s", invoice.From))
	pdf.Cell(100, 10, fmt.Sprintf("To: %s", invoice.To))
	pdf.Ln(6)
	pdf.Cell(100, 10, fmt.Sprintf("From Address: %s", invoice.FromAddress))
	pdf.Cell(100, 10, fmt.Sprintf(" To Address: %s", invoice.ToAddress))
	pdf.Ln(10)

	// Add invoice table headers
	headers := []string{"Item", "Description", "Quantity", "Price", "Total"}
	columnWidths := []float64{40, 70, 20, 30, 30}

	for i, header := range headers {
		pdf.CellFormat(columnWidths[i], 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Add items
	var grandTotal float64
	for _, item := range invoice.Items {
		total := float64(item.Quantity) * item.Price
		grandTotal += total

		pdf.CellFormat(columnWidths[0], 10, item.ID, "1", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidths[1], 10, item.Description, "1", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidths[3], 10, fmt.Sprintf("$%.2f", item.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(columnWidths[4], 10, fmt.Sprintf("$%.2f", total), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Add total amount
	pdf.CellFormat(160, 10, "Total", "1", 0, "R", false, 0, "")
	pdf.CellFormat(30, 10, fmt.Sprintf("$%.2f", grandTotal), "1", 0, "C", false, 0, "")

	// Save PDF
	err := pdf.OutputFileAndClose(os.Getenv("PDF_FOLDER") + invoice.ID + ".pdf")
	if err != nil {
		log.Fatal("Error creating PDF:", err)
	}
}
