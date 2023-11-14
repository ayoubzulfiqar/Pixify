package tinypng

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ayoubzulfiqar/Pixify/errors"
)

// Post request to tiny
func (t Tiny) Post(mime string) (*http.Response, error) {
	//serialized body for post request
	postBytesReader := bytes.NewReader(t.body)

	// new *http.Request{}
	request, err := http.NewRequest("POST", urlSender, postBytesReader)
	if err != nil {
		return nil, &errors.ErrorStat{Message: "Can't create new POST request to Tiny"}
	}

	//sanitize header for auth header
	authValue := headerToBase64(username, key)
	request.Header.Add("Content-Type", mime)
	request.Header.Add("Authorization", authValue)

	// new *http.Client struct
	cli := &http.Client{}

	// make POST request to tiny api
	resp, err := cli.Do(request)
	if err != nil {
		return nil, &errors.ErrorStat{Message: "POST request to Tiny api failed"}
	}

	if resp.StatusCode > http.StatusPartialContent {
		return nil, &errors.ErrorStat{Message: fmt.Sprintf("%s %s", resp.Status, "POST request to Tiny api failed")}
	}
	// for debugging reasons
	fmt.Printf("%+v\n", resp)

	return resp, nil
}

// HeaderToBase64 Base64 encoding defined by rfc RFC2045-MIME
// plus concat the "Basic: "word.
func headerToBase64(username, key string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s", username, key)))
	return fmt.Sprintf("Basic %s", sEnc)
}

// Get image from tiny api
func (t Tiny) Get(url string) ([]byte, error) {
	respImage, errGET := http.Get(url)
	if errGET != nil {
		return nil, errors.ErrorStat{Message: "Error on getting image from Tiny api"}
	}
	//todo
	readIMG, errREAD := io.ReadAll(respImage.Body)
	if errREAD != nil {
		return nil, errors.ErrorStat{Message: "Error on reading response body image from Tiny api"}
	}

	func() {
		err := respImage.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return readIMG, nil
}

// PostGetJSON function that sends image and receives json
func (t Tiny) PostGetJSON(mime string) (*TinyJSON, error) {
	resp, err := t.Post(mime)
	if err != nil {
		return nil, err
	}

	// allocate TinyJSON struct for parsing JSON response
	v := &TinyJSON{}

	// parse JSON from response
	err = json.NewDecoder(resp.Body).Decode(v)
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

	return v, nil
}
