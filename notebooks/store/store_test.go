package store

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/influxdata/influxdb/v2/sqlite"
	sqliteMigrations "github.com/influxdata/influxdb/v2/sqlite/migrations"
	influxdbtesting "github.com/influxdata/influxdb/v2/testing"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetNotebook(t *testing.T) {
	t.Parallel()

	store, _ := sqlite.NewTestStore(t)
	// defer clean(t)
	ctx := context.Background()

	sqliteMigrator := sqlite.NewMigrator(store, zap.NewNop())
	err := sqliteMigrator.Up(ctx, &sqliteMigrations.All{})
	require.NoError(t, err)

	service, err := NewService(zap.NewNop(), store)
	require.NoError(t, err)

	var wg sync.WaitGroup

	for i := 0; i < 9999; i++ {
		wg.Add(1)
		go func() {
			_, err := service.GetNotebook(ctx, *influxdbtesting.IDPtr(1))
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
