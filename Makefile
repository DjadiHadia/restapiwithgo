createdb:
docker exec -it postres12 creatdb --username=root --owner=root RentalCar
sqlc:
sqlc generate

.PHONY : sqlc