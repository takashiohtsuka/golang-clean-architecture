package users

import "golang-clean-architecture/pkg/usecase/input"

/** request bodyをマッピングする構造体 */
type Post struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (req *Post) ToInput() input.CreateUserInput {
	return input.CreateUserInput{
		Name: req.Name,
		Age:  req.Age,
	}
}
