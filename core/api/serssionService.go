package api

import "github.com/gorilla/sessions"

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func (api *apiService) SaveSession() error {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := store.Get(api.Request, "user-session")
	if err != nil {
		return err
	}

	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	// Save it before we write to the response/return from the handler.
	err = session.Save(api.Request, api)
	if err != nil {
		return err
	}

	return nil
}
