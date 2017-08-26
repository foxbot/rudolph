package server

import (
	"fmt"
	"github.com/foxbot/rudolph/generator"
	"net/http"
)

type Server struct {
	host string
	gen  generator.SnowflakeGenerator
}

func NewServer(host string, gen generator.SnowflakeGenerator) Server {
	return Server{
		host: host,
		gen:  gen,
	}
}

func (self *Server) Run() error {
	http.HandleFunc("/snowflake", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, self.gen.Generate())
	})

	return http.ListenAndServe(self.host, nil)
}
