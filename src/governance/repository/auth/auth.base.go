package auth

import (
	"reflect"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryAuth struct {
	db *pgxpool.Pool
}

func (r RepositoryAuth) RepositoryName() string {
	return reflect.TypeOf(r).Name()
}

func (r *RepositoryAuth) SetDB(db *pgxpool.Pool) { r.db = db }
func (r *RepositoryAuth) GetDB() *pgxpool.Pool   { return r.db }
func (r *RepositoryAuth) IsOnlyRead() bool       { return false }
