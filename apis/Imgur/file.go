package imgur

import (
	"io"
	"log"
	"os"

	"github.com/ayoubzulfiqar/Pixify/errors"
	"github.com/ayoubzulfiqar/Pixify/utils"
)

// DownloadImage make's http get request and save's the file
// with mime extension
func (i Imgur) DownloadImage(url, mime, name string) error {
	var (
		out *os.File
	)

	switch mime {
	case "image/png":
		var (
			err error // for scope reasons
			res string
		)
		res, err = utils.Concat(name, ".png")
		if err != nil {
			return err
		}
		out, err = os.Create(res)
		if err != nil {
			return err
		}
	case "image/jpeg":
		var (
			err error // for scope reasons
			res string
		)
		res, err = utils.Concat(name, ".jpg")
		if err != nil {
			return err
		}
		out, err = os.Create(res)
		if err != nil {
			return nil
		}
	}

	resp, err := i.Get(url)
	if err != nil {
		return err
	}

	_, errCopy := io.Copy(out, resp.Body)
	if errCopy != nil {
		return err
	}

	defer func() {
		if err := out.Close(); err != nil {
			log.Fatal(err)
		}

	}()

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

// SaveImageByte saves images form byte serialize to hdd
func (i Imgur) SaveImageByte(img []byte, path string) error {
	//TODO better opt way
	err := os.WriteFile(path, img, 0644)
	if err != nil {
		return errors.ErrorStat{Message: "Can't save image from ImageURL"}
	}

	return nil
}
