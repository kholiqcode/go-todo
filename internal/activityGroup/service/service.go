package service

import (
	"context"
	"fmt"

	"github.com/kholiqcode/go-todolist/internal/activityGroup/dtos"
	querier "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"
	"github.com/kholiqcode/go-todolist/utils"
)

type ActivityGroupService interface {
	FindAll(ctx context.Context) ([]dtos.ActivityGroupResponse, error)
	FindByID(ctx context.Context, id int32) (*dtos.ActivityGroupResponse, error)
	Store(ctx context.Context, request dtos.CreateActivityGroupRequest) (*dtos.ActivityGroupResponse, error)
	Update(ctx context.Context, id int32, request dtos.UpdateActivityGroupRequest) (*dtos.ActivityGroupResponse, error)
	Delete(ctx context.Context, id int32) error
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
		return nil, utils.CustomErrorWithTrace(err, "failed to get activity groups", 400)
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
		return nil, utils.CustomErrorWithTrace(err, fmt.Sprintf("Activity with ID %v Not Found", id), 404)
	}

	activityGroupResp := dtos.ToActivityGroupResponse(activityGroup)

	return &activityGroupResp, nil
}

func (s *activityGroupServiceImpl) Store(ctx context.Context, request dtos.CreateActivityGroupRequest) (*dtos.ActivityGroupResponse, error) {

	params := querier.CreateActivityGroupParams{
		Title: request.Title,
		Email: request.Email,
	}
	res, err := s.repo.CreateActivityGroup(ctx, params)

	if err != nil {
		return nil, utils.CustomErrorWithTrace(err, "failed to create activity group", 400)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, utils.CustomErrorWithTrace(err, "failed to get last inserted id", 400)
	}

	activityGroup, err := s.repo.GetActivityGroup(ctx, int32(insertedID))
	if err != nil {
		return nil, utils.CustomErrorWithTrace(err, "failed to get activity group", 400)
	}

	activityGroupResp := dtos.ToActivityGroupResponse(activityGroup)

	return &activityGroupResp, nil
}

func (s *activityGroupServiceImpl) Update(ctx context.Context, id int32, request dtos.UpdateActivityGroupRequest) (*dtos.ActivityGroupResponse, error) {

	params := querier.UpdateActivityGroupParams{
		ID:    id,
		Title: request.Title,
	}
	err := s.repo.UpdateActivityGroup(ctx, params)

	if err != nil {
		return nil, utils.CustomErrorWithTrace(err, "failed to update activity group", 400)
	}

	activityGroup, err := s.repo.GetActivityGroup(ctx, id)
	if err != nil {
		return nil, utils.CustomErrorWithTrace(err, fmt.Sprintf("Activity with ID %v Not Found", id), 404)
	}

	activityGroupResp := dtos.ToActivityGroupResponse(activityGroup)

	return &activityGroupResp, nil
}

func (s *activityGroupServiceImpl) Delete(ctx context.Context, id int32) error {
	_, err := s.repo.GetActivityGroup(ctx, id)
	if err != nil {
		return utils.CustomErrorWithTrace(err, fmt.Sprintf("Activity with ID %v Not Found", id), 404)
	}

	s.repo.DeleteActivityGroup(ctx, id)

	return nil
}
