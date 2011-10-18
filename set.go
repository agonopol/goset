package set

import "json"
import "os"

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (this *Set) Add(x interface{}) {
	this.set[x] = true
}

func (this *Set) Remove(x interface{}) {
	this.set[x] = false, false
}

func (this *Set) Reset() {
	this.set = make(map[interface{}]bool)
}

func (this *Set) Has(x interface{}) bool {
	if _, found := this.set[x]; found {
		return true
	}
	return false
}

func (this *Set) MarshalJSON() ([]byte, os.Error) {
	set := make([]interface{}, 0)
	for k, _ := range this.set {
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
