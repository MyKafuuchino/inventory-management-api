MIGRATE=migrate
MIGRATION_PATH=./database/migration
DB_URL=mysql://root@tcp(localhost:3306)/inventory_management

build:
	@if [ ! -d "dist" ]; then mkdir dist; fi
	go build -o dist/myapp main.go

run:
	@./dist/myapp

migration:
	@$(MIGRATE) create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migrate-up:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migrate-down:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" down $(steps)

migrate-status:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" version

migrate-reset:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" down
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" up

