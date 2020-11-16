package ctxt

var globalCtxt = Default()

// Returns the current Context used in all operations that do not receive a Context
// as part of its parameters.
func GlobalCtxt() Ctxt {
	return globalCtxt
}

// Sets the Context that will be used by all operations that do not receive a new Context as part
// of its parameters..
func SetGlobalCtxt(ctx Ctxt) {
	globalCtxt = ctx
}

// Re-Initializes the global Context to its default value
func ResetGlobalCtxt() {
	globalCtxt = Default()
}
