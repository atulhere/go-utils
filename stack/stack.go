package stack

import(
	"errors"
)

type Stack struct{

	item []interface{}

}

func (s *Stack) Push( item interface{}){

	s.item = append(s.item, item)

}

func (s *Stack) Pop() {

	l := len(s.item)
	if(l >0){
		s.item = s.item[:l-1]
	}
}

func (s *Stack) Top() (interface{}, error) {

	l := len(s.item)
	if(len(s.item)==0){
		return nil, errors.New("The Stack is empty")
	}
	return s.item[l-1], nil
}

func (s *Stack) IsEmpty() bool{

	if(len(s.item) == 0){
		return true
	}
	return false
}
