package core

// EventRouter est une fonction qui route un événement vers une commande
type EventRouter func(event Event[any], deps *Dependencies) Command

type Runtime struct {
	state        *State
	dependencies *Dependencies
	eventQueue   chan Event[any]
	routers      []EventRouter  // Liste de tous les routers de tous les modules
	reducers     []EventReducer // Liste de tous les reducers de tous les modules
}

func NewRuntime(state *State, dependencies *Dependencies, routers []EventRouter, reducers []EventReducer) *Runtime {
	return &Runtime{
		state:        state,
		dependencies: dependencies,
		eventQueue:   make(chan Event[any], 100), // Buffer de 100 événements
		routers:      routers,
		reducers:     reducers,
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
	// Essayer chaque router jusqu'à trouver une commande
	for _, router := range r.routers {
		cmd := router(event, r.dependencies)
		if cmd != nil {
			events := cmd.Execute()
			// Pour chaque événement retourné, passer par tous les reducers
			for _, evt := range events {
				r.applyReducers(evt)
			}
			// Une fois qu'on a trouvé une commande, on arrête (ou on pourrait continuer pour permettre plusieurs commandes)
			break
		}
	}
}

func (r *Runtime) applyReducers(event Event[any]) {
	// Tous les reducers reçoivent l'événement
	for _, reducer := range r.reducers {
		r.state = reducer(event, r.state)
	}
}
