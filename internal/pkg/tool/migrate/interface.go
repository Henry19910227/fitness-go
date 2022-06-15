package migrate

type Tool interface {
	Up(step *int) error
	Down(step *int) error
	Force(version int) error
	Migrate(version uint) error
	Version() (uint, bool, error)
}
