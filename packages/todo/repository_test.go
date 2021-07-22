package todo

import (
	"hexa-go/infra"
	"hexa-go/infra/config"
	"hexa-go/infra/storage"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/stretchr/testify/assert"
)

func cleanInserted(db *pg.DB, id string) {
	_, _ = db.Model(&Todo{}).Where("id = ?", id).Delete()
}

func TestNewRepoError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)
	_ = db.Close()

	_, err := NewRepository(db, logger)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "pg: database is closed")
	storage.DBClose()
}

func TestInsert(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	assert.Nil(t, err)

	cleanInserted(db, "1234")
}

func TestInsertError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)
	_ = db.Model((*Todo)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ERROR #42P01 relation \"todos\" does not exist")
}

func TestFetch(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	assert.Nil(t, err)

	todos, _ := repo.Fetch()

	assert.Equal(t, len(todos), 1)
	assert.Equal(t, todos[0].ID, "1234")
	assert.Equal(t, todos[0].Title, "Test")
	assert.Equal(t, todos[0].Completed, false)
	assert.Equal(t, todos[0].Order, 0)

	cleanInserted(db, "1234")
}

func TestFetchError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)
	_ = db.Model((*Todo)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})

	_, err := repo.Fetch()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ERROR #42P01 relation \"todos\" does not exist")
}

func TestFindByID(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	assert.Nil(t, err)

	created, _ := repo.FindByID("1234")
	assert.Equal(t, created.ID, "1234")
	assert.Equal(t, created.Title, "Test")
	assert.Equal(t, created.Completed, false)
	assert.Equal(t, created.Order, 0)

	cleanInserted(db, "1234")
}

func TestFindByIDError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)
	_ = db.Model((*Todo)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})

	_, err := repo.FindByID("1234")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ERROR #42P01 relation \"todos\" does not exist")
}

func TestUpdate(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	assert.Nil(t, err)

	err = repo.Update(&Todo{
		ID:        "1234",
		Title:     "TestUpdated",
		Completed: true,
		Order:     1,
	})
	assert.Nil(t, err)

	updated, _ := repo.FindByID("1234")
	assert.Equal(t, updated.ID, "1234")
	assert.Equal(t, updated.Title, "TestUpdated")
	assert.Equal(t, updated.Completed, true)
	assert.Equal(t, updated.Order, 1)

	cleanInserted(db, "1234")
}

func TestUpdateError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)
	_ = db.Model((*Todo)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})

	err := repo.Update(&Todo{
		ID:        "1234",
		Title:     "TestUpdated",
		Completed: true,
		Order:     1,
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ERROR #42P01 relation \"todos\" does not exist")
}

func TestDelete(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)

	err := repo.Insert(&Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	assert.Nil(t, err)

	err = repo.Delete("1234")
	assert.Nil(t, err)

	res, _ := repo.FindByID("1234")
	assert.Empty(t, res)
}

func TestDeleteError(t *testing.T) {
	logger := infra.GetLogger()
	conf, _ := config.LoadConfig("../..", logger)
	db, _ := storage.DBConnect(conf.GetConfig(), logger)

	repo, _ := NewRepository(db, logger)
	_ = db.Model((*Todo)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})

	err := repo.Delete("1234")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ERROR #42P01 relation \"todos\" does not exist")
}
