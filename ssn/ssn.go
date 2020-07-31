package ssn

import (
    "github.com/gorilla/sessions"
    "net/http"
)

type Req = http.Request
type Resp = http.ResponseWriter

type Sessioner interface {
    Session(r *Req) *sessions.Session
    IsAuthenticated(r *Req) bool
    ClearSession(w Resp, r *Req) error
    Authenticated(r *Req)
    Put(key string, v interface{}, r *Req) *sessions.Session
    SaveThisSession(w Resp, r *Req) error
    SaveSession(w Resp, r *Req) error
    DeleteSession(w Resp, r *Req) error
}
