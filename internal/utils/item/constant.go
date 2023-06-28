package item

// enum item category
var (
	HEAVY_ITEM          = "Barang Berat"
	WORSE_OF_PLACES     = "Barang Makan Tempat"
	KEY_HEAVY_ITEM      = "barang_berat"
	KEY_WORSE_OF_PLACES = "barang_makan_tempat"

	ItemCategory = map[string]*string{
		KEY_HEAVY_ITEM:      &HEAVY_ITEM,
		KEY_WORSE_OF_PLACES: &WORSE_OF_PLACES,
	}
)
