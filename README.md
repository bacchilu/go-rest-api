# Go Rest API

A test project to learn the language

## test HTTP POST

    curl -X POST -H "Content-Type: application/json" -d '{"name": "Luca", "description": "me"}' http://0.0.0.0:8080/events | jq