package set

import "json"
import "os"

type Set map[interface{}]bool

func New() *Set {
	set := make(Set)
	return &set
}

func (this *Set) Add(x interface{}) {
	(*this)[x] = true
}

func (this *Set) Remove(x interface{}) {
	(*this)[x] = false, false
}

func (this *Set) Reset() {
	(*this) = make(map[interface{}]bool)
}

func (this *Set) Has(x interface{}) bool {
	_, found := (*this)[x]
	return found
}

func (this *Set) MarshalJSON() ([]byte, os.Error) {
	set := make([]interface{}, 0)
	for k, _ := range *this {
		set = append(set, k)
	}
	return json.Marshal(set)
}

func (this *Set) UnmarshalJSON(body []byte) os.Error {
	var set []interface{}
	err := json.Unmarshal(body, &set)
	if err != nil {
		return err
	} else {
		for _, v := range set {
			this.Add(v)
		}
	}
	return nil
}
