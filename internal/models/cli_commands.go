package models

type CLICommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}
