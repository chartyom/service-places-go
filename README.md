GOServicePlaces
===

Start
---
`go run main.go server` 

Documentation
---
./Test_places.postman_collection.json

Tests
---
go test -timeout 30s github.com/HenkCord/GOServicePlaces/usecases/places

go test -timeout 30s github.com/HenkCord/GOServicePlaces/cmd

go test -timeout 30s github.com/HenkCord/GOServicePlaces/entities


---
## Database
1. Start MongoDB
2. Create DB **places**
3. Result
* places
  * _id
  * name
  * city
  * rating
  * menu
    * _id
    * name
    * cost
  * updateAt
  * createdAt
