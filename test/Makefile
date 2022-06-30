test:
	docker compose up --build --force-recreate --always-recreate-deps --renew-anon-volumes -d
	docker exec test-mysqlrouter-go-test-1 sh -c "bash /go/src/mysqlrouter-go/entrypoint.sh"

down:
	docker compose down
	docker compose rm -f

local: test down