package interactor

type Application interface {
	CreateEvent(Event) (Event, error)
	GetEvent(int64) (Event, error)
	ListEvents() ([]Event, error)
}

type app struct {
	store DataGateway
}

func NewApplication(s DataGateway) Application {
	return app{store: s}
}

func (a app) CreateEvent(event Event) (Event, error) {
	return a.store.Create(event)
}

func (a app) GetEvent(id int64) (Event, error) {
	return a.store.GetByID(id)
}

func (a app) ListEvents() ([]Event, error) {
	return a.store.List()
}
