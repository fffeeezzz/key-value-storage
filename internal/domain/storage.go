package domain

import (
	"bytes"
	"fmt"
)

type Storage map[string]*Element

var Data = make(Storage)

func (s *Storage) Add(k string, v *Element) bool {
	if k == "" {
		return false
	}

	if _, ok := Data[k]; !ok {
		Data[k] = v
		return true
	}

	return false
}

func (s *Storage) Delete(k string) bool {
	if _, ok := Data[k]; ok {
		delete(Data, k)
		return true
	}

	return false
}

func (s *Storage) Change(k string, v *Element) bool {
	if k == "" {
		return false
	}

	Data[k] = v

	return true
}

func (s *Storage) String() string {
	var buf bytes.Buffer

	for k, v := range Data {
		buf.WriteString(fmt.Sprintf("element: %s with value: %v\n", k, v))
	}

	return buf.String()
}
