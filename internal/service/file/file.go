package file

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

type ImageType string

const (
	JpgImageType     ImageType = ".jpg"
	PngImageType     ImageType = ".png"
	SvgImageType     ImageType = ".svg"
	UnknownImageType ImageType = "unknown"
)

const (
	MaxImageMemorySize = 2 * (1024 * 1024)
)

const (
	FilePathAdvantages         string = "./assets/images/advantages"
	FilePathStaysImages        string = "./assets/images/stays_images"
	FilePathStaysReportsImages string = "./assets/images/stays_reports_images"
	FilePathUsersReportsImages string = "./assets/images/users_reports_images"
)

const filePath = "./assets/images/advantages"

type Service struct{}

func (s *Service) UploadImage(img []byte, t ImageType, category string) (string, error) {
	const op = "service.FileService.CreateImage"

	// Ensure the category directory exists
	categoryPath := fmt.Sprintf(category)
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
