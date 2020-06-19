package object

// Env represnets environment in a scope
type Env struct {
	store map[string]Object
	outer *Env
}

// NewEnv instantiates an empty environment
func NewEnv() *Env {
	return &Env{
		store: make(map[string]Object),
		outer: nil,
	}
}

// Get retrieves a symbol
func (e *Env) Get(symbol string) (Object, bool) {
	obj, ok := e.store[symbol]
	if !ok && e.outer != nil {
		return e.outer.Get(symbol)
	}
	return obj, ok
}

// Set assigns a value to the symbol and returns the value
func (e *Env) Set(symbol string, val Object) Object {
	e.store[symbol] = val
	return val
}

// NewEnclosedEnv instantiates a new extended environment
func NewEnclosedEnv(outer *Env) *Env {
	env := NewEnv()
	env.outer = outer
	return env
}
