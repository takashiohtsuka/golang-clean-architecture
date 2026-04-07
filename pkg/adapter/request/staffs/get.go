package staffs

import (
	"errors"
	"golang-clean-architecture/pkg/usecase/input"
	"strconv"
)

/** request queryをマッピングする構造体 */
type Get struct {
	Name     string `query:"name"`
	Age      string `query:"age"`
	IsActive string `query:"is_active"`
}

func (req *Get) ToInput() (input.ListStaffInput, error) {
	input := input.ListStaffInput{
		Name: req.Name,
		Age:  req.Age,
	}

	if req.IsActive != "" {
		isActive, err := strconv.ParseBool(req.IsActive)
		if err != nil {
			return input, errors.New("is_active は true/false で指定してください")
		}
		input.IsActive = &isActive
	}

	return input, nil
}
