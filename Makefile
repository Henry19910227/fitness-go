release:
	docker build -t toyokoyo199/fitness-backend:2.0.4 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:2.0.4

migrate_up_latest:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up
migrate_up:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up 1
migrate_down:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ down 1
migrate_force:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ force 20220705105809
migrate_create:
	migrate create -ext sql -dir migrations create_feedback_images_table_v_2_0_5

test-mysql:
	docker-compose up --build -d test-mysql

redis:
	docker-compose up --build -d redis