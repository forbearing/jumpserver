package jumpserver

// Operator is jumpserver object action operator.
// It has the ability to create, delete, update, list and get jumpserver objects.
type Operator interface {
	// Create creates a jumpserver object.
	Create(...[]Object) (Object, error)
	// Delete deletes a jumpserver object.
	Delete(Parameter) error
	// Update updates jumpserver objects.
	Update(...[]Object) ([]Object, error)
	// UpdatePartial partial updates jumpserver objects.
	UpdatePartial(...[]Object) ([]Object, error)
	// List lists all jumpserver objects
	List(Parameter) ([]Object, error)
	// Get gets meet conditions jumpserver objects.
	Get(id string) ([]Object, error)
}

//var _ = Operator((*AssetExecutor)(nil))

// Object represents Jumpserver object
type Object interface {
	String() string
	PrettyString() string
	GetID() string
}

// Parameter structure represents url query string.
type Parameter interface {
	Query() (string, error)
	URL(api string) (string, error)
	GetID() string
}

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

// StructuredLogger is structured logger interface used for user to custom themselves logger.
type StructuredLogger interface {
	Debugw(msg string, args ...any)
	Infow(msg string, args ...any)
	Warnw(msg string, args ...any)
	Errorw(msg string, args ...any)
	Fatalw(msg string, args ...any)
}
