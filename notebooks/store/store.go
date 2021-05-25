package store

import (
	"context"
	"database/sql"
	"sync"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/influxdata/influxdb/v2"
	"github.com/influxdata/influxdb/v2/kit/platform"
	"github.com/influxdata/influxdb/v2/notebooks/service"
	"github.com/influxdata/influxdb/v2/snowflake"
	"github.com/influxdata/influxdb/v2/sqlite"
	"github.com/upper/db/v4"
	"go.uber.org/zap"
)

var _ service.NotebookService = (*Store)(nil)

type Store struct {
	mu            sync.Mutex
	sess          db.Session
	db            *sql.DB
	log           *zap.Logger
	timeGenerator influxdb.TimeGenerator
	idGenerator   platform.IDGenerator
}

func NewService(logger *zap.Logger, s *sqlite.SqlStore) (*Store, error) {
	return &Store{
		sess:          s.Sess,
		db:            s.DB,
		log:           logger,
		timeGenerator: influxdb.RealTimeGenerator{},
		idGenerator:   snowflake.NewIDGenerator(),
	}, nil
}

func (s *Store) GetNotebook(ctx context.Context, id platform.ID) (*service.Notebook, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.sess.SQL().
		InsertInto("notebooks").
		Columns(
			"id",
			"org_id",
			"name",
			"spec",
		).
		Values(
			s.idGenerator.ID(),
			s.idGenerator.ID(),
			"some name",
			"some spec",
		).
		Exec()

	return nil, err
}

func (s *Store) CreateNotebook(ctx context.Context, create *service.NotebookReqBody) (*service.Notebook, error) {

	return nil, nil
}

func (s *Store) UpdateNotebook(ctx context.Context, id platform.ID, update *service.NotebookReqBody) (*service.Notebook, error) {

	return nil, nil
}

func (s *Store) DeleteNotebook(ctx context.Context, id platform.ID) error {

	return nil
}

func (s *Store) ListNotebooks(ctx context.Context, filter service.NotebookListFilter) ([]*service.Notebook, error) {

	return nil, nil
}
