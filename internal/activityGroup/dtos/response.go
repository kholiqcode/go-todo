package dtos

import "time"

type ActivityGroupResponse struct {
	ID        int32      `json:"id"`
	Title     string     `json:"title"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
