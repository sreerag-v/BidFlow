package domain

import "time"

type Work struct {
	ID                 int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Street             string     `json:"street"`
	DistrictID         int        `json:"district_id"`
	District           District   `json:"-" gorm:"foreignkey:DistrictID;constraint:OnDelete:CASCADE"`
	StateID            int        `json:"state_id"`
	State              State      `json:"-" gorm:"foreignkey:StateID;constraint:OnDelete:CASCADE"`
	TargetProfessionID int        `json:"target_profession"`
	Profession         Profession `json:"-" gorm:"foreignkey:TargetProfessionID;constraint:OnDelete:CASCADE"`
	UserID             int        `json:"user_id"`
	User               User       `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	ProID              int        `json:"pro_id"`
	Provider           Provider   `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	WorkStatus         string     `json:"work_status" gorm:"column:work_status;default:'listed';check:work_status IN ('listed','committed','completed')"`
	BiddedPrice        float32    `json:"bidded_price"`
	PaymentStatus      bool       `json:"payment_status" gorm:"Default:false" `
}

type ReqWork struct {
	UserID             int    `json:"user_id"`
	TargetProfessionID int    `json:"target_profession"`
	Street             string `json:"street"`
	DistrictID         int    `json:"district_id"`
	StateID            int    `json:"state_id"`
}

type WorkspaceImages struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID int    `json:"work_id"`
	Work   Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	Image  string `json:"image"`
}

type CompletedImages struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID int    `json:"work_id"`
	Work   Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	Image  string `json:"image"`
}

type Rating struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Rating   int    `json:"rating" gorm:"rating"`
	Feedback string `json:"feedback"`
	WorkID   int    `json:"work_id"`
	Work     Work   `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
}

type Bid struct {
	ID          int      `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID      int      `json:"work_id"`
	Work        Work     `json:"-" gorm:"foreignkey:WorkID;constraint:OnDelete:CASCADE"`
	ProID       int      `json:"pro_id"`
	Provider    Provider `json:"-" gorm:"foreignkey:ProID;constraint:OnDelete:CASCADE"`
	Estimate    float64  `json:"estimate"`
	Description string   `json:"description"`
	UserId      int      `json:"user_id"`
	AcceptedBid bool     `json:"accepted_bid" gorm:"Default:false"`
	IsDeleted   bool     `json:"is_deleted" gorm:"Default:false"`
}

type AcceptedBidRes struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkID      int     `json:"work_id"`
	Estimate    float64 `json:"estimate"`
	Description string  `json:"description"`
}

type RazorPay struct {
	UserID          int    `JSON:"userid"`
	RazorPaymentId  string `JSON:"razorpaymentid" gorm:"primaryKey"`
	RazorPayOrderID string `JSON:"razorpayorderid"`
	Signature       string `JSON:"signature"`
	AmountPaid      string `JSON:"amountpaid"`
}

type Payment struct {
	PaymentId     int `JSON:"payment_id" gorm:"primarykey"`
	UserId        int
	PaymentMethod string `jSON:"payment_method" gorm:"not null"`
	Totalamount   int    `jSON:"total_amount" gorm:"not null"`
	Status        string `jSON:"Status" gorm:"not null"`
	Date          time.Time
}
