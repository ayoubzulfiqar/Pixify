package imgur

// Image struct represents the JSON fromat image reponse from
// api imgur. Note that every field it's exported and global.
type Image struct {
	Data struct {
		ID          string // The ID for the image
		Title       string // The title of the image.
		Description string // Description of the image.
		DateTime    int    // Time inserted into the gallery, epoch time
		Type        string // Image MIME type.
		Width       int    // The width of the image in pixels
		Height      int    // The height of the image in pixels
		Size        int    // The size of the image in bytes
		Views       int    // The number of image views
		Link        string // Link to to the image
	}
	Success bool   // success request
	Status  uint16 //
}
