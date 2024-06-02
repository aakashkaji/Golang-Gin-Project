package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
)

type EmpService interface {
}

type EmpolyeeService struct {
	repo domain.EmpolyeeRepoDB
}

func NewEmpolyeeService(repo domain.EmpolyeeRepoDB) *EmpolyeeService {
	return &EmpolyeeService{repo: repo}
}

func (s *EmpolyeeService) NewCreateEmpolyee(ctx context.Context, req domain.EmpolyeeRequestDto) (*domain.Empolyee, error) {

	if apiErr := NewEmpolyeeValidate(req); apiErr != nil {
		return nil, apiErr
	}

	empolyee := domain.Empolyee{
		Name:      req.Name,
		Position:  req.Position,
		Salary:    req.Salary,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	// todo create empolyee

	createEmpolyee, err := s.repo.CreateEmpolyee(ctx, empolyee)

	if err != nil {
		return nil, err
	}

	// response dto if you want to define

	return createEmpolyee, nil

}

func (s *EmpolyeeService) UpdateEmpolyee(ctx context.Context, req domain.EmpolyeeUpdateRequestDto) (*domain.Empolyee, error) {

	// perform validation

	_, err := s.repo.FindByEmpolyeeId(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	// call database to update
	req.UpdateAt = time.Now()

	emp, errs := s.repo.UpdateEmpolyee(ctx, &req)
	if errs != nil {
		return nil, errs
	}
	// then return

	return emp, nil
}

func (s *EmpolyeeService) GetEmpolyeeId(ctx context.Context, empolyeeId int) (*domain.Empolyee, error) {

	empolyee, err := s.repo.FindByEmpolyeeId(ctx, empolyeeId)
	if err != nil {
		return nil, err
	}
	return empolyee, nil

}

func (s *EmpolyeeService) GetAllEmpolyee(ctx context.Context, pageNo, pageSize string) (*domain.EmpolyeeResponseDto, error) {

	// validation

	opts, _ := validateFindAllEmpolyeeList(pageNo, pageSize)

	empolyees, err := s.repo.AllEmpolyee(ctx, opts)

	if err != nil {
		return nil, err
	}
	return empolyees, nil

}

func (s *EmpolyeeService) DeleteEmpolyee(ctx context.Context, id int) error {

	// validation

	// check record exists

	empolyee, errs := s.repo.FindByEmpolyeeId(ctx, id)
	if errs != nil {
		return errs //errors.New("Record Not found!")
	}

	fmt.Println(empolyee)

	err := s.repo.DeleteEmpolyeeRecord(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
