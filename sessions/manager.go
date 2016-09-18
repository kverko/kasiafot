//Package sessions for sessions management
package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//Manager - struct to keep session ids related to session cookie with given CookieName
type Manager struct {
	lock       sync.Mutex
	CookieName string
	Lifetime   int64
	Sessions   []string
}

//NewManager - session manager contructor
func NewManager(cookieName string, lifetime int64) *Manager {
	newman := new(Manager)
	newman.CookieName = cookieName
	newman.Lifetime = lifetime
	newman.Sessions = make([]string, 0)
	return newman
}

//SessionInit - method to add a session id to sessions manager
func (manager *Manager) SessionInit(sid string) error {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.Sessions = append(manager.Sessions, sid)
	return nil
}

//SessionStart - method that starts a new session and adds it to session manager
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(manager.CookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.newSessionID()
		cookie := http.Cookie{Name: manager.CookieName, Value: url.QueryEscape(sid), MaxAge: int(manager.Lifetime), HttpOnly: true}
		http.SetCookie(w, &cookie)
		err = manager.SessionInit(sid)
	} else {
		sid, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			err = fmt.Errorf("Session manager: couldn't decode cookie: %s\n", manager.CookieName)
		}

		if ok := manager.sessionExists(sid); !ok {
			err = fmt.Errorf("Session manager: session of sid: %s not found\n", sid)
		}
	}
	return err

}

//SessionID - returns ID of session related to given http request
func (manager *Manager) SessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil {
		err = fmt.Errorf("Session Manager: cannot find cookie with name: %s", manager.CookieName)
	}
	sid, err := url.QueryUnescape(cookie.Value)
	return sid, err
}

//RemoveSession - method removing session when needed
func (manager *Manager) RemoveSession(sid string) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	for k, v := range manager.Sessions {
		if v == sid {
			manager.Sessions = append(manager.Sessions[:k], manager.Sessions[k+1:]...)
		}
	}
}

//DelSessionCookie - method deleting cookie related to current session
func (manager *Manager) DelSessionCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(manager.CookieName)
	if (err != nil && err == http.ErrNoCookie) || cookie.Value == "" {
		return nil
	}
	deadCookie := http.Cookie{Name: manager.CookieName, Value: "", Expires: time.Now(), MaxAge: -1}
	http.SetCookie(w, &deadCookie)
	return err
}

//IsLoggedIn - method checking if user sending given http request is logged in
//the user is logged in if there is cookie with proper name and the id of the session
//exists in session manager sessions storage
func (manager *Manager) IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return false
	}

	if sid, _ := manager.SessionID(r); !manager.sessionExists(sid) {
		return false
	}

	return true
}

func (manager *Manager) sessionExists(sid string) bool {
	for _, s := range manager.Sessions {
		if s == sid {
			return true
		}
	}
	return false
}

//create new sess id
func (manager *Manager) newSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
