package imageshack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ayoubzulfiqar/Pixify/errors"
)

// Get request
func (i ImageShack) Get(url string) (*http.Response, error) {

	// make new pointer to *http.Request
	// with GET method, url and a new byes.Buffer body
	req, err := http.NewRequest("GET", url, new(bytes.Buffer))
	if err != nil {
		return nil, &errors.ErrorStat{Message: "Can't create new GET request to Image Shack"}
	}

	// allocate new *http.Client
	client := &http.Client{}

	// make client request with the http.Request above declared
	resp, err := client.Do(req)
	if err != nil {
		return nil, &errors.ErrorStat{Message: "GET request to ImageShack api failed"}
	}

	if resp.StatusCode > http.StatusPartialContent {
		return nil, &errors.ErrorStat{Message: fmt.Sprintf("%s %s", resp.Status, "POST request to ImageShack api failed")}
	}

	// return it
	return resp, nil
}

// ImageJSON request imageshack
func (i ImageShack) ImageJSON(url string) (*Image, error) {

	resp, err := i.Get(url)
	if err != nil {
		return nil, err
	}

	// allocate imgJSON struct for decoding the response body
	imgJSON := &Image{}

	// decode into JSON
	err = json.NewDecoder(resp.Body).Decode(imgJSON)
	if err != nil {
		return nil, err
	}

	// close body response
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// return the JSON
	return imgJSON, nil
}

// ImageByte reads images from http response
// and serialize it into byte
func (i ImageShack) ImageByte(url string) ([]byte, error) {
	resp, err := i.Get(url)
	if err != nil {
		return nil, err
	}

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// close body response
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return read, nil
}
