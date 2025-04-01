// Package server implements the gRPC service layer for Gai Keep.
//
// This package provides the implementation of the Protocol Buffer service
// defined in service.proto. It handles incoming gRPC requests, performs
// necessary validation and processing, and delegates to the storage layer
// for persistence. The server package acts as the interface between clients
// and the core application functionality.
package server

import (
	"context"
	"fmt"
	"time"

	"github.com/ChristiMahu/gaikeep/internal/storage"
	"github.com/ChristiMahu/gaikeep/pkg/logger"
	pb "github.com/ChristiMahu/gaikeep/proto/gaikeep"
	"github.com/google/uuid"
)

// GaiKeepServer implements the gRPC service defined in the protobuf file.
// This struct satisfies the GaiKeepServer interface generated from the
// Protocol Buffer definition. It wraps a storage implementation and
// provides methods that correspond to the service's RPC endpoints.
type GaiKeepServer struct {
	pb.UnimplementedGaiKeepServer
	store *storage.FileStorage
}

// New returns a new gRPC server instance backed by the provided storage.
// It initializes the server with the given storage implementation, which
// will be used for all persistence operations requested through the gRPC API.
//
// The store parameter is a pointer to a storage.FileStorage instance that
// will handle the actual data persistence.
//
// Returns a configured GaiKeepServer ready to handle gRPC requests.
func New(store *storage.FileStorage) *GaiKeepServer {
	logger.LogInfo("Creating new Gai Keep server instance")
	return &GaiKeepServer{store: store}
}

// SaveEntry handles the gRPC call to persist an entry via the FileStorage backend.
// This method implements the SaveEntry RPC defined in the service.proto file.
// It generates a unique ID and timestamp for the entry, then delegates to
// the storage layer to save the data.
//
// The context parameter provides request scoping and cancellation.
// The req parameter contains the entry content to be saved.
//
// Returns a response with the generated entry ID, or an error if saving fails.
func (s *GaiKeepServer) SaveEntry(ctx context.Context, req *pb.SaveEntryRequest) (*pb.SaveEntryResponse, error) {
	id := uuid.New().String()
	now := time.Now().UTC().Format(time.RFC3339)

	logger.LogInfo("Saving new entry with ID: " + id)

	entry := storage.Entry{
		ID:        id,
		Timestamp: now,
		Content:   req.Content,
	}

	if err := s.store.Save(entry); err != nil {
		logger.LogError("Failed to save entry: " + err.Error())
		return nil, err
	}

	return &pb.SaveEntryResponse{
		Id: id,
	}, nil
}

// GetEntry handles the gRPC call to retrieve an entry by ID.
// This method implements the GetEntry RPC defined in the service.proto file.
// It delegates to the storage layer to retrieve an entry by its unique ID.
//
// The context parameter provides request scoping and cancellation.
// The req parameter contains the ID of the entry to retrieve.
//
// Returns the entry data if found, or an error if retrieval fails.
func (s *GaiKeepServer) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.GetEntryResponse, error) {
	logger.LogInfo("Fetching entry with ID: " + req.Id)

	entry, err := s.store.Get(req.Id)
	if err != nil {
		logger.LogError("Failed to get entry: " + err.Error())
		return nil, err
	}

	return &pb.GetEntryResponse{
		Id:        entry.ID,
		Content:   entry.Content,
		CreatedAt: entry.Timestamp,
		Tags:      entry.Tags,
	}, nil
}

// ListEntries handles the gRPC call to retrieve all entries.
// This method implements the ListEntries RPC defined in the service.proto file.
// It delegates to the storage layer to retrieve all entries and converts them
// to the protocol buffer format expected by clients.
//
// The context parameter provides request scoping and cancellation.
// The req parameter is empty for this method as no inputs are required.
//
// Returns a list of all entries, or an error if retrieval fails.
func (s *GaiKeepServer) ListEntries(ctx context.Context, req *pb.ListEntriesRequest) (*pb.ListEntriesResponse, error) {
	logger.LogInfo("Listing all entries")

	entries, err := s.store.List()
	if err != nil {
		logger.LogError("Failed to list entries: " + err.Error())
		return nil, err
	}

	response := &pb.ListEntriesResponse{
		Entries: make([]*pb.Entry, 0, len(entries)),
	}

	for _, entry := range entries {
		response.Entries = append(response.Entries, &pb.Entry{
			Id:        entry.ID,
			Content:   entry.Content,
			CreatedAt: entry.Timestamp,
			Tags:      entry.Tags,
		})
	}

	logger.LogInfo(fmt.Sprintf("Retrieved %d entries", len(entries)))
	return response, nil
}
