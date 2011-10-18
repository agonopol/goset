package goset

import "testing"
import "json"
import "fmt"


func compare(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("value differs. Expected [%v], actual [%v]", expected, actual)
		panic(fmt.Sprintf("value differs. Expected [%v], actual [%v]", expected, actual))
	}
}

func TestSetAddRemoveContains(t *testing.T) {
	set := NewSet()
	set.Add(10)
	compare(t, true, set.Has(10))
	set.Remove(10)
	compare(t, false, set.Has(10))
	compare(t, false, set.Has("shie"))

}

func TestSetMarshalAndUnMarshal(t *testing.T) {
	set := NewSet()
	set.Add("hello")
	set.Add("there")
	set.Add("you")
	set.Add("shie")
	compare(t, true, set.Has("shie"))
	marshaled, err := json.Marshal(set)
	if err != nil {
		panic(err)
	}
	compare(t, "[\"hello\",\"there\",\"shie\",\"you\"]", string(marshaled))
	unmarshaled := NewSet()
	err = json.Unmarshal(marshaled, unmarshaled)
	compare(t, true, unmarshaled.Has("shie"))

}