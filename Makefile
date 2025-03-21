release:
	docker build --platform linux/amd64 -t toyokoyo199/fitness-backend:2.0.30 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:2.0.30
production:
	docker build --platform linux/amd64 -t toyokoyo199/fitness-production-backend:2.0.30 --build-arg mode=production .
	docker push toyokoyo199/fitness-production-backend:2.0.30

migrate_up_latest:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up
migrate_up:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up 1
migrate_down:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ down 1
migrate_force:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ force 20230411145507
migrate_create:
	migrate create -ext sql -dir migrations update_course_training_avg_statistics_table_course_id_column_v_2_0_29

test-mysql:
	docker-compose up --build -d test-mysql

redis:
	docker-compose up --build -d redis