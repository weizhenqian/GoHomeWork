package sort

type Person struct {
	Fristname string
	Lastname  string
}

type Persons []Person

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Fristname > p[j].Fristname
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(data Persons) {
	Plen := data.Len()
	for pass := 1; pass < Plen; pass++ {
		for i := 0; i < Plen-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}
