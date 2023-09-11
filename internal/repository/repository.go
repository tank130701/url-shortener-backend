package repository

import( 
	"github.com/redis/go-redis/v9"
	"github.com/tank130701/url-shortener-back-end/internal/repository/redis_storage"
	"github.com/tank130701/url-shortener-back-end/internal/repository/sqlite_storage"
	"database/sql"
)


type Storage interface{
	SaveShortUrl(shortURL, fullURL string) error
	GetFullUrl(shortURL string) (string, error)
}

type Repository struct {
	Storage
}

func NewRepositoryRedis(rdb *redis.Client) *Repository {
	return &Repository{
		Storage: redis_storage.NewRedisStorage(rdb),
	}
}

func NewRepositorySqlite(db *sql.DB) *Repository{
	return &Repository{
		Storage: sqlite_storage.NewSqliteStorage(db),
	}
}
