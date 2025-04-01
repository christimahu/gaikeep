// Package app contains the core application logic and server management
//
// This package coordinates the various components of the Gai Keep application,
// including configuration loading, storage initialization, and gRPC server setup.
// It serves as the central hub that ties together all the different parts of the
// system and manages their lifecycle.
package app

import (
	"net"

	"github.com/ChristiMahu/gaikeep/internal/config"
	"github.com/ChristiMahu/gaikeep/internal/server"
	"github.com/ChristiMahu/gaikeep/internal/storage"
	"github.com/ChristiMahu/gaikeep/pkg/logger"
	"github.com/ChristiMahu/gaikeep/proto/gaikeep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents the Gai Keep application server
// It encapsulates all the components needed to run the service, including
// configuration, storage, and the gRPC server itself. The Server handles
// the complete lifecycle from initialization to shutdown.
type Server struct {
	config  *config.Config
	storage *storage.FileStorage
	grpc    *grpc.Server
	address string
}

// NewServer creates a new application server instance
//
// It loads configuration, initializes the storage layer, and sets up
// the gRPC server with appropriate service handlers. The address parameter
// specifies the network address and port to listen on (e.g., ":50051").
//
// Returns an error if any initialization step fails.
func NewServer(address string) (*Server, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	store, err := storage.NewFileStorage(cfg.Directory)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	gaikeep.RegisterGaiKeepServer(s, server.New(store))
	reflection.Register(s)

	return &Server{
		config:  cfg,
		storage: store,
		grpc:    s,
		address: address,
	}, nil
}

// Start begins listening and serving requests
//
// This method blocks until the server is stopped or an error occurs.
// It creates a TCP listener on the server's configured address and
// starts the gRPC server on that listener.
//
// Returns an error if the listener cannot be created or the server
// fails to start.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	logger.LogInfo("Gai Keep server listening on " + s.address)
	return s.grpc.Serve(lis)
}

// Stop gracefully shuts down the server
//
// This method signals the gRPC server to stop accepting new requests and
// to finish processing existing requests before shutting down. It should
// be called when the application is terminating to ensure clean shutdown.
func (s *Server) Stop() {
	if s.grpc != nil {
		s.grpc.GracefulStop()
	}
}
