package core

// EventRouter est une fonction qui route un événement vers une commande
type EventRouter func(event Event[any], deps *Dependencies, state *State) Command

type Runtime struct {
	state        *State
	dependencies *Dependencies
	eventQueue   chan Event[any]
	router       EventRouter
}

func NewRuntime(state *State, dependencies *Dependencies, router EventRouter) *Runtime {
	return &Runtime{
		state:        state,
		dependencies: dependencies,
		eventQueue:   make(chan Event[any], 100), // Buffer de 100 événements
		router:       router,
	}
}

func (r *Runtime) Send(event Event[any]) {
	select {
	case r.eventQueue <- event:
		// Événement envoyé
	default:
		// Channel plein (optionnel: log ou erreur)
	}
}

// State retourne l'état actuel du runtime
func (r *Runtime) State() *State {
	return r.state
}

func (r *Runtime) RunUntilIdle() {
	for {
		select {
		case event := <-r.eventQueue:
			r.handleEvent(event)
		default:
			// Channel vide, on arrête
			return
		}
	}
}

func (r *Runtime) handleEvent(event Event[any]) {
	cmd := r.router(event, r.dependencies, r.state)
	if cmd != nil {
		cmd.Execute()
	}
}
