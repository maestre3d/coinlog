gen-di:
	cd di && wire .

append-and-deploy-migration:
	go run -mod=mod ./ent/migrate/main.go "${MIGRATION_NAME}" && \
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:root@localhost:6432/coinlog?sslmode=disable"


gen-entity:
	ent init "${ENTITY_NAME}"

gen-ent:
	ent generate --feature sql/upsert --feature sql/versioned-migration ./ent/schema

new-migration:
	go run -mod=mod ent/migrate/main.go "${MIGRATION_NAME}" && \
	atlas migrate apply --dir "file://ent/migrate/migrations" --url postgres://postgres:root@localhost:6432/coinlog?sslmode=disable