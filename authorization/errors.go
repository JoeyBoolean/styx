package authorization

// Types
type (
	errNotFound interface {
		error
		IsErrNotFound()
	}
	errDeniedAccess interface {
		error
		IsDeniedAccess()
	}
)

// Behaviors
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
