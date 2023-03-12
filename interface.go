package jumpserver

// Executor is jumpserver object action executor.
// It has the ability to create, delete, update, list and get jumpserver objects.
type Executor interface {
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

//var _ = Executor((*AssetExecutor)(nil))

// Object represents Jumpserver object
type Object interface {
	String() string
	PrettyString() string
	GetName() string
	GetID() string
}

// Parameter structure represents url query string.
type Parameter interface {
	Path() (string, error)
	URL(api string) (string, error)
	GetID() string
}
