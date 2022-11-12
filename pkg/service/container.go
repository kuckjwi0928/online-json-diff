package service

type Container struct {
	diffService DifferService
	httpService HttpService
}

func NewServiceContainer() *Container {
	return &Container{
		diffService: NewGoJsonDiffer(),
		httpService: NewHttpService(),
	}
}

func (c *Container) DiffService() DifferService {
	return c.diffService
}

func (c *Container) HttpService() HttpService {
	return c.httpService
}
