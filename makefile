development:
	CompileDaemon -build 'go build -buildvcs=false' -command='./example-go-restful-api'

generate:
	rm -rf ./sqlc && sqlc generate

migrate:
	cd ./commands/migrate && go build -buildvcs=false && ./migrate;

rollback:
	cd ./commands/rollback && go build -buildvcs=false && ./rollback;

seed:
	cd ./commands/seed && go build -buildvcs=false && ./seed

drop:
	cd ./commands/drop && go build -buildvcs=false && ./drop

reset: drop generate migrate seed

clean:
	rm -rf ./example-go-restful-api
	rm -rf ./commands/migrate/migrate
	rm -rf ./commands/migrate/rollback
	rm -rf ./commands/migrate/seed
	rm -rf ./commands/migrate/drop
