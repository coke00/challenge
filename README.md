# intalation
export GOPATH=$HOME/go
mkdir -p $GOPATH/src/github.com/coke00/challenge

# Build and Run

- go build

`` Creara un archivo ejecutable .exe``

- ./name.exe

`` Ejecuta archivo y despliega servidor en el puerto configurado``

- go run main.go

`` Ejecuta el archivo elegido y su entorno``

- go test -v

`` Ejecuta unitTest``

# Routes

- http://localhost:4203

``Welcome``

- http://localhost:4203/contagiados

``Obtener estadisticas globales y paises ``


- http://localhost:4203/contagiados/{iso}

`` Obtener las estadisticas de los contagiados por pais, ingresando el codigo del pais o iso, ejemplo: CL
request de referencia http://localhost:4203/contagiados/CL ``

# Use .env files

go get github.com/joho/godotenv

how use it
godotenv.Load()

// or

godotenv.Load(".env")
