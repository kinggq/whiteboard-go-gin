package core

import "github.com/kinggq/initialize"

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routes()

	s := initServer(":8088", Router)
	s.ListenAndServe()
}
