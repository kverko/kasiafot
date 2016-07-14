//session itself
package sessions

import "time"

type Session struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

func (s *Session) Set(key, value interface{}) error {
	s.value[key] = value
	return nil
}
func (s *Session) Get(key interface{}) interface{} {
	if v, ok := s.value[key]; ok {
		return v
	}
	return nil
}
func (s *Session) Delete(key interface{}) error {
	delete(s.value, key)
	return nil
}
func (s *Session) SessionID() string {
	return s.sid
}
