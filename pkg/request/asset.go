package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func SaveAsset(baseURL string, filepath string) error {
	// Skip absolute url and only save relative files
	if !strings.HasPrefix(filepath, "/") {
		return nil
	}

	resp, err := http.Get(fmt.Sprintf("%s%s", baseURL, filepath))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		return err
	}

	f, err := create(fmt.Sprintf("./%s", filepath))

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(body)

	if err != nil {
		return err
	}

	return nil
}
