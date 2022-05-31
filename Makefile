up:
	docker compose up -d
	
up.prod:
	docker compose up --profile release -d

down:
	docker compose down --remove-orphans

logs:
	docker compose logs -tf
