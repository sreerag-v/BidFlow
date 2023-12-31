package helper

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	cfg "github.com/sreerag_v/BidFlow/pkg/config"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/helper/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type helper struct {
	config cfg.Config
}

func NewHelper(cfg cfg.Config) interfaces.Helper {
	return &helper{
		config: cfg,
	}
}

func (h *helper) CreateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (helper *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	tokenClaims := &models.AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  admin.Previlege,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("adminsecret")) //take this from runtime in future avoid hardcoding
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *helper) UploadToS3(file *multipart.FileHeader) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		fmt.Println("configuration error:", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	f, openErr := file.Open()
	if openErr != nil {
		fmt.Println("opening error:", openErr)
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("bucket-diplo"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if uploadErr != nil {
		fmt.Println("uploading error:", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}

func (helper *helper) GenerateTokenProvider(details domain.Provider) (string, error) {
	accessTokenClaims := &models.AuthCustomClaims{
		Id:    details.ID,
		Email: details.Email,
		Role:  "provider",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("providersecret")) //take this from runtime in future avoid hardcoding
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (helper *helper) GenerateTokenUser(details domain.User) (string, error) {
	accessTokenClaims := &models.AuthCustomClaims{
		Id:    details.ID,
		Email: details.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("usersecret")) //take this from runtime in future avoid hardcoding
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (helper *helper) StringToUInt(str string) (uint, error) {
	if str == "" {
		return 0, errors.New("empty string ")
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	// fmt.Println("xxxxx", uint(val))
	return uint(val), nil
}
