package cherry

// ErrLogger -- interface for logging origin and returned errors due to origin errors discarding
type ErrorLogger interface {
	Log(origin error, returning *Err)
}

// LogOrigin -- logs origin error for returning error using ErrLogger, Chainable.
func (err *Err) LogOrigin(origin error, logger ErrorLogger) *Err {
	logger.Log(origin, err)
	return err
}
