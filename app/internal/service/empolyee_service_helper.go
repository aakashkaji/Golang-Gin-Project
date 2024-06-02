package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
)

// Empolyee data validation

func NewEmpolyeeValidate(req domain.EmpolyeeRequestDto) error {

	var errorMessages []string

	if req.Name == "" {
		errorMessages = append(errorMessages, "Name cannot blank ")
	}

	if req.Position == "" {
		errorMessages = append(errorMessages, "Position cannot blank ")
	}

	if req.Salary <= 0 {

		errorMessages = append(errorMessages, "Salary must be greater than 0")
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "; "))
	}
	return nil

}

func validateFindAllEmpolyeeList(pageNo, pageSize string) (domain.Pagination, error) {

	opts := domain.Pagination{
		RecordPerPage: 20,
		CurrentPage:   1,
	}

	if pageSize != "" {
		vl, _ := strconv.Atoi(pageSize)
		opts.RecordPerPage = vl
	}

	if pageNo != "" {
		vl, _ := strconv.Atoi(pageNo)
		if vl > 1 {
			opts.CurrentPage = vl * opts.RecordPerPage
		} else {
			opts.CurrentPage = 0
		}
	}

	return opts, nil

}
