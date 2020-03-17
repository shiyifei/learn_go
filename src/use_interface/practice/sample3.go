package practice

import "sort"

type SortInterface interface {
	Len() int
	Less(i int ,j int) bool
	Swap(i,j int)
}

type Sortable interface {
	sort.Interface
	Sort()
}


type SortableStrs [3]string

func (this *SortableStrs) Len() int {
	return len(this)
}

func (this *SortableStrs) Less(i,j int) bool {
	return this[i] < this[j]
}

func (this *SortableStrs) Swap(i,j int ) {
	this[i], this[j] = this[j], this[i]
}

func (this *SortableStrs) Sort() {
	sort.Sort(this)
}




