# Answers

Sí, Al acceder a /hello y /world (que no están registrados como rutas), el servidor sigue respondiendo: devuelve "This is the home page" en lugar de un 404.

Esto ocurre porque en el http.ServeMux de Go, el patrón / actúa como un capturador global. Cualquier petición que no coincida con una ruta más específica cae en /. Por eso /hello, /world o cualquier ruta desconocida siempre llega al manejador de la página principal.

Esto es diferente a muchos otros frameworks donde una ruta no registrada devuelve un 404 automáticamente. En el mux por defecto de Go, tienes que ser intencional si quieres devolver un 404 para rutas desconocidas.

  curl http://localhost:8080/
  curl http://localhost:8080/about
  curl http://localhost:8080/blog
  curl http://localhost:8080/hello



