up:
	docker compose up -d
	
deploy:
	docker compose --profile release up -d --build

down:
	docker compose down --remove-orphans

logs:
	docker compose logs -tf
