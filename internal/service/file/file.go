package file

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type ImageType string

const (
	JpgImageType     ImageType = ".jpg"
	PngImageType     ImageType = ".png"
	SvgImageType     ImageType = ".svg"
	UnknownImageType ImageType = "unknown"

	MaxImageMemorySize = 2 * (1024 * 1024)
)

type PathType string

const (
	FilePathAdvantages         PathType = "/images/advantages"
	FilePathStaysImages        PathType = "/images/stays_images"
	FilePathUsersReportsImages PathType = "/images/users_reports_images"
	FilePathStaysReportsImages PathType = "/images/stays_reports_images"
)

type Service struct{}

func (s *Service) UploadImage(img []byte, t ImageType, filePath PathType) (string, error) {
	const op = "service.FileService.CreateImage"

	fileName, err := s.GenRandomFileName()
	if err != nil {
		return "", err
	}

	fileWithPath := fmt.Sprintf("./static%s/%s%s", filePath, fileName, t)

	if err := os.MkdirAll(filepath.Dir(fileWithPath), os.ModePerm); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	file, err := os.Create(fileWithPath)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	if _, err := file.Write(img); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Sprintf("%s/%s%s", filePath, fileName, t), nil
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
