package models

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt"
)

type AdminLogin struct {
	Email    string `json:"email" gorm:"validate:required,email" validate:"email"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DoubleTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AdminDetailsResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email" `
	Previlege string `json:"previlege"`
}

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type VerificationDetails struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	DocumentImage string   `json:"document_images"`
	Services      []string `json:"services"`
}

type Verification struct {
	ID   int
	Name string
}

type ProviderRegister struct {
	Name       string                `json:"name"`
	Email      string                `json:"email "`
	Password   string                `json:"password"`
	RePassword string                `json:"re-password"`
	Phone      string                `json:"phone+6+"`
	Document   *multipart.FileHeader `json:"document"`
}

type PageNation struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}

type UserDetails struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" `
	Phone     string `json:"phone" `
	IsBlocked bool   `json:"is_blocked"`
}

type OTPLoginStruct struct {
	Name  string `json:"name" binding:"omitempty,min=3,max=16"`
	Email string `json:"email" binding:"omitempty,email"`
	Phone string `json:"phone" binding:"omitempty,min=10,max=10"`
}

type OTPVerifyStruct struct {
	OTP    string `json:"otp" binding:"required,min=4,max=8"`
	UserID uint   `json:"user_id" binding:"required,numeric"`
	Email  string `json:"email" `
}

type ProviderDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsVerified bool   `json:"is_verified"`
}

type WorkDetails struct {
	ID            int      `json:"id"`
	Street        string   `json:"street"`
	District      string   `json:"district" `
	State         string   `json:"state" `
	Profession    string   `json:"profession"`
	User          string   `json:"user_name"`
	Provider      string   `json:"provider"`
	Images        []string `json:"images" `
	Participation bool     `json:"participation"`
	WorkStatus    string   `json:"work_status" `
}

type MinWorkDetails struct {
	ID         int    `json:"id"`
	Street     string `json:"street"`
	District   string `json:"district" `
	State      string `json:"state" `
	Profession string `json:"profession"`
	User       string `json:"user_name"`
	Provider   string `json:"provider"`
	WorkStatus string `json:"work_status" `
}

type GetServices struct {
	ID          int    `json:"id"`
	ServiceName string `json:"service"`
	Category_id int    `json:"category_id"`
}

type GetLocations struct {
	ID       int    `json:"id"`
	District string `json:"district"`
	State    string `json:"state"`
}

type BidDetails struct {
	Work_id          int     `json:"work_id"`
	Provider    string  `json:"provider"`
	ProviderID  int     `json:"provider_id"`
	Estimate    float64 `json:"estimate"`
	Description string  `json:"description"`
}

type AddNewState struct {
	State string `json:"state"`
}

type AddNewDistrict struct {
	StateID  int    `json:"state_id"`
	District string `json:"district"`
}

type CreateCategory struct {
	Category string `json:"category"`
}

type AddServicesToACategory struct {
	CategoryID  int    `json:"category_id"`
	ServiceName string `json:"service"`
}

type PlaceBid struct {
	WorkID      int     `json:"-"`
	ProID       int     `json:"-"`
	Estimate    float64 `json:"estimate"`
	Description string  `json:"description"`
}

type UserSignup struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,number"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,eqfield=Password"`
}

type ProviderDetailsForUser struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	AverageRating int    `json:"rating"`
}
type ProviderProfile struct {
	Image         string `json:"image"`
	Name          string `json:"name" gorm:"not null"`
	Email         string `json:"email" gorm:"not null" validate:"required,email"`
	Phone         string `json:"phone" gorm:"not null"`
	Profession    string `json:"profession" gorm:"not null"`
	District      string `json:"district" gorm:"not null"`
	AverageRating int    `json:"rating" gorm:"not null"`
}
type RatingModel struct {
	Rating   float32 `json:"rating"`
	Feedback string  `json:"feedback"`
}

type UpdateUser struct {
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Phone string `json:"phone" gorm:"unique;not null"`
}

type Forgott struct {
	Email string `json:"email" gorm:"not null"`
}

type ChangePassword struct {
	Email           string
	Otp             string
	Password        string
	ConfirmPassword string
}
