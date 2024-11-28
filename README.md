# pruebaCrudGo

Para realizar esta prueba de crud inicalmente se diseño un modelo relacion-entidad de la base de datos

Poseteriormente se realizo la cracion de una base de datos en mi local llamada school y un script para crear las tablas anteriormente diseñadas en el modelo.

El modelo y el script de creacion estan anexados en la carpeta Docs

se opto por una arquitectura clean, agrupando los repositorios, servicios y handlers

para el repositorio de estudiantes se realizaron las correspondientes operaciones crud: obtener todos, obtener por id, crear, eliminar y actualizar.

el framework utilizado para exponer el api es Gin Gonic, utilizando este se crearon las rutas y se lee el contexto de cada operacion a realizar.

para utilizar el programa primero se debe logear accediendo al endpoint /login
