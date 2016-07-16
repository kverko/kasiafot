//package for sessions management
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

type Manager struct {
	lock       sync.Mutex
	CookieName string
	Lifetime   int64
	Sessions   map[string]Session
}

func NewManager(cookieName string, lifetime int64) *Manager {
	newman := new(Manager)
	newman.CookieName = cookieName
	newman.Lifetime = lifetime
	newman.Sessions = make(map[string]Session, 0)
	return newman
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionInit(sid string) (Session, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	sess := Session{sid: sid, timeAccessed: time.Now()}
	manager.Sessions[sid] = sess
	return sess, nil
}

func (manager *Manager) SessionUse(sid string) (Session, bool) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	sess, ok := manager.Sessions[sid]
	if !ok {
		return sess, false
	}
	sess.timeAccessed = time.Now()
	return sess, true
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	cookie, err := r.Cookie(manager.CookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		cookie := http.Cookie{Name: manager.CookieName, Value: url.QueryEscape(sid), MaxAge: int(manager.Lifetime)}
		http.SetCookie(w, &cookie)
		session, err = manager.SessionInit(sid)
	} else {
		sid, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			err = fmt.Errorf("Session manager: no cookie with %s", manager.CookieName)
		}
		ok := true
		session, ok = manager.SessionUse(sid)
		if ok != true {
			err = fmt.Errorf("Session manager: session of sid: %s not found", sid)
		}
	}

	return
}

func (manager *Manager) IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie(manager.CookieName)
	if err != nil || cookie.Value == "" {
		return false
	}

	return true
}
