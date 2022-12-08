release:
	docker build -t toyokoyo199/fitness-backend:2.0.21 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:2.0.21

migrate_up_latest:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up
migrate_up:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up 1
migrate_down:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ down 1
migrate_force:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ force 20221117111948
migrate_create:
	migrate create -ext sql -dir migrations add_user_course_assets_table_source_comment_v_2_0_20

test-mysql:
	docker-compose up --build -d test-mysql

redis:
	docker-compose up --build -d redis