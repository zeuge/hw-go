package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/tracing"
)

const (
	defaultEntityCap = 64
	tracerName       = "repository"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.PgConfig) (*UserRepository, error) {
	pool, err := pgxpool.New(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	repo := &UserRepository{
		pool: pool,
	}

	return repo, nil
}

func (r *UserRepository) Close() {
	if r.pool != nil {
		r.pool.Close()
	}
}

func (r *UserRepository) Ping(ctx context.Context) error {
	err := r.pool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("r.pool.Ping: %w", err)
	}

	return nil
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "Save")
	defer span.End()

	sql, args, err := goqu.Insert("users").
		Rows(goqu.Record{
			"id":         user.ID.String(),
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		}).
		ToSQL()
	if err != nil {
		return fmt.Errorf("goqu.Insert: %w", err)
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "FindByID")
	defer span.End()

	sql, args, err := goqu.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("users").
		Where(goqu.Ex{
			"id": id.String(),
		}).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu.Select: %w", err)
	}

	user := entity.User{}

	err = r.pool.QueryRow(ctx, sql, args...).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		}

		return nil, fmt.Errorf("r.pool.QueryRow.Scan: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "FindAll")
	defer span.End()

	sql, _, err := goqu.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("users").
		Order(goqu.C("name").Asc()).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("goqu.Select: %w", err)
	}

	rows, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("r.pool.Query: %w", err)
	}
	defer rows.Close()

	users := make([]*entity.User, 0, defaultEntityCap)

	for rows.Next() {
		user := entity.User{}

		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	tracer := tracing.GetTracer(tracerName)

	ctx, span := tracer.Start(ctx, "DeleteByID")
	defer span.End()

	sql, args, err := goqu.Delete("users").
		Where(goqu.Ex{
			"id": id.String(),
		}).
		ToSQL()
	if err != nil {
		return fmt.Errorf("goqu.Delete: %w", err)
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	return nil
}
