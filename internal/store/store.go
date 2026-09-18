// the key-vallue expire 
// store is the database 

/* 
*/ 
package store 

import (
	"time"
)

type Store struct{
	data map[string]entry
}

type entry struct {
	value string 
	expiresAt time.Time 
}

// go versoin for  self in py
// attachinng this method to the store, thats what the func (s *Store) FunctionName(param1, param2)
func (s *Store) Set(key string, value string) {
	s.data[key] = entry{value: value}
}


func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val.value, ok
}

func (s *Store) Exist(key string) (bool) {
	_ , ok := s.data[key]
	return ok
}

///variadic parameterm any number of key strings
func (s *Store) Del(keys ...string) (int) {

	count := 0 
	for _, key := range keys {
		if _, ok := s.data[key]; ok {
			delete(s.data, key)
			count ++
		}
	}
	return count 
} 

func New() *Store{
	return &Store{
		data : make(map[string]entry)}
}
