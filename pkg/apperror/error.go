package apperror

// NotFoundException は対象リソースが存在しない場合のエラー
type NotFoundException struct {
	Message string
}

func (e *NotFoundException) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "not found"
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{Message: message}
}
