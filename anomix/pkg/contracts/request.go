package contracts

// Request defines the base request
//go:generate goscinny validator -t Request -f $GOFILE
type Request struct {
	Method    *string `json:"-" required:"true"`
	RequestID *string `json:"-" required:"true"`
}
