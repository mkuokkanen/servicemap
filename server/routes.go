package server

import (
	"net/http"
)

func (s *Server) rootPath(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}

func (s *Server) renderIndex(w http.ResponseWriter, r *http.Request) {
	t, d := s.model.TemplateAndData()
	err := t.ExecuteTemplate(w, "main.gohtml", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) reloadData(w http.ResponseWriter, r *http.Request) {
	err := s.model.Load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	_, err = w.Write([]byte("Reloaded"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func staticFileHandler(staticDir string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(staticDir))
	handler := http.StripPrefix("/static", fs)
	return handler.ServeHTTP
}
