package commons

type FailedTypeCastError struct{}

func (FailedTypeCastError) Error() string {
	return "型のキャスト先が誤っています"
}
