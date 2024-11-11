package stays

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	staysInterface "github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"mime/multipart"
	"sync"
)

type Service struct {
	Repo    staysInterface.StaysRepo
	LocSvc  staysInterface.LocationService
	FileSvc staysInterface.FileService
	UserSvc staysInterface.UserService
}

func (s *Service) CreateStay(ctx context.Context, stay *stays.StayEntity) error {
	const op = "service.stays.CreateStay"

	foundLocation, err := s.LocSvc.GetByID(ctx, stay.LocationID)
	if err != nil {
		return err
	}
	if foundLocation.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	user, err := s.UserSvc.GetUserByID(ctx, stay.UserID.String())
	if err != nil {
		return err
	}

	if user.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	err = s.Repo.CreateStay(ctx, stay)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStayByID(ctx context.Context, id uuid.UUID) (*stays.Stay, error) {
	const op = "service.stays.GetStayByID"

	stay, err := s.Repo.GetStayByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if stay == nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	return stay, nil
}

func (s *Service) GetStays(ctx context.Context) ([]*stays.Stay, error) {
	const op = "service.stays.GetStays"

	stays, err := s.Repo.GetStays(ctx)
	if err != nil {
		return nil, err
	}

	return stays, nil
}

func (s *Service) DeleteStayByID(ctx context.Context, id uuid.UUID) error {
	const op = "service.stays.DeleteStay"

	exists, err := s.Repo.CheckStayIfExistsByID(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	err = s.Repo.DeleteStayByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateStayByID(ctx context.Context, stay *stays.StayEntity, id uuid.UUID) (*stays.Stay, error) {
	const op = "service.stays.UpdateStayByID"

	foundLocation, err := s.LocSvc.GetByID(ctx, stay.LocationID)
	if err != nil {
		return &stays.Stay{}, err
	}

	if foundLocation == nil {
		return &stays.Stay{}, fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	err = s.Repo.UpdateStayByID(ctx, stay, id)
	if err != nil {
		return &stays.Stay{}, err
	}

	updatedStay, err := s.Repo.GetStayByID(ctx, id)
	if err != nil {
		return &stays.Stay{}, err
	}

	return updatedStay, nil
}

func (s *Service) GetStaysByUserID(ctx context.Context, userId uuid.UUID) ([]*stays.Stay, error) {
	const op = "service.stays.GetStaysByUserID"

	user, err := s.UserSvc.GetUserByID(ctx, userId.String())
	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	usrStays, err := s.Repo.GetStaysByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return usrStays, nil
}

func (s *Service) GetImagesByStayID(ctx context.Context, id uuid.UUID) ([]stays.StayImage, error) {
	const op = "service.stays.GetImagesByStayID"

	foundStay, err := s.GetStayByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if foundStay == nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	images, err := s.Repo.GetImagesByStayID(ctx, id)
	if err != nil {
		return nil, err
	}

	return images, nil

}

func (s *Service) GetMainImageByStayID(ctx context.Context, id uuid.UUID) (stays.StayImage, error) {
	const op = "service.stays.GetMainImageByStayID"

	foundStay, err := s.GetStayByID(ctx, id)
	if err != nil {
		return stays.StayImage{}, err
	}

	if foundStay == nil {
		return stays.StayImage{}, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	image, err := s.Repo.GetMainImageByStayID(ctx, id)
	if err != nil {
		return stays.StayImage{}, err
	}

	return image, nil
}

func (s *Service) CreateImages(ctx context.Context, filesHeaders []*multipart.FileHeader, stayID uuid.UUID) error {
	const op = "service.stays.CreateImages"

	isExists, err := s.Repo.CheckStayIfExistsByID(ctx, stayID)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	errChan := make(chan error, len(filesHeaders))

	var wg sync.WaitGroup
	wg.Add(len(filesHeaders))

	for _, fh := range filesHeaders {
		go func(fileHeader *multipart.FileHeader) {
			defer wg.Done()

			img, err := fileHeader.Open()
			if err != nil {
				errChan <- fmt.Errorf("%s: %w", op, err)
			}
			defer img.Close()

			buf := make([]byte, fileHeader.Size)

			n, err := img.Read(buf)
			if err != nil {
				errChan <- fmt.Errorf("%s: %w", op, err)
			}

			fileName, err := s.FileSvc.UploadImage(buf[:n], file.JpgImageType, file.FilePathStaysImages)
			if err != nil {
				errChan <- err
			}

			err = s.Repo.CreateStayImage(ctx, fileName, false, stayID)
			if err != nil {
				errChan <- err
			}
		}(fh)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) CreateMainImage(ctx context.Context, fileHeader *multipart.FileHeader, stayID uuid.UUID) error {
	const op = "service.stays.CreateMainImages"

	isExists, err := s.Repo.CheckStayIfExistsByID(ctx, stayID)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	img, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer img.Close()

	buf := make([]byte, fileHeader.Size)

	n, err := img.Read(buf)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fileName, err := s.FileSvc.UploadImage(buf[:n], file.JpgImageType, file.FilePathStaysImages)
	if err != nil {
		return err
	}
	err = s.Repo.CreateStayImage(ctx, fileName, true, stayID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteStayImage(ctx context.Context, imageID uuid.UUID) error {
	const op = "service.stays.DeleteStayImage"

	stayImage, err := s.Repo.GetStayImageByID(ctx, imageID)
	if err != nil {
		return err
	}

	if stayImage.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrStayImageNotFound)
	}

	err = s.FileSvc.RemoveFile(stayImage.ImageName, file.FilePathStaysImages)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteStayImage(ctx, imageID)
	if err != nil {
		return err
	}
	return nil
}
