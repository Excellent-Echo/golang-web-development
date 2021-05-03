package entity

type Product struct {
	ID           int
	Name         string
	Price, Stock int
}

func (p *Product) StockStatus() string {
	if p.Stock <= 15 {
		return "stock tinggal sedikit"
	} else {
		return "stock masih banyak"
	}
}
