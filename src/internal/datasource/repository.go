// package datasource
package datasource

type Repository interface {
	SaveGame() error
	GetGame() error
}
