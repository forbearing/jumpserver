package jumpserver

// Operator is jumpserver object action operator.
// It has the ability to create, delete, update, list and get jumpserver objects.
type Operator interface {
	// Create creates a jumpserver object.
	Create(Object) (Object, error)
	// Delete deletes a jumpserver object.
	Delete(Parameter) error
	// Update updates jumpserver objects.
	Update([]Object) ([]Object, error)
	// UpdatePartial partial updates jumpserver objects.
	UpdatePartial([]Object) ([]Object, error)
	// List lists all jumpserver objects
	List() ([]Object, error)
	// Get gets meet conditions jumpserver objects.
	Get() ([]Object, error)
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

type Logger interface{}

type StructuredLogger interface{}
