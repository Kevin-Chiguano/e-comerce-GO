# E-commerce Go

Este proyecto corresponde a una aplicación web de e-commerce desarrollada en Go utilizando Gin, PostgreSQL y autenticación con JWT. Su propósito es ofrecer una solución básica pero funcional para la gestión de productos, carrito de compras, pedidos y usuarios, con un enfoque claro en arquitectura por capas y separación de responsabilidades.

## Resumen del proyecto

La aplicación permite:

- Registrar usuarios y autenticar sesiones
- Gestionar un catálogo de productos
- Agregar productos a un carrito por usuario
- Crear pedidos a partir del carrito
- Administrar productos y pedidos desde un panel de acceso restringido
- Exponer una interfaz web sencilla para interacción con el sistema

## Tecnologías utilizadas

- Go 1.22
- Gin Gonic
- PostgreSQL
- JWT con golang-jwt/jwt/v5
- bcrypt para el hashing de contraseñas
- godotenv para cargar variables de entorno
- HTML, CSS y JavaScript para la interfaz web

## Estructura del proyecto

```bash
ecommerce-manager/
├── cmd/
│   └── main.go
├── configs/
│   └── .env
├── internal/
│   ├── database/
│   │   └── connection.go
│   ├── functional/
│   │   ├── filters.go
│   │   ├── mappers.go
│   │   └── reducers.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── carrito_handler.go
│   │   ├── pedido_handler.go
│   │   ├── producto_handler.go
│   │   └── usuario_handler.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── role.go
│   ├── models/
│   │   ├── carrito.go
│   │   ├── categoria.go
│   │   ├── pedido.go
│   │   ├── producto.go
│   │   └── usuario.go
│   ├── repositories/
│   │   ├── carrito_repository.go
│   │   ├── pedido_repository.go
│   │   ├── producto_repository.go
│   │   └── usuario_repository.go
│   ├── routes/
│   │   └── routes.go
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── carrito_service.go
│   │   ├── pedido_service.go
│   │   ├── producto_service.go
│   │   └── usuario_service.go
│   └── utils/
│       └── validator.go
├── web/
│   ├── css/
│   │   └── styles.css
│   ├── js/
│   │   └── app.js
│   └── templates/
│       ├── admin_editar_producto.html
│       ├── admin_nuevo_producto.html
│       ├── admin_pedidos.html
│       ├── admin_productos.html
│       ├── carrito.html
│       ├── index.html
│       ├── login.html
│       ├── pedidos.html
│       ├── productos.html
│       └── registro.html
├── go.mod
└── readme.md
```

## Arquitectura y organización del código

### Punto de entrada

El archivo principal de ejecución se encuentra en `cmd/main.go`. Allí se realiza lo siguiente:

- Carga de variables de entorno desde `configs/.env`
- Conexión con PostgreSQL
- Configuración del servidor Gin
- Registro de rutas y middlewares
- Servicio de archivos estáticos desde la carpeta `web`

### Capa de modelos

Los modelos definidos en `internal/models` representan las entidades del dominio:

- `Usuario`: datos del usuario, correo electrónico, rol y fecha de registro
- `Producto`: información del producto, incluidas descripción, precio, stock e imagen
- `Carrito` y `CarritoDetalle`: gestión del estado temporal del carrito
- `Pedido` y `DetallePedido`: representación de compras realizadas

### Capa de handlers

Los handlers reciben las solicitudes HTTP y delegan el procesamiento a los servicios correspondientes. Entre ellos se incluyen:

- `AuthHandler`: registro e inicio de sesión
- `ProductoHandler`: listado, creación, actualización, eliminación y ajuste de stock
- `CarritoHandler`: gestión de productos en el carrito
- `PedidoHandler`: creación y consulta de pedidos
- `UsuarioHandler`: registro de usuarios desde la API

### Capa de servicios

La lógica de negocio se encuentra en `internal/services` y es responsable de:

- Validar reglas de negocio
- Coordinar operaciones entre repositorios
- Generar tokens JWT
- Encriptar contraseñas
- Convertir un carrito en un pedido

### Capa de repositorios

Los repositorios en `internal/repositories` interactúan directamente con PostgreSQL y encapsulan las operaciones de base de datos como:

- Inserción, consulta, actualización y eliminación de registros
- Lectura de productos, usuarios, carritos y pedidos
- Operaciones transaccionales vinculadas al carrito y al pedido

### Middleware

El proyecto implementa middlewares para seguridad y control de acceso:

- `AuthMiddleware`: valida JWT en el encabezado `Authorization`
- `RoleMiddleware`: restringe rutas según el rol del usuario

Roles definidos:

- Administrador: `1`
- Usuario: `2`

## Flujo funcional

### 1. Registro y autenticación

1. El usuario registra sus datos.
2. La contraseña se almacena de forma segura mediante bcrypt.
3. Al iniciar sesión se valida la credencial y se emite un token JWT.
4. El token se reutiliza para autenticar peticiones protegidas.

### 2. Catálogo de productos

- Los productos pueden consultarse públicamente.
- El administrador puede crear, modificar y eliminar productos.

### 3. Carrito de compras

- Cada usuario autenticado cuenta con un carrito asociado.
- Se pueden agregar, remover o vaciar productos.

### 4. Proceso de compra

- Al confirmar la compra, el sistema convierte el contenido del carrito en un pedido.
- Se calcula el total del pedido.
- El carrito se limpia automáticamente tras la creación del pedido.

## Rutas principales

### Rutas públicas

- `GET /`
- `GET /login`
- `GET /registro`
- `GET /productos`
- `POST /api/register`
- `POST /api/login`
- `GET /api/productos`
- `GET /api/productos/:id`

### Rutas protegidas por autenticación

- `GET /api/carrito`
- `POST /api/carrito`
- `POST /api/carrito/clear`
- `DELETE /api/carrito/:producto_id`
- `POST /api/pedidos`
- `GET /api/pedidos`

### Rutas protegidas por rol de administrador

- `POST /api/productos`
- `PUT /api/productos/:id`
- `DELETE /api/productos/:id`
- `PUT /api/productos/:id/stock`
- `GET /api/admin/pedidos`
- `PUT /api/admin/pedidos/:id/aprobar`

## Requisitos previos

- Go 1.22 o superior
- PostgreSQL instalado y en ejecución
- Un usuario y una base de datos creados para la aplicación

## Variables de entorno

El proyecto carga la configuración desde `configs/.env` o desde el entorno actual. Se esperan las siguientes variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=ecommerce
JWT_SECRET=tu_secreto_jwt
PORT=8080
```

## Instalación y ejecución

1. Clona el repositorio.
2. Crea el archivo `configs/.env` con las variables anteriores.
3. Asegúrate de tener PostgreSQL disponible.
4. Instala las dependencias:

```bash
go mod tidy
```

5. Ejecuta la aplicación:

```bash
go run ./cmd/main.go
```

6. Abre el navegador en:

```text
http://localhost:8080
```

## Esquema de base de datos sugerido

La aplicación espera tablas como las siguientes:

- `usuarios`
- `productos`
- `carrito`
- `carrito_detalle`
- `pedidos`
- `detalle_pedidos`

Ejemplo básico:

```sql
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    rol_id INT NOT NULL DEFAULT 2,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE productos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL,
    descripcion TEXT,
    precio NUMERIC(10,2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    categoria_id INT,
    imagen TEXT
);

CREATE TABLE carrito (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL
);

CREATE TABLE carrito_detalle (
    id SERIAL PRIMARY KEY,
    carrito_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL DEFAULT 1
);

CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL,
    total NUMERIC(10,2) NOT NULL,
    estado VARCHAR(50) NOT NULL DEFAULT 'PENDIENTE',
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE detalle_pedidos (
    id SERIAL PRIMARY KEY,
    pedido_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,
    precio_unitario NUMERIC(10,2) NOT NULL,
    subtotal NUMERIC(10,2) NOT NULL
);
```

## Observaciones

- El proyecto se encuentra en una etapa inicial y algunas vistas web son relativamente simples.
- El flujo principal de compra está implementado en el backend.
- La seguridad depende del correcto uso de `JWT_SECRET` y de una configuración adecuada de la base de datos.

## Próximos pasos recomendados

- Mejorar la experiencia visual del panel administrativo
- Implementar una vista completa de historial de pedidos
- Agregar validaciones de negocio más robustas
- Incorporar pruebas unitarias e integración
- Optimizar la gestión de stock en el proceso de compra

