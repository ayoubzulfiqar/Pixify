package utils

import (
	"io"
	"log"
	"net/http"
)

// ResponseStringReader reads resp.Body *http.Response
// into a string container and closes the resp.Body.Close()
func ResponseStringReader(resp *http.Response) (stringContainer string, err error) {
	var (
		buffer    []byte
		n         int
		container []byte
	)
	// declare buffer to be 1kb
	buffer = make([]byte, 1024)
	for {
		n, err = resp.Body.Read(buffer)
		// test if err
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return "", err
		}

		// append buffer to container
		container = append(container, buffer...)
	}

	return string(container), nil
}

// ResponseByteReader reads resp.Body *http.Response
// into a byte container anc closes the resp.Body.Close()[
func ResponseByteReader(resp *http.Response) (byteContainer []byte, err error) {
	var (
		buffer    []byte
		n         int
		container []byte
	)
	// declare buffer to be 1kb
	buffer = make([]byte, 1024)
	for {
		n, err = resp.Body.Read(buffer)
		// test if err
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return nil, err
		}

		// append buffer to container
		container = append(container, buffer...)
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return container, nil
}
