package domain

// CodeLength represents the amount of digits in a code
const CodeLength = 4

type CodeUsecase interface {
	SendRandomCodeAndWaitForPair(writer CodeWriter) error
	Pair(code string) error
}

type CodeRepository interface {
	CreateCodesIfNeeded() error
	GetRandomCode() (string, error)
}

// CodePubSub can publish a code to- and subscribe to a channel
type CodePubSub interface {
	Publish(code string) error
	Subscribe(code string) <-chan bool
}

// CodeReader reads a code from the client
type CodeReader interface {
	Read() (string, error)
}

// CodeWriter writes a code to the client
type CodeWriter interface {
	Write(code string) error
}
