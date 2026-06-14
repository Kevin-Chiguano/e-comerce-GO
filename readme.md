# Estructura del Sistema

```bash
ecommerce-manager/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── models/
│   │   ├── producto.go
│   │   ├── usuario.go
│   │   ├── carrito.go
│   │   └── pedido.go
│   │
│   ├── handlers/
│   │   ├── producto_handler.go
│   │   ├── usuario_handler.go
│   │   ├── carrito_handler.go
│   │   └── auth_handler.go
│   │
│   ├── services/
│   │   ├── producto_service.go
│   │   ├── usuario_service.go
│   │   ├── carrito_service.go
│   │   └── pedido_service.go
│   │
│   ├── repositories/
│   │   ├── producto_repository.go
│   │   ├── usuario_repository.go
│   │   ├── carrito_repository.go
│   │   └── pedido_repository.go
│   │
│   ├── functional/
│   │   ├── filters.go
│   │   ├── mappers.go
│   │   └── reducers.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logger.go
│   │
│   ├── database/
│   │   └── connection.go
│   │
│   ├── utils/
│   │   ├── json.go
│   │   ├── files.go
│   │   └── validator.go
│   │
│   └── routes/
│       └── routes.go
│
├── pkg/
│   └── response/
│       └── response.go
│
├── configs/
│   └── config.env
│
├── go.mod
└── README.md
```

## Descripción de Carpetas

| Carpeta | Descripción |
|---|---|
| `cmd/` | Contiene el punto de entrada principal de la aplicación (`main.go`). |
| `internal/models/` | Define las estructuras de datos y modelos del sistema. |
| `internal/handlers/` | Maneja las peticiones HTTP y respuestas de la API REST. |
| `internal/services/` | Contiene la lógica de negocio de la aplicación. |
| `internal/repositories/` | Gestiona el acceso y operaciones con la base de datos. |
| `internal/functional/` | Implementa programación funcional como filtros, mapeos y reducciones. |
| `internal/middleware/` | Incluye middlewares para autenticación y registro de logs. |
| `internal/database/` | Configuración y conexión con la base de datos. |
| `internal/utils/` | Funciones auxiliares reutilizables del sistema. |
| `internal/routes/` | Define las rutas y endpoints de la API. |
| `pkg/response/` | Estructuras estandarizadas para respuestas HTTP. |
| `configs/` | Archivos de configuración y variables de entorno. |

## Tecnologías Utilizadas

- Go (Golang)
- API REST
- GORM
- Middleware
- Arquitectura por capas
- Variables de entorno (`.env`)
- Programación funcional
