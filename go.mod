module github.com/Xevion/todoist-late-reset

go 1.22.3

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-co-op/gocron/v2 v2.12.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
)


require internal/api v1.0.0
replace internal/api => ./internal/api