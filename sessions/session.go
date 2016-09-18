package sessions

//session itself: structure and functionality

type Session struct {
	id     string
	values map[string]interface{}
}
