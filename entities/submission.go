package entities

import "github.com/gofiber/fiber/v2"

const (
	SubmissionStatusPending = "PENDING"
	SubmissionStatusCorrect = "CORRECT"
	SubmissionStatusWrong   = "WRONG"
)

type Submission struct {
	SubmissionID uint      `json:"submission_id" gorm:"primaryKey"`
	ChallengeID  uint      `json:"challenge_id"`
	Challenge    Challenge `json:"challenge" gorm:"foreignKey:ChallengeID"`
	UserID       uint      `json:"user_id"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	Language     string    `json:"language"`
	SourceCode   string    `json:"source_code"`
	Status       string    `json:"status" gorm:"default:PENDING"`
}

type SubmissionCreateDTO struct {
	ChallengeID uint   `json:"challenge_id" validate:"required"`
	Language    string `json:"language" validate:"required"`
	SourceCode  string `json:"source_code" validate:"required"`
}

func ValidateSubmissionCreateDTO(c *fiber.Ctx) SubmissionCreateDTO {
	var dto SubmissionCreateDTO

	if err := c.BodyParser(&dto); err != nil {
		panic(err)
	}

	if err := validate.Struct(&dto); err != nil {
		panic(err)
	}

	return dto
}
