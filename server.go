package main

type Config struct {
	ListenAddr string
	Strore     Storer
}

type Server struct {
	*Config
	Store Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{Config: cfg}, nil
}
