API curls:
- GET albums: curl http://localhost:8080/albums
- GET album by ID: curl http://localhost:8080/albums/<id>
- POST create a new album: 
    curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "20","title": "Olala coca","artist": "No Name la","price": 99.99}'