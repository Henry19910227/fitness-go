debug:
	docker build -t toyokoyo199/fitness-backend:1.8.6 --build-arg mode=develop .
	docker push toyokoyo199/fitness-backend:1.8.6
release:
	docker build -t toyokoyo199/fitness-backend:1.8.6 --build-arg mode=release .
	docker push toyokoyo199/fitness-backend:1.8.6