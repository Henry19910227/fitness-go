release:
	docker build -t toyokoyo199/fitness-backend:1.8.7 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:1.8.7