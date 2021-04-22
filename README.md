# bookstore_users-api
Book store Users API

## Initialize MySQL in docker container
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=pass -v ./db-vol:/var/lib/mysql -p 3306:3306 -d mysql:5.6

## Exporting DB details
export mysql_username="root"
export mysql_password="pass"
export mysql_host="127.0.0.1:3306"
export mysql_schema="users_db"


## Creating requests
curl -X POST localhost:8000/users -d '{"name": "ravishen", "email": "shen@1234"}' -v

curl -X GET localhost:8000/users/0