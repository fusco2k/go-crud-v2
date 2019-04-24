##go-crud-v2

The basic purpose of this project is to solidify the usage of the mongodb driver on a primitive crud webapp

#BASIC USAGE

Use this exported postman collection to generate the requests 

```
{
	"info": {
		"_postman_id": "4e24398f-3fc6-4712-92a2-63fb40c29822",
		"name": "go-crud-v2",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "http://localhost:8080/user",
			"request": {
				"method": "PUT",
				"header": [
					{
						"key": "Content-Type",
						"name": "Content-Type",
						"value": "application/json",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "[{\n\t\"name\":\"Tony\"\n},\n{\n\t\"name\":\"Frodo\"\n}]"
				},
				"url": {
					"raw": "http://localhost:8080/user",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"user"
					]
				}
			},
			"response": []
		},
		{
			"name": "http://localhost:8080/user",
			"request": {
				"method": "DELETE",
				"header": [
					{
						"key": "Content-Type",
						"name": "Content-Type",
						"value": "application/json",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"name\": \"Frodo\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/user",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"user"
					]
				}
			},
			"response": []
		},
		{
			"name": "http://localhost:8080/user",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"name": "Content-Type",
						"value": "application/json",
						"type": "text"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n\t\"name\" : \"Frodo\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/user",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"user"
					]
				}
			},
			"response": []
		},
		{
			"name": "http://localhost:8080/user",
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": ""
				},
				"url": {
					"raw": "http://localhost:8080/user",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"user"
					]
				}
			},
			"response": []
		}
	]
}
```