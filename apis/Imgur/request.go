package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	io "io/ioutil"
	"net/http"

	"github.com/ayoubzulfiqar/Pixify/errors"
)

// Get request to imgur api with setting the mime content and url
// returns a pointer to http.Response and in case of error  a pointer to ErrorStat
func (i Imgur) Get(url string) (*http.Response, *errors.ErrorStat) {
	req, err := http.NewRequest("GET", url, new(bytes.Buffer))
	if err != nil {
		return nil, &errors.ErrorStat{Message: "Can't create new request to ImageURL"}
	}
	// add the corresponding header
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", i.clientID))
	//allocate new *http.Client
	client := &http.Client{}

	//make client request with the http.Request above declared.
	resp, err := client.Do(req)
	if err != nil {
		return nil, &errors.ErrorStat{Message: "Get request to ImageURL api failed"}
	}

	// check http status to be ok(200)
	errStatHTTP := errorHTTPStatus(resp.StatusCode)
	if errStatHTTP != nil {
		return nil, errStatHTTP
	}

	// return it
	return resp, nil
}

// ImageJSON returns serialize get response
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageJSON(url string) (*Image, error) {
	resp, errGet := i.Get(url)
	if errGet != nil {
		return nil, errGet
	}

	imgJSON := &Image{}
	errJSON := json.NewDecoder(resp.Body).Decode(imgJSON)
	if errJSON != nil {
		return nil, errJSON
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	return imgJSON, nil
}

// ImageByte http request and return image serialized body
// https://api.imgur.com/3/gallery/image/{id}
func (i Imgur) ImageByte(url string) ([]byte, error) {
	resp, errGet := i.Get(url)
	if errGet != nil {
		return nil, errGet
	}

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	return read, nil
}
