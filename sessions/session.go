package sessions

//Session - structure for keeping session id and related values
type Session struct {
	id     string
	values map[string]interface{}
}
