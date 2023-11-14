package imageshack

// "success":true,
//    "process_time":32,
//    "result":{
//       "id":"idggZf2Aj",
//       "server":661,
//       "bucket":8761,
//       "filename":"ggZf2A.jpg",
//       "original_filename":"e1ab262c288441e252ea507676e6c8a8.jpg",
//       "direct_link":"imagizer.imageshack.com\/img661\/8761\/ggZf2A.jpg",

// Image struct that hold JSON http response
type Image struct {
	Result struct {
		DirectLink string
	}
}

// ImageShack naked struct
// Note that naked structs are zero cost memory
type ImageShack struct {
}

// NewImageShack allocates and returns a new *ImageShack
func NewImageShack() (i *ImageShack) {
	return &ImageShack{}
}
