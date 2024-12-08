package file

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

type ImageType string

const (
	JpgImageType ImageType = ".jpg"
	PngImageType ImageType = ".png"
	SvgImageType ImageType = ".svg"
)

const (
	FilePathAdvantages  string = "./static/images/advantages"
	FilePathStaysImages string = "./static/images/stays_images"
)

const filePath = "./static/images/advantages"

type Service struct{}

func (s *Service) UploadImage(img []byte, t ImageType, category string) (string, error) {
	const op = "service.FileService.CreateImage"

	// Ensure the category directory exists
	categoryPath := fmt.Sprintf("./static/images/%s", category)
	err := os.MkdirAll(categoryPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("%s: failed to create directory %s: %w", op, categoryPath, err)
	}

	// Generate a random file name
	fileName, err := s.GenRandomFileName()
	if err != nil {
		return "", err
	}

	// Create the full file path
	fileWithPath := fmt.Sprintf("%s/%s%s", categoryPath, fileName, t)

	// Create and write to the file
	file, err := os.Create(fileWithPath)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	if _, err := file.Write(img); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return fileWithPath, nil
}

func (s *Service) RemoveFile(fileName string) error {
	const op = "service.FileService.RemoveFile"

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s: %s", op, "file does not exist")
		} else {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	err = os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	return nil
}

func (s *Service) GenRandomFileName() (string, error) {
	const op = "service.FileService.GenRandomFileName"

	rBytes := make([]byte, 16)
	_, err := rand.Read(rBytes)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	rString := hex.EncodeToString(rBytes)

	return rString, nil
}
