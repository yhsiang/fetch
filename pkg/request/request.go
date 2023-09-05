package request

import (
	"io"
	"net/http"
	"os"
)

func GetHtmlPage(webPage string) (string, error) {
	resp, err := http.Get(webPage)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		return "", err
	}

	return string(body), nil
}

func SaveFile(filename string, content string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}

	return nil
}
