package configs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageKitResponse struct {
	URL          string `json:"url"`
	FileID       string `json:"fileId"`
	Name         string `json:"name"`
	FilePath     string `json:"filePath"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func UploadToImageKit(file *multipart.FileHeader) (string, error) {
	// Read file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Read file content
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, src); err != nil {
		return "", err
	}

	// Convert to base64
	base64Data := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// Prepare the request
	url := "https://upload.imagekit.io/api/v1/files/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("file", base64Data)
	_ = writer.WriteField("fileName", file.Filename)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	// Add authentication
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("IMAGEKIT_PRIVATE_KEY") + ":"))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Parse response
	var result ImageKitResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to upload image: %s", res.Status)
	}

	return result.URL, nil
}
