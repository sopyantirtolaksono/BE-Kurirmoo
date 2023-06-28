package storage

var (
	MimeTypePNG  = "image/png"
	MimeTypeJPEG = "image/jpeg"

	MimetypeImageAllowed = map[string]string{
		MimeTypePNG:  MimeTypePNG,
		MimeTypeJPEG: MimeTypeJPEG,
	}
)

// CustomerType is a map for customer type with boolean value
var CustomerType = map[string]bool{
	"corporate": true,
	"personal":  false,
}
