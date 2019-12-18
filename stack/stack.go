package stack

import (
	//"fmt"
)

// IsEmpty Pop Push Top

type Stack struct {
	Data		[]interface{}
	minData		[]map[interface{}]int
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *Stack) Push(data interface{}) {
	s.Data = append(s.Data, data)

	func() {
		if len(s.minData) == 0 {
			tmpMap := make(map[interface{}]int)
			tmpMap[data] = 1
			s.minData = append(s.minData, tmpMap)
			return
		} 

		if _, ok := s.minData[len(s.minData) - 1][data]; ok {
			// top is min
			s.minData[len(s.minData) - 1][data]++
			return
		}

		for tmpMin, _ := range s.minData[len(s.minData) - 1] {
			if tmpMin.(int) > data.(int) {
				tmpMap := make(map[interface{}]int)
				tmpMap[data] = 1
				s.minData = append(s.minData, tmpMap)
				return		
			}
		}
	}()
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	tmpData := s.Data[len(s.Data) - 1]
	s.Data = s.Data[:len(s.Data) - 1]

	if tmpMin, ok := s.minData[len(s.minData) - 1][tmpData]; ok {
		if tmpMin == 1 {
			s.minData = s.minData[:len(s.minData) - 1]
		} else {
			s.minData[len(s.minData) - 1][tmpData]--
		}
	}

	return tmpData
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.Data[len(s.Data) - 1]
}

func (s *Stack) Min() interface{} {
	if s.IsEmpty() {
		return nil
	}

	for minData, _ := range s.minData[len(s.minData) - 1] {
		return minData
	}

	return nil
}

func main() {
	ms := Stack{Data:make([]interface{}, 0, 2), minData:make([]map[interface{}]int, 0, 1)}
	ms.Push(10)
	ms.Push(14)
	ms.Push(18)
	ms.Push(5)
	ms.Push(11)
	ms.Push(5)
	ms.Pop()
	ms.Pop()
	ms.Pop()
	ms.Pop()
	fmt.Println(ms.Min())
}













