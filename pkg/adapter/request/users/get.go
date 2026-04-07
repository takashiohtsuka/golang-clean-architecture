package users

import (
	"errors"
	"golang-clean-architecture/pkg/usecase/input"
	"strconv"
)

/** request queryをマッピングする構造体 */
type Get struct {
	Name string `query:"name"`
	Age  string `query:"age"`
}

func (req *Get) ToInput() (input.ListUserInput, error) {
	input := input.ListUserInput{
		Name: req.Name,
	}

	if req.Age != "" {
		age, err := strconv.Atoi(req.Age)
		if err != nil {
			return input, errors.New("age は整数で指定してください")
		}
		input.Age = &age
	}

	return input, nil
}
