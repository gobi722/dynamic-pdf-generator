package pdf_generate

// Invoice structure (adjust as per your MongoDB data)
type Invoice struct {
	ID          string `bson:"_id"`
	From        string `bson:"from"`
	FromAddress string `bson:"fromAddress"`
	To          string `bson:"to"`
	ToAddress   string `bson:"toAddress"`
	Items       []Item `bson:"items"`
}

type Item struct {
	ID          string  `bson:"id"`
	Description string  `bson:"description"`
	Quantity    int     `bson:"quantity"`
	Price       float64 `bson:"price"`
}
