package runtime

// Env represents the environment the application is running in.
type Env string

const (
	Development Env = "dev"
	Test        Env = "test"
	Production  Env = "prod"
)

// String returns the string representation of the environment.
func (e Env) String() string {
	return string(e)
}
