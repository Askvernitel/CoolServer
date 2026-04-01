package interfaces

type OperationType int
type OperationCaller int
type Id int

const (
	RUN_RIGHT_OPERATION = iota
	RUN_LEFT_OPERATION
	RUN_UP_OPERATION
	RUN_DOWN_OPERATION
)

type Updatable interface {
	Update([]OperationBundle)
}

type ReadConn interface {
	Read()
}

type WriteConn interface {
	Write(any)
}

type WithData interface {
	GetData() []map[string]any
}

type State interface {
	Updatable
}

type Operation interface {
	GetType() OperationType
}
type OperationBundle interface {
	GetOperations() []Operation
	GetCaller() OperationCaller
}
