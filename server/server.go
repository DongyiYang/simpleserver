package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

// Server is a http.Handler which exposes kubelet functionality over HTTP.
type Server struct {
	mux *http.ServeMux
}

// NewServer initializes and configures a kubelet.Server object to handle HTTP requests.
func NewServer() Server {
	server := Server{
		mux: http.NewServeMux(),
	}
	server.InstallDefaultHandlers()
	return server
}

// ServeHTTP responds to HTTP requests on the Kubelet.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(w, req)
}

// InstallDefaultHandlers registers the default set of supported HTTP request patterns with the mux.
func (s *Server) InstallDefaultHandlers() {
	s.mux.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Web Server\n")
}

// Add parameters passed from command line.
func AddFlags(fs *pflag.FlagSet) {
	return
}

// TODO: For now the address and port number is hardcoded. The actual port number need to be discussed.
func ListenAndServe() {
	glog.V(3).Infof("Start simple web server\n")
	handler := NewServer()
	s := &http.Server{
		Addr:           net.JoinHostPort("0.0.0.0", "9090"),
		Handler:        &handler,
		MaxHeaderBytes: 1 << 20,
	}
	glog.Fatal(s.ListenAndServe())
}
