package filehelpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url, filePath, arch string) error {

	// Create the file into the filepath
	filePath = filepath.Join(filePath, "Contents-"+arch+".gz")

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create the file: %v", err)
	}
	defer out.Close()

	// Create the url containing the content indices for the specific arch
	fileUrl := url + "Contents-" + arch + ".gz"

	resp, err := http.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// copy body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}
