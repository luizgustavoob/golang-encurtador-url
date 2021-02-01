package url

type memoryRepository struct {
	urls   map[string]*Url
	clicks map[string]int
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		make(map[string]*Url),
		make(map[string]int),
	}
}

func (self *memoryRepository) ExistsID(id string) bool {
	_, existe := self.urls[id]
	return existe
}

func (self *memoryRepository) FindByID(id string) *Url {
	return self.urls[id]
}

func (self *memoryRepository) FindByURL(destino string) *Url {
	for _, u := range self.urls {
		if u.Destination == destino {
			return u
		}
	}
	return nil
}

func (self *memoryRepository) Save(url Url) error {
	self.urls[url.ID] = &url
	return nil
}

func (self *memoryRepository) AddLogClick(id string) {
	self.clicks[id]++
}

func (self *memoryRepository) FindLogClicks(id string) int {
	return self.clicks[id]
}
