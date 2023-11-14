package server

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	imageshack "github.com/ayoubzulfiqar/Pixify/apis/ImageShack"
	imgur "github.com/ayoubzulfiqar/Pixify/apis/Imgur"
	tinypng "github.com/ayoubzulfiqar/Pixify/apis/TinyPNG"
	"github.com/ayoubzulfiqar/Pixify/utils"
)

const (
	clientID     = "40aea5f08c0f717"
	clientSecret = "a72c35d27b38d27114b4503e5b9acc835861ed8c"
)

// ImgurToTiny just send image to tinyPNG
func imgurToTiny() {
	// allocate
	imgur := imgur.NewImgur()
	tiny := tinypng.NewTiny()
	// set
	imgur.SetClientID(clientID)
	imgur.SetClientSecret(clientSecret)
	imgur.SetBody(new(bytes.Buffer))
	// get request JSON response
	v, err := imgur.ImageJSON("https://api.imgur.com/3/gallery/image/CFzq6zN")
	if err != nil {
		log.Fatal(err)
	}
	if v.Data.Type == "image/png" || v.Data.Type == "image/jpeg" {
		//err := imgur.DownloadImage(v.Data.Link, v.Data.Type, "myNewEPicImage")
		byteImage, err := imgur.ImageByte(v.Data.Link)
		if err != nil {
			log.Fatal(err)
		}
		// prepare tiny for post Request
		// setting body with the image downloaded from imgur
		tiny.SetBody(byteImage)
		// post request and get json response parsed
		// t parsed json
		t, err := tiny.PostGetJSON(v.Data.Type)
		if err != nil {
			log.Fatal(err)
		}

		// get pic from tiny api from json info
		pic, err := tiny.Get(t.Output.URL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		}

		// save image pic
		err = tiny.SaveImage(pic, "newImageCompressedImgur", v.Data.Type)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
		os.Exit(1)
	}
}

// ShackToTiny send image from shack to tiny compression
func shackToTiny() {
	// alloc
	shack := imageshack.NewImageShack()
	tiny := tinypng.NewTiny()
	//v, err := shack.ImageJSON("https://api.imageshack.com/v2/images/pbzPCsEij")
	v, err := shack.ImageJSON("https://api.imageshack.com/v2/images/idbwalqaj")
	if err != nil {
		log.Fatal(err)
	}
	// add https:// to link
	url, err := utils.Concat("https://", v.Result.DirectLink)

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(url, ".jpg") {
		byteImage, err := shack.ImageByte(url)
		if err != nil {
			log.Fatal(err)
		}
		tiny.SetBody(byteImage)
		// prepare tiny for post Request
		// setting body with the image downloaded from imgur
		tiny.SetBody(byteImage)
		// post request and get json response parsed
		// t parsed json
		t, err := tiny.PostGetJSON("image/jpeg")
		if err != nil {
			log.Fatal(err)
		}

		// get pic from tiny api from json info
		pic, err := tiny.Get(t.Output.URL) // byte []byte
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		}

		// save image pic
		err = tiny.SaveImage(pic, "NewImageShackCompressedJPG", "image/jpeg")
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.Contains(url, ".png") {
		byteImage, err := shack.ImageByte(url)
		if err != nil {
			log.Fatal(err)
		}
		tiny.SetBody(byteImage)
		// prepare tiny for post Request
		// setting body with the image downloaded from imgur
		tiny.SetBody(byteImage)
		// post request and get json response parsed
		// t parsed json
		t, err := tiny.PostGetJSON("image/png")
		if err != nil {
			log.Fatal(err)
		}

		// get pic from tiny api from json info
		pic, err := tiny.Get(t.Output.URL) // byte []byte
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		}

		// save image pic
		err = tiny.SaveImage(pic, "NewImageShackCompressedPNG", "image/png")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: image is not PNG/JPG type\n")
		fmt.Fprintf(os.Stderr, "Please enter a valid PNG/JPG file type\n")
		os.Exit(1)
	}
}
