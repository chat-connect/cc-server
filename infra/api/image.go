package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Request struct {
	ImageFile string `json:"image_file"`
	ImageName string `json:"image_name"`
	ImagePath string `json:"image_path"`
}

func UploadImage(imageFile string, imageName string, imagePath string) error {
	request, err := json.Marshal(&Request{
		ImageFile: imageFile,
		ImageName: imageName,
		ImagePath: imagePath,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/image/upload", os.Getenv("GC_IMAGE_URL")), bytes.NewBuffer(request))
	if err != nil {
		fmt.Println("Failed to create request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
	}
	defer resp.Body.Close()

	return nil
}
