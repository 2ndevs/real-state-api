package utils

type WatchedList[T comparable] struct {
	initialState []T
	items        []T
	removed      []T
	added        []T
}

func (self *WatchedList[T]) Create(items []T) {
	self.items = items
	self.initialState = items
	self.added = make([]T, 1)
	self.removed = make([]T, 1)
}

func (self *WatchedList[T]) Add(newItem T) {
	for _, item := range self.items {
		if item == newItem {
			return
		}
	}

	self.items = append(self.items, newItem)
	self.added = append(self.added, newItem)
}

func (self *WatchedList[T]) Remove(itemToRemove T) {
	for idx, item := range self.items {
		if item != itemToRemove {
			continue
		}

		self.items = append(self.items[:idx], self.items[idx+1:]...)
		self.removed = append(self.removed, itemToRemove)
	}
}

func (self WatchedList[T]) GetRemoved() []T {
	return self.removed
}

func (self WatchedList[T]) GetAdded() []T {
	return self.added
}

func (self WatchedList[T]) GetItems() []T {
	return self.items
}

func (self WatchedList[T]) GetInitialState() []T {
	return self.initialState
}
