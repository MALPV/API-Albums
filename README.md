# API-Albums

_API Fake sobre albums creada para el aprendizaje del lenguaje Go usando el framework Gin_

## Comenzando 🚀

_Para probar y levantar la API deberá tener instalado Go en su sistema._

_Para verificar si go esta instalado ingresa en la terminal_.

_Ejecuta el comando_

```
go version
```

### Para iniciar un proyecto 🔧

_Desde la ubicación del archivo main.go iniciar la terminal_

_Ejecutar el comando_

```
go run main.go
```

### Para crear el ejecutable 🔧

```
go build main.go
```

### Probando las rutas ⚙️

_Ruta base_

```
http://localhost:8080/
```

_Para obtener la lista de los albums [GET]_

```
http://localhost:8080/albums
```
_Respuesta ejemplo_

```json
[
    {
        "id": "1",
        "title": "Reir o llorar",
        "artist": "Jacqueline Valenzuela",
        "price": 10
    },
    {
        "id": "2",
        "title": "Arcoiris",
        "artist": "Marcos Palma",
        "price": 17.99
    }
]
```

_Para crear un nuevo album [POST]_

```
http://localhost:8080/albums
```
_Desde postman_

```json
{
    "id":"10",
    "title":"Sobredosis",
    "artist":"Mora",
    "price": 36.99
}
```

## Construido con 🛠️

* [Lenguaje Go](https://go.dev/) - El lenguaje para crear software rápido, confiable y eficiente a escala

* [Framework Gin](https://go.dev/) - Si necesitas rendimiento y buena productividad, Gin te encantará.

## Autores ✒️

_Quien les escribe_

* **Marcos Palma Valenzuela** - *Trabajo Inicial* 

También puedes mirar la lista de todos [mis proyectos](https://github.com/MALPV)