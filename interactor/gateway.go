package interactor

type DataGateway interface {
	Create(Event) (Event, error)
	GetByID(id int64) (Event, error)
	List() ([]Event, error)
	Update(Event) error
}
