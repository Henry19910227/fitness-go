release:
	docker build -t toyokoyo199/fitness-backend:1.9.7 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:1.9.7

migrate_up_latest:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up
migrate_up:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ up 1
migrate_down:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ down 1
migrate_force:
	migrate -database mysql://henry:aaaa8027@tcp\(localhost:8889\)/fitness -path migrations/ force 20220527181622
migrate_create:
	migrate create -ext sql -dir migrations create_diets_table_v_1_9_8