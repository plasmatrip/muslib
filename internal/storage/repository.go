package storage

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/plasmatrip/muslib/internal/logger"
	"github.com/plasmatrip/muslib/internal/model"
	"github.com/plasmatrip/muslib/internal/storage/queries"
)

type Repository struct {
	db  *pgxpool.Pool
	log logger.Logger
}

func NewRepository(ctx context.Context, dsn string, log logger.Logger) (*Repository, error) {
	// запускаем миграцию
	err := StartMigration(dsn)
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		} else {
			log.Sugar.Infoln("the database exists, there is nothing to migrate")
		}
	} else {
		log.Sugar.Infoln("database migration was successful")
	}

	// открываем БД
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:  db,
		log: log,
	}, nil
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func StartMigration(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
		return err
	}
	return nil
}

func (r Repository) Ping(ctx context.Context) error {
	return r.db.Ping(ctx)
}

func (r Repository) Close() {
	r.db.Close()
}

func (r Repository) AddSong(ctx context.Context, song model.Song) error {
	ct, err := r.db.Exec(ctx, queries.AddSong, pgx.NamedArgs{
		"group_name":   song.Group,
		"song_name":    song.Song,
		"release_date": time.Time(song.ReleaseDate),
		"lyrics":       song.Text,
		"link":         song.Link,
	})
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		r.log.Sugar.Debugw("song not added", "group", song.Group, "song", song.Song)
		return errors.New("song not added")
	}

	return nil
}

func (r Repository) DeleteSong(ctx context.Context, song model.Song) error {
	ct, err := r.db.Exec(ctx, queries.DeleteSong, pgx.NamedArgs{
		"group_name": song.Group,
		"song_name":  song.Song,
	})
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		r.log.Sugar.Debugw("song not deleted", "group", song.Group, "song", song.Song)
		return errors.New("song not deleted")
	}

	return nil
}

func (r Repository) UpdateSong(ctx context.Context, song model.Song) error {
	ct, err := r.db.Exec(ctx, queries.UpdateSong, pgx.NamedArgs{
		"group_name":   song.Group,
		"song_name":    song.Song,
		"release_date": song.ReleaseDate.NilIfZero(),
		"lyrics":       song.Text,
		"link":         song.Link,
	})
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		r.log.Sugar.Debugw("song not updated", "group", song.Group, "song", song.Song, "release_date", song.ReleaseDate, "lyrics", song.Text, "link", song.Link)
		return errors.New("song not updated")
	}

	return nil
}
