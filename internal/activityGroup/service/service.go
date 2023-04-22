package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/kholiqcode/go-todolist/internal/activityGroup/dtos"
	querier "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"
)

type ActivityGroupService interface {
	FindAll(ctx context.Context) ([]dtos.ActivityGroupResponse, error)
	FindByID(ctx context.Context, id int32) (*dtos.ActivityGroupResponse, error)
}

type activityGroupServiceImpl struct {
	repo querier.ActivityGroupRepo
}

func NewActivityGroupService(repo querier.ActivityGroupRepo) ActivityGroupService {
	return &activityGroupServiceImpl{
		repo: repo,
	}
}

func (s *activityGroupServiceImpl) FindAll(ctx context.Context) ([]dtos.ActivityGroupResponse, error) {
	activityGroups, err := s.repo.ListActivityGroups(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get activity groups: %v", err))
	}

	activityGroupsResp := make([]dtos.ActivityGroupResponse, len(activityGroups))

	for i, activityGroup := range activityGroups {
		activityGroupResp := dtos.ToActivityGroupResponse(activityGroup)
		activityGroupsResp[i] = activityGroupResp
	}

	return activityGroupsResp, nil
}

func (s *activityGroupServiceImpl) FindByID(ctx context.Context, id int32) (*dtos.ActivityGroupResponse, error) {
	activityGroup, err := s.repo.GetActivityGroup(ctx, id)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get activity group: %v", err))
	}

	activityGroupResp := dtos.ToActivityGroupResponse(activityGroup)

	return &activityGroupResp, nil
}
