# Wybieramy obraz GoLang jako bazowy
FROM golang:1.21-alpine

# Ustawiamy katalog roboczy
WORKDIR /app

# Kopiujemy pliki projektu do kontenera
COPY . .

# Instalujemy zależności
RUN go mod tidy

# Kompilujemy aplikację
RUN go build -o main .

# Uruchamiamy aplikację
CMD ["./main"]
