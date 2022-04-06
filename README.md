# Backend - Beers

## Introducción (Problema)
Bender es fanático de las cervezas, y quiere tener un registro de todas las cervezas que prueba y cómo calcular el precio que necesita para comprar una caja de algún tipo especifico de cervezas. Para esto necesita una API REST con esta información que posteriormente compartir con sus amigos. 

## Requisitos
```sh
docker 
docker-compose
```

## Clonar la aplicación
```sh
git clone https://github.com/freischarler/test_ms_beer
```

## Correr la aplicación
Para correr la aplicación es necesario ejecutar el siguiente comando:
```sh
cd test_ms_beer
docker-compose up -d --build 

```
Por ultimo abrir el navegador y poner la siguiente dirección (agregar API endpoints luego):
```sh
http://localhost:9000
```

### API endpoints

| Método | URL                             | Descripción                             |
|--------|---------------------------------|-----------------------------------------|
| GET    | /beers                          | Lista todas las cervezas                |
| POST   | /beers                          | Ingresa una nueva cerveza               |
| GET    | /beers/{beerID}                 | Lista el detalle de la marca de cerveza |
| GET    | /beers/{beerID}/boxprice        | Lista el precio de una caja             |

## Contribuir
Para contribuir realizar un pull request con las sugerencias.
## Licencia
GPL
