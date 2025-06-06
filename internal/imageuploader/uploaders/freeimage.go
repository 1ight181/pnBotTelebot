package uploaders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	loggerifaces "pnBot/internal/logger/interfaces"
)

const uploadURL = "https://freeimage.host/api/1/upload"

type FreeImageUploader struct {
	apiKey string
	logger loggerifaces.Logger
}

func NewFreeImageUploader(apiKey string, logger loggerifaces.Logger) *FreeImageUploader {
	return &FreeImageUploader{
		apiKey: apiKey,
		logger: logger,
	}
}

type UploadResponse struct {
	StatusCode int `json:"status_code"`
	Success    struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"success"`
	Image struct {
		URL       string `json:"url"`
		URLViewer string `json:"url_viewer"`
		Thumb     struct {
			URL string `json:"url"`
		} `json:"thumb"`
		Medium struct {
			URL string `json:"url"`
		} `json:"medium"`
	} `json:"image"`
	StatusTxt string `json:"status_txt"`
}

func (fiu *FreeImageUploader) UploadImage(file io.Reader, filename string) (string, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("source", filename)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}

	writer.WriteField("key", fiu.apiKey)
	writer.WriteField("format", "json")

	if err := writer.Close(); err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	responce, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer responce.Body.Close()

	var responceData UploadResponse
	if err := json.NewDecoder(responce.Body).Decode(&responceData); err != nil {
		return "", err
	}

	if responceData.Success.Code != 200 {
		return "", fmt.Errorf("загрузка завершилась неудачей с кодом: %d", responceData.StatusCode)
	}

	return responceData.Image.URL, nil
}
