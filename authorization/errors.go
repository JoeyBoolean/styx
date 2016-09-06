package authorization

type errDeniedAccess interface {
	error
	IsDeniedAccess()
}

type errDeniedAccessBehavior struct{}

func (err errDeniedAccessBehavior) IsDeniedAccess() {}

func withErrDeniedAccess(err error) error {
	return struct {
		error
		errDeniedAccessBehavior
	}{
		err,
		errDeniedAccessBehavior{},
	}
}

func isErrDeniedAccess(err error) bool {
	_, ok := err.(errDeniedAccess)
	return ok
}
