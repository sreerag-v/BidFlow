package interfaces

import (
	"mime/multipart"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type Helper interface {
	CreateHashPassword(string) (string, error)
	CompareHashAndPassword(a string, b string) error
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	UploadToS3(file *multipart.FileHeader) (string, error)

	GenerateTokenProvider(details domain.Provider) (string, error)
	GenerateTokenUser(details domain.User) (string, error)

	StringToUInt(str string) (uint, error)

}