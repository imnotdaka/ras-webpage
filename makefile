start-docker:
	PROMETHEUS_USER= PROMETHEUS_PASSWORD= MYSQL_PASSWORD= docker-compose up -d --build

start-back:
	cd backend && go run ./...

start-front:
	cd frontend && npm run dev