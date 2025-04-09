# Dynamic PDF Generator 📄

A backend utility written in Go to generate custom PDF files in bulk using data fetched from MongoDB. Each PDF is tailored uniquely for every record in the collection — ideal for reports, invoices, letters, certificates, or any document automation needs.

## ⚙️ Features

- Fetches data dynamically from MongoDB
- Generates `N` number of unique PDFs based on document content
- Supports customizable layout and formatting
- Efficient batch processing for large data
- Output files are named meaningfully (like `invoice_<id>.pdf`)

## 🧰 Tech Stack

- Go (Golang)
- MongoDB (as source of records)
- gofpdf (PDF generation library)

## 🧪 Example Use Cases

- Generating invoices for all users
- Certificates for event participants
- Reports or summaries fetched from database
- Letters or notices with user-specific data

## 🗃️ MongoDB Sample Record

```json
{
  "_id": "abc123",
  "name": "Gobinath R",
  "amount": "500.00",
  "date": "2024-04-01",
  "purpose": "Training Invoice"
}
