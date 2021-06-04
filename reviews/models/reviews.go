package models

import (
	"errors"
	"time"
)

const maxLengthInComments = 400

type Review struct {
	Id      int64
	Stars   int       // 1 - 5
	Comment string    // max 400 chars
	Date    time.Time // created at
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	// Validate that stars are within range
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must between 1 and 5")
	}
	// Validate comment length
	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("comment must be less than 400 characters")
	}
	return nil
}
