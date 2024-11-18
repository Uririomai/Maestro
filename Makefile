compose_up:
	docker compose --env-file .\backend\.env up -d

compose_down:
	docker compose down

compose_drop:
	docker compose down -v

compose_rebuild:
	docker compose down -v
	docker compose --env-file .\backend\.env up -d --build --force-recreate

