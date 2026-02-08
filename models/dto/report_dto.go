package dto

type ReportDto struct {
	TotalRevenue   int            `json:"total_revenue"`
	TotalTransaksi int            `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
	Periode        string         `json:"periode"`
}

type ProdukTerlaris struct {
	ProductName string `json:"name"`
	QtyTerjual  int    `json:"qty_terjual"`
}
