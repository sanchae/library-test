# library-test

V GoLang spiši API, ki mora vsebovati naslednje endpointe:
- dodajanje uporabnika (uporabnik naj vsebuje le ime in priimek)
- prikaz uporabnikov
- prikaz knjig ki so na voljo za izposojo s količino (naziv knjige, količina)
- izposoja knjige (kateri uporabnik katero knjigo)
- vračanje knjige (kateri uporabnik katero knjigo)

# run the container
docker compose up --build

# after container starts 
install this as a cli tool: https://github.com/golang-migrate/migrate

run this command in your terminal: make migrate-up

your database should have tables created and populated

api:

get post
http://localhost:8080/users

{
    "first_name": "",
    "last_name": ""
}

get
http://localhost:8080/books

post
http://localhost:8080/books/lend

{
    "book_id": int,
    "user_id": int
}

post
http://localhost:8080/books/return

{
    "book_id": int,
    "user_id": int
}
