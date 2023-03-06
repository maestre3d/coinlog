k8s-create-ns:
	kubectl apply -f deployments/global.yml

k8s-set-ns:
	kubectl config set-context --current --namespace=coinlog

build-kafka-image:
	docker build ./deployments/kafka -t coinlog/kafka-kraft

build-image:
	docker build -f ./deployments/coinlog-http-api/Dockerfile -t coinlog/http-api:latest .

build-web:
	docker build -f ./deployments/coinlog-web-client/Dockerfile -t coinlog/web-client:latest ./client/coinlog-web

k8s-setup: k8s-create-ns k8s-set-ns build-kafka-image
	kubectl apply -f deployments/kafka/kafka.yml && \
	kubectl apply -f deployments/postgres/postgres.yml

set-host:
	echo "127.0.0.1 api.coinlog.info" >> /etc/hosts

gen-di:
	cd di && wire .

append-and-deploy-migration:
	go run -mod=mod ./ent/migrate/main.go "${MIGRATION_NAME}" && \
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:root@coinlog-postgres.coinlog.svc.cluster.local:5432/coinlog?sslmode=disable"

gen-entity:
	ent init "${ENTITY_NAME}"

gen-ent:
	ent generate --feature sql/upsert --feature sql/versioned-migration ./ent/schema

new-migration:
	go run -mod=mod ent/migrate/main.go "${MIGRATION_NAME}" && \
	atlas migrate apply --dir "file://ent/migrate/migrations" --url postgres://postgres:root@coinlog-postgres.coinlog.svc.cluster.local:5432/coinlog?sslmode=disable

new-mock:
	mockery --dir=./domain/user --name=Repository --structname=UserRepository --filename=user_repository.go

gen-coverage:
	go test ./... -coverprofile coverage.out . && go tool cover -html=coverage.out
