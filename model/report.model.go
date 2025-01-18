package model

type TopProduct struct {
	ProductID   uint
	ProductName string
	TotalSold   int
}

type LowStockProduct struct {
	ProductID   uint
	ProductName string
	Stock       int
	LowStock    int
}
