package _map

import "sync"

//Set 构造不重复集合
type Set struct {
	//空结构体的内存地址都一样，并且不占用内存空间
	Map  map[int]struct{}
	Len  int
	Lock sync.Mutex
}

//New 初始化列表
func (s *Set) NewSet(cap int) *Set {
	m := make(map[int]struct{}, cap)
	return &Set{
		Map: m,
	}
}

//Add 新增键
func (s *Set) Add(key int) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Map[key] = struct{}{} // 实际往字典添加这个键
	s.Len = len(s.Map)
}

//Remove 移除键
func (s *Set) Remove(key int) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.Len == 0 {
		return
	}
	//移除map里的key
	delete(s.Map, key)
	s.Len = len(s.Map)
}

//Has 查看是否有值
func (s *Set) Has(key int) bool {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	_, ok := s.Map[key]
	return ok
}

//GetLen 获取长度
func (s *Set) GetLen() int {
	return s.Len
}

//IsEmpty 是否为空
func (s *Set) IsEmpty() bool {
	return s.Len == 0
}

//Clear 清空数据
func (s *Set) Clear() {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//释放原来内存
	s.Map = map[int]struct{}{}

	s.Len = 0
}

func (s *Set) List() []int {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	list := make([]int, 0, s.Len)
	for item := range s.Map {
		list = append(list, item)
	}
	return list
}
