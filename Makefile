compose_up:
	docker compose --env-file .\backend\.env --env-file .\frontend\.env up -d

compose_down:
	docker compose down

compose_drop:
	docker compose down -v