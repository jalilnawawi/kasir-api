package models

type Kategori struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var ListKategori = []Kategori{
	{
		ID:          1,
		Name:        "Elektronik",
		Description: "Berbagai perangkat elektronik seperti TV, audio, kamera, dan peralatan rumah tangga elektronik.",
	},
	{
		ID:          2,
		Name:        "Komputer & Aksesoris",
		Description: "Perangkat komputer, laptop, serta aksesoris pendukung seperti keyboard, mouse, dan storage.",
	},
	{
		ID:          3,
		Name:        "Handphone & Aksesoris",
		Description: "Smartphone dan aksesoris pendukung seperti charger, casing, kabel data, dan headset.",
	},
}
