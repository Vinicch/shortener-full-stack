up:
	docker compose up -d
	
up.prod:
	docker compose --profile release up -d --build

down:
	docker compose down --remove-orphans

logs:
	docker compose logs -tf
