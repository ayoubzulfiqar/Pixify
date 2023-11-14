package utils

import (
	"fmt"
	"io"
	"net/http"
)

// DumpResponse for debugging information
func DumpResponse(resp *http.Response) {
	for key, val := range resp.Header {
		fmt.Print(key)
		fmt.Print(" : ")
		fmt.Println(val)
	}
	fmt.Printf("Status : %s\n", resp.Status)
	read, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}
	fmt.Println()
	fmt.Println("===================== BODY ====================")
	fmt.Println()
	fmt.Println(string(read))

	// for any case
	defer resp.Body.Close()
}
