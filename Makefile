include local.env
LOCAL_BIN:=$(CURDIR)/bin

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0


local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

force-start:
	docker compose up -d

CREATE_URL=http://localhost:8080/create/12
REFRESH_URL=http://localhost:8080/refresh
JSON_FILE=$(CURDIR)/response.json

check: fetch send

fetch:
	curl -X POST $(CREATE_URL) -H "Content-Type: application/json" -o $(JSON_FILE)

send:
	curl -X POST $(REFRESH_URL) -H "Content-Type: application/json" --data @$(JSON_FILE)

clean:
	rm -f $(JSON_FILE)