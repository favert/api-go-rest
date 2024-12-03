{
	"info": {
		"_postman_id": "ea845874-b275-44bb-9a4a-665901de63ff",
		"name": "go rest API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "10917455"
	},
	"item": [
		{
			"name": "Todas Personalidades",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8000/personalidades",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"personalidades"
					]
				}
			},
			"response": []
		},
		{
			"name": "Cria Personalidade",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"id\" : 3,\r\n    \"nome\" : \"Nome Postman\",\r\n    \"historia\" : \"Historia Postman\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8000/personalidades",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"personalidades"
					]
				}
			},
			"response": []
		},
		{
			"name": "Retorna Uma Personalidade",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8000/personalidades/2",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"personalidades",
						"2"
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete Uma Personalidade",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8000/personalidades/2",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"personalidades",
						"2"
					]
				}
			},
			"response": []
		},
		{
			"name": "Edita Personalidade",
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"id\" : 3,\r\n    \"nome\" : \"Nome Postman Nova\",\r\n    \"historia\" : \"Historia Postman Nova\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8000/personalidades/2",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"personalidades",
						"2"
					]
				}
			},
			"response": []
		}
	]
}