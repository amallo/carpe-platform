package core

// Command représente une commande exécutable
type Command interface {
	Execute() error
}

