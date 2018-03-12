package cherry

type ErrConstruct func(...func(*Err)) *Err

func (constr ErrConstruct) Error() string {
	return constr().Error()
}

func (constr ErrConstruct) AddDetails(details ...string) ErrConstruct {
	err := constr().AddDetails(details...)
	return func(options ...func(*Err)) *Err {
		for _, option := range options {
			option(err)
		}
		return err
	}
}

func (constr ErrConstruct) AddDetailsErr(details ...error) ErrConstruct {
	err := constr().AddDetailsErr(details...)
	return func(options ...func(*Err)) *Err {
		for _, option := range options {
			option(err)
		}
		return err
	}
}

func (constr ErrConstruct) AddDetailF(f string, vals ...interface{}) ErrConstruct {
	err := constr().AddDetailF(f, vals...)
	return func(options ...func(*Err)) *Err {
		for _, option := range options {
			option(err)
		}
		return err
	}
}
