package interfaces

type OperationType int
type OperationCaller string
type Id string

const (
	RUN_RIGHT_OPERATION = iota
	RUN_LEFT_OPERATION
	RUN_UP_OPERATION
	RUN_DOWN_OPERATION
)

type Updatable interface {
	Update([]Operation)
}

type State interface {
	Updatable
}

type Operation interface {
	GetType() OperationType
	GetCaller() OperationCaller
}
