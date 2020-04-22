package heap

type Heap interface {
	Insert(int)
	Poll() int
	Remove(int)
	Search(int) int
	Print()
}
