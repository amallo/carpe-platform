package core

// EventReducer est une fonction qui réduit un événement en modifiant le state
// Tous les reducers de tous les modules reçoivent tous les événements
// Chaque reducer décide s'il veut réagir à l'événement ou non
// Si le reducer ne réagit pas, il retourne le state inchangé
type EventReducer func(event Event[any], state *State) *State

