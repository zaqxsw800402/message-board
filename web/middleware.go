package main

//func (app *application) Auth(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if !app.Session.Exists(r.Context(), "userID") {
//			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
