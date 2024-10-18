# library-test
docker compose up --build

your database should have tables created and populated

- prikaz uporabnikov
```
curl http://localhost:8080/users
```

- dodajanje uporabnika (uporabnik naj vsebuje le ime in priimek)
```
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe"
  }'
```

- prikaz knjig ki so na voljo za izposojo s količino (naziv knjige, količina)
```
curl http://localhost:8080/books
```

- izposoja knjige (kateri uporabnik katero knjigo)
```
curl -X POST http://localhost:8080/books/lend \
  -H "Content-Type: application/json" \
  -d '{
    "book_id": 1,
    "user_id": 2
  }'
```

- vračanje knjige (kateri uporabnik katero knjigo)
```
curl -X POST http://localhost:8080/books/return \
  -H "Content-Type: application/json" \
  -d '{
    "book_id": 1,
    "user_id": 2
  }'
```