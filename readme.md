Estructura de del Sistema

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
