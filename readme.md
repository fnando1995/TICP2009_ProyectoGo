# Proyecto GO

## Consumo de API REST en GO

Escribirán un pequeño servicio web que se pueda consultar a través de un REST API, a su vez este servicio web tiene que consumir al menos 2 otros REST APIs de cualquier fuente que le interese ustedes.

Su servicio/API debe de permitir un mínimo de una consulta que resulte de mezclar la información de o acciones de los 2 (o más) APIs que estén utilizando.



## Ambiente 

Las pruebas realizadas en este proyecto fueron hechas bajo un sistema con las siguientes características:


- S.O. -> Linux Ubuntu 20.04
- Go -> 1.16.6

## Descripcion de APIs usados

El head de API es https://api.coingecko.com/api/v3/, donde se puede notar que la API de coingecko esta en su versión número 3.


### coins/list

Esta consulta retorna una JSON con todos las las monedas incluídas en coingecko. Al ser bastante pesada por la magnitud de los datos suele tardarse, por lo que se considera grabar la información en una base local y utilizarla cada día para actualizar la información de ser necesario.

Esta API no entrega información relevante más allá de nombre, símbolo y id de la moneda.

### 