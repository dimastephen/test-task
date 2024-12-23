include local.env
LOCAL_BIN:=$(CURDIR)/bin

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0


local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

force-start:
	docker compose up -d
	curl --request POST -sL \
	     --url 'http://example.com'\
	     --output './path/to/file'
