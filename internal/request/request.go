package request

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading body of request: %v", err)
		return nil, err
	}

	return body, nil
}
