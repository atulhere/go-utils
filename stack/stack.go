package stack

import(
	"errors"
)

type Stack struct{

	Item []interface{}

}

func (s *Stack) Push( item interface{}){

	s.Item = append(s.Item, item)

}

func (s *Stack) Pop() {

	l := len(s.Item)
	if(l >0){
		s.Item = s.Item[:l-1]
	}
}

func (s *Stack) Top() (interface{}, error) {

	l := len(s.Item)
	if(len(s.Item)==0){
		return nil, errors.New("The Stack is empty")
	}
	return s.Item[l-1], nil
}

func (s *Stack) IsEmpty() bool{

	if(len(s.Item) == 0){
		return true
	}
	return false
}
