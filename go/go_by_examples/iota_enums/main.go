package main

import "fmt"

type ServerState int

const (
	IDLE ServerState = iota
	CONNECTED
	ERROR
	RETRYING
)

var stateNames = map[ServerState]string{
	IDLE:      "idle",
	CONNECTED: "connected",
	ERROR:     "error",
	RETRYING:  "retry",
}

func (ss ServerState) String() string {
	return stateNames[ss]
}
func transition(s ServerState) ServerState {
	switch s {
	case IDLE:
		return CONNECTED
	case CONNECTED, RETRYING:
		return IDLE
	case ERROR:
		return ERROR
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
func main() {
	ns := transition(IDLE)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
	fmt.Println(ns.String())
}
