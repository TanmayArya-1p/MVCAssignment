package utils

import (
	"errors"
	"fmt"
	"inorder/pkg/config"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func GenerateImageUploadPath(filename string) (string, string, error) {

	var outputDir string = config.Config.InOrder.ITEM_IMAGE_DIRECTORY

	err := os.MkdirAll(outputDir, 0644)
	if err != nil {
		return "", "", err
	}

	ext := strings.ToLower(filepath.Ext(filename))
	if ext != ".jpg" && ext != ".png" {
		return "", "", errors.New("Invalid image file. Only png/jpg allowed")
	}

	name := filepath.Base(filename)
	if strings.HasSuffix(name, ".png") {
		name = strings.TrimSuffix(name, ".png")
	} else if strings.HasSuffix(name, ".jpg") {
		name = strings.TrimSuffix(name, ".jpg")
	}

	nonce := fmt.Sprintf("%04d", rand.Intn(10000))
	fileName := name + "_" + nonce + ext
	filePath := outputDir + "/" + fileName
	fileURL := strings.TrimPrefix(filePath, "public")
	return filePath, fileURL, nil
}
