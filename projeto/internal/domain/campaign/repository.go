package campaign

type Repository interface {
	Create(c *Campaign) error
	Get() ([]Campaign, error)
	GetStarted() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
	Update(c *Campaign) error
	Delete(c *Campaign) error
}
