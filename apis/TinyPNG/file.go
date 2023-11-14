package tinypng

import (
	"os"

	"github.com/ayoubzulfiqar/Pixify/utils"
)

func (t Tiny) SaveImage(img []byte, name, mime string) error {

	if mime == "image/jpeg" {
		//Write compressed file
		nameJPG, err := utils.Concat(name, ".jpg")
		if err != nil {
			nameJPG = "defaultCompressed.jpg"
		}
		errWrite := os.WriteFile(nameJPG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	} else if mime == "image/png" {
		namePNG, err := utils.Concat(name, ".png")
		if err != nil {
			namePNG = "defaultCompressed.png"
		}
		errWrite := os.WriteFile(namePNG, img, 0644)
		if errWrite != nil {
			return errWrite
		}
	}
	return nil
}
