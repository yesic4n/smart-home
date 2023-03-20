package interfaces

type IEntity interface {
	DeleteEntity(param string) error
	GetEntity(param string) (interface{}, error)
	UpdateData(msg string, payload interface{}, param string) error
	InsertData(payload interface{}) error
	SetElement(typ string, value interface{}) error
	GetElement(msg string) (*string, error)
	FindDocument(key string, val string) (interface{}, error)
}
