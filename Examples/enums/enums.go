package main
import "fmt"

type ServerState int
 
const (
	StateIdle = iota
	StateConnect 
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle: "Idle",
	StateConnect: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string{
	return stateName[ss]
}

func main() {
	fmt.Println(StateIdle)
	fmt.Println(stateName[StateIdle])
}