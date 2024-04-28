package user

import (
	"reflect"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryUser struct {
	db *pgxpool.Pool
}

func (r RepositoryUser) RepositoryName() string {
	return reflect.TypeOf(r).Name()
}

func (r *RepositoryUser) SetDB(db *pgxpool.Pool) { r.db = db }
func (r *RepositoryUser) GetDB() *pgxpool.Pool   { return r.db }
func (r *RepositoryUser) IsOnlyRead() bool       { return false }
