package models

type Item struct {
	Id    uint64  `json:"id"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

//Items inmemory store
var Items = []Item{}

var id uint64 = 0

//GetItem finds the item by id and return a link to the item and true if item exist
func GetItem(id uint64) (*Item, bool) {
	for i := range Items {
		if Items[i].Id == id {
			return &Items[i], true
		}
	}
	return &Item{}, false
}

//CreateItem creates a new item and adds it to Items
func (item *Item) CreateItem() error {
	id++
	item.Id = id
	Items = append(Items, *item)
	return nil
}

//DeleteItem finds and deletes Items by Id
func DeleteItem(id uint64) {
	for i := range Items {
		if Items[i].Id == id {
			Items = append(Items[:i], Items[i+1:]...)
			return
		}
	}
}
