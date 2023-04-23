package dtos

import querier "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"

func ToActivityGroupResponse(activityGroup querier.Activity) ActivityGroupResponse {
	return ActivityGroupResponse{
		ID:        activityGroup.ID,
		Title:     activityGroup.Title,
		Email:     activityGroup.Email,
		CreatedAt: &activityGroup.CreatedAt.Time,
		UpdatedAt: &activityGroup.UpdatedAt.Time,
	}
}
