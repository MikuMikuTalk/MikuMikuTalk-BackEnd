package set

import "fmt"

type void struct{}
type Set struct {
	m map[any]void
}

// 创建一个新的集合
func NewSet(items ...any) *Set {
	s := &Set{}
	s.m = make(map[any]void)
	s.Add(items...)
	return s
}

// 向集合中添加元素
func (s *Set) Add(items ...any) {
	for _, item := range items {
		s.m[item] = void{}
	}
}

// 判断集合中是否包含item
func (s *Set) Contains(item any) bool {
	_, ok := s.m[item]
	return ok
}

// 返回集合大小
func (s *Set) Size() int {
	return len(s.m)
}

// 移除集合元素
func (s *Set) Remove(item any) {
	delete(s.m, item)
}

// 清空集合
func (s *Set) Clear() {
	s.m = make(map[any]void)
}

// 判断两个集合是否相等
func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.m {
		// 只要有一个不相等，就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) IsSubSet(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// 求并集
func Union(set1 *Set, set2 *Set) *Set {
	newSet := NewSet()
	for key := range set1.m {
		newSet.Add(key)
	}
	for key := range set2.m {
		newSet.Add(key)
	}
	return newSet
}

// 求交集
func InterSet(set1, set2 *Set) *Set {
	newSet := NewSet()
	for key := range set1.m {
		if set2.Contains(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

// 求差集
// 差集（A - B）：所有属于集合 A 但不属于集合 B 的元素。
func Difference(set1, set2 *Set) *Set {
	newSet := NewSet()
	for key := range set1.m {
		if !set2.Contains(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (s *Set) String() string {
	str := ""
	for key, _ := range s.m {
		str += fmt.Sprintf("%v,", key)
	}
	if len(s.m) > 0 {
		str = str[:len(str)-1] // 去掉最后的逗号和空格
	}
	return str
}
