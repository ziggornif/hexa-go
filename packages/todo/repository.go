package todo

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/sirupsen/logrus"
)

// Repository - repository interface
type Repository interface {
	Fetch() ([]Todo, error)
	Insert(todo *Todo) error
	FindByID(id string) (*Todo, error)
	Update(todo *Todo) error
	Delete(id string) error
}

type repository struct {
	db     *pg.DB
	logger *logrus.Logger
}

// NewRepository - create repository instance
func NewRepository(db *pg.DB, logger *logrus.Logger) (Repository, error) {
	repo := &repository{
		db:     db,
		logger: logger,
	}

	err := repo.createSchema()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// createSchema - create table schema
func (r *repository) createSchema() error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := r.db.Model((*Todo)(nil)).CreateTable(opts)
	if createErr != nil {
		r.logger.Errorf("[Repository] createSchema - Error while creating table todo, Reason: %v", createErr)
		return createErr
	}

	r.logger.Info("[Repository] createSchema - Table created")
	return nil
}

// Insert - insert entity
func (r *repository) Insert(todo *Todo) error {
	_, insertErr := r.db.Model(todo).Insert()
	if insertErr != nil {
		r.logger.Errorf("[Repository] Insert - Error while inserting item, Reason: %v", insertErr)
		return insertErr
	}
	r.logger.Info("[Repository] Insert - item created")
	return nil
}

// Fetch - fetch entities
func (r *repository) Fetch() ([]Todo, error) {
	var todos []Todo
	getError := r.db.Model(&todos).Select()
	if getError != nil {
		r.logger.Errorf("[Repository] List - Error while getting items, Reason: %v", getError)
		return nil, getError
	}
	return todos, nil
}

// FindByID - find entity by ID
func (r *repository) FindByID(id string) (*Todo, error) {
	found := Todo{}
	err := r.db.Model(&found).Where("id = ?", id).Select()
	if err != nil {
		r.logger.Errorf("[Repository] FindById - Error while getting item, Reason: %v", err)
		return nil, err
	}
	return &found, nil
}

// Update - update entity
func (r *repository) Update(todo *Todo) error {
	_, err := r.db.Model(todo).WherePK().Update()
	if err != nil {
		r.logger.Errorf("[Repository] Update - Error while inserting item, Reason: %v", err)
		return err
	}
	r.logger.Info("[Repository] Update - item updated")
	return nil
}

// Delete - delete entity
func (r *repository) Delete(id string) error {
	_, err := r.db.Model(&Todo{ID: id}).WherePK().Delete()
	if err != nil {
		r.logger.Errorf("[Repository] Delete - Error while removing item, Reason: %v", err)
		return err
	}
	r.logger.Info("[Repository] Delete - item removed")
	return nil
}
