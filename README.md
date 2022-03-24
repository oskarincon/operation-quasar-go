# 📡 operation-quasar-go 📡

![enter image description here](https://static.wikia.nocookie.net/esstarwars/images/5/59/QuasarFireCarrier-SWR.png)


Operación Fuego de Quasar

## Requirementos 🛠️ 
**GO 1.7**

**Docker**

**Postman**

## Clonar Proyecto ⏬
> $ git clone git@github.com:oskarincon/operation-quasar-go.git

## Ejecucion de proyecto 🚀
Ejecutar en con docker-compose:
> $ docker-compose -f ./docker-compose.yml up --build
Ejecutar en local con go:
> $ go run .

## Endpoints
    ENDPOINT AWS : https://w13tb32mml.execute-api.us-east-1.amazonaws.com/v1
    ENDPOINT DOCKER : http://127.0.0.1:3000

**topsecret**
Obtenga la posición y el mensaje completo publicando tres satélites con la distancia y el mensaje completo.

`POST -> /alliance/topsecret`

EJEMPLO request body:

    {
    "satellites": [
         {
            "name": "kenobi",
            "distance": 10.0,
            "message": ["este", "", "", "mensaje", ""]
         },
              {
            "name": "skywalker",
            "distance": 1.0,
            "message": ["", "es", "", "", ""]
         },
              {
            "name": "sato",
            "distance": 10.0,
            "message": ["este", "", "un", "", ""]
         }
      ]
    }

RESPUESTA 200:

    {
        "position": {
            "x": -487.314375,
            "y": 1574.38125
        },
        "message": "este es un mensaje "
    }

RESPUESTA 400:

    {
        "data": "json: invalid character } as slice",
        "success": false
    }

RESPUESTA 404:

    {
        "data": "no se tienen distancias para todos los satelites",
        "success": false
    }
<br>
<br>

**opsecret_split**
Posts de distancia y un mensaje de un satellite (kenobi, skywalker, sato) incompleto que se guardará en la memoria del sistema. Recibirá el nombre del satélite como path param.

`POST -> /alliance/topsecret_split/{satellite_name}`

EJEMPLO request body:

    {
            "distance": 100.0,
            "message": ["este", "", "un", "", ""]
    }

EJEMPLO 200:

    {
    "message": [
        "este",
        "",
        "un",
        "",
        ""
    ],
    "name": "sato",
    "distance": 100
   }

RESPUESTA 400:

    {
        "data": "json: invalid character } as slice",
        "success": false
    }

RESPUESTA 404:

    {
        "data": "el satelite a guardar debe ser kenobi, skywalker o sato",
        "success": false
    }

`GET -> /alliance/topsecret_split/`

GET el cual toma los datos en memoria guardados en el servicio /topsecret_split/{satellite_name} si se encuentran los datos necesarios, se decifra el mensaje y se genera la posición .

Example 200 response:

    {
        "position": {
            "x": -487.5,
            "y": 1575
        },
        "message": "este un "
    }

RESPUESTA 404:

    {
        "data": "no hay data suficiente para calcular datos",
        "success": false
    }

**Desarrollo**
- Por medio del metodo determinantes 3x2 se resuelve la ubicación de los satellites.
- Por medio de validación de desfase de mensajes (igualar longitud de mensajes) y validacion de cada mensaje se resuelve el mensaje.
- La api es realizada por medio del framework fiber Express para golang, se genera el microservicio
- Se guarda el cache en memoria con ReneKroon/ttlcache/v2 
- se realiza la implementación del api por medio de AWS
- En la carpeta collectionsPostman se encuentra una colección que se puede importar en postman para el consumo de los servicios desplegados por docker o aws