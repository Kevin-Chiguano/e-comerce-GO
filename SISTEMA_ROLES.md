# Sistema de Roles y Control de Acceso - E-Commerce

## 🎯 Resumen de Cambios Implementados

### 1. **Middleware de Roles** 
Archivo: `internal/middleware/role.go`
- Nuevo middleware `RoleMiddleware` que valida permisos por rol
- Constantes: `RoleAdmin = 1` y `RoleUser = 2`
- Retorna 403 Forbidden si el usuario no tiene permisos

### 2. **CRUD de Productos (Admin)** 
Endpoints disponibles (requieren `rol_id = 1`):

#### Crear Producto
```http
POST /api/productos
Authorization: Bearer {token}
Content-Type: application/json

{
  "nombre": "Producto",
  "descripcion": "Descripción",
  "precio": 99.99,
  "stock": 50,
  "categoria_id": 1,
  "imagen": "url_imagen.jpg"
}
```

#### Actualizar Producto
```http
PUT /api/productos/{id}
Authorization: Bearer {token}
Content-Type: application/json

{
  "nombre": "Nuevo nombre",
  "descripcion": "Nueva descripción",
  "precio": 79.99,
  "stock": 100,
  "categoria_id": 1,
  "imagen": "url_imagen.jpg"
}
```

#### Eliminar Producto
```http
DELETE /api/productos/{id}
Authorization: Bearer {token}
```

#### Actualizar Stock
```http
PUT /api/productos/{id}/stock
Authorization: Bearer {token}
Content-Type: application/json

{
  "cantidad": 10  # Puede ser positivo o negativo
}
```

#### Obtener Todos los Productos (Público)
```http
GET /api/productos
```

#### Obtener Producto por ID (Público)
```http
GET /api/productos/{id}
```

---

### 3. **Gestión de Carrito (Usuario)**
Endpoints disponibles (requieren autenticación):

#### Obtener Carrito
```http
GET /api/carrito
Authorization: Bearer {token}
```

#### Añadir Producto al Carrito
```http
POST /api/carrito
Authorization: Bearer {token}
Content-Type: application/json

{
  "producto_id": 1,
  "cantidad": 2
}
```

#### Remover Producto del Carrito
```http
DELETE /api/carrito/{producto_id}
Authorization: Bearer {token}
```

#### Vaciar Carrito
```http
POST /api/carrito/clear
Authorization: Bearer {token}
```

---

### 4. **Compra de Productos (Simular Pedido)**
Endpoint disponible (requiere autenticación):

#### Crear Pedido desde Carrito
```http
POST /api/pedidos
Authorization: Bearer {token}
```

**Respuesta exitosa:**
```json
{
  "message": "Pedido creado exitosamente",
  "pedido": {
    "id": 1,
    "usuario_id": 1,
    "total": 199.98,
    "estado": "PENDIENTE",
    "fecha": "2026-06-14T01:15:09Z",
    "detalles": [
      {
        "producto_id": 1,
        "cantidad": 2,
        "precio_unitario": 99.99,
        "subtotal": 199.98
      }
    ]
  }
}
```

✅ **Automático**: El carrito se vacía después de crear el pedido

---

### 5. **Ver Mis Pedidos Realizados**
Endpoint disponible (requiere autenticación):

```http
GET /api/pedidos
Authorization: Bearer {token}
```

**Respuesta:**
```json
{
  "message": "Lista de pedidos",
  "pedidos": [
    {
      "id": 1,
      "usuario_id": 1,
      "total": 199.98,
      "estado": "PENDIENTE",
      "fecha": "2026-06-14T01:15:09Z"
    }
  ]
}
```

---

### 6. **Cambios en Handlers**

#### ProductoHandler
- ✅ `Create()` - Crear nuevo producto
- ✅ `Update()` - Actualizar producto
- ✅ `Delete()` - Eliminar producto
- ✅ `UpdateStock()` - Actualizar stock
- ✅ `GetAll()` - Obtener todos (ya existía)
- ✅ `GetByID()` - Obtener por ID (ya existía)

#### CarritoHandler
- ✅ `RemoveItem()` - Remover producto del carrito
- ✅ `ClearCarrito()` - Vaciar carrito
- ✅ `AddItem()` - Añadir producto (ya existía)
- ✅ `GetCarrito()` - Obtener carrito (ya existía)

#### PedidoHandler
- ✅ `CrearPedido()` - Crear pedido desde carrito
- ✅ `GetMisPedidos()` - Obtener pedidos del usuario

---

### 7. **Cambios en Repositories**

#### ProductoRepository
- ✅ `Create()` - Crear producto en BD
- ✅ `Update()` - Actualizar producto
- ✅ `Delete()` - Eliminar producto
- ✅ `UpdateStock()` - Actualizar stock

#### CarritoRepository
- ✅ `RemoveItem()` - Remover item del carrito
- ✅ `Clear()` - Limpiar todo el carrito

#### PedidoRepository
- ✅ `GetByUsuario()` - Traer pedidos del usuario (ya existía)

---

### 8. **Cambios en Services**

#### ProductoService
- ✅ `CreateProducto()`
- ✅ `UpdateProducto()`
- ✅ `DeleteProducto()`
- ✅ `UpdateStock()`

#### CarritoService
- ✅ `RemoveFromCarrito()`
- ✅ `ClearCarrito()`

#### PedidoService
- ✅ `GetMisPedidos()` - Obtener pedidos del usuario
- ✅ Automáticamente limpia el carrito después de crear un pedido

---

### 9. **Rutas Organizadas por Rol**

```
/api
├── /register (Público)
├── /login (Público)
├── /productos (Público)
│   ├── GET / - Listar todos
│   └── GET /:id - Obtener por ID
│
├── Protegidas (con AuthMiddleware)
│   ├── /carrito (Usuario)
│   │   ├── GET / - Obtener carrito
│   │   ├── POST / - Añadir producto
│   │   ├── DELETE /:producto_id - Remover producto
│   │   └── POST /clear - Vaciar carrito
│   │
│   ├── /pedidos (Usuario)
│   │   ├── POST / - Crear pedido (Checkout)
│   │   └── GET / - Ver mis pedidos
│   │
│   └── /productos (Admin Solo - RoleMiddleware)
│       ├── POST / - Crear producto
│       ├── PUT /:id - Actualizar producto
│       ├── DELETE /:id - Eliminar producto
│       └── PUT /:id/stock - Actualizar stock
```

---

## 📋 Archivos Modificados/Creados

- ✅ `internal/middleware/role.go` (NUEVO)
- ✅ `internal/handlers/producto_handler.go` (ACTUALIZADO)
- ✅ `internal/handlers/carrito_handler.go` (ACTUALIZADO)
- ✅ `internal/handlers/pedido_handler.go` (ACTUALIZADO)
- ✅ `internal/services/producto_service.go` (ACTUALIZADO)
- ✅ `internal/services/carrito_service.go` (ACTUALIZADO)
- ✅ `internal/services/pedido_service.go` (ACTUALIZADO)
- ✅ `internal/repositories/producto_repository.go` (ACTUALIZADO)
- ✅ `internal/repositories/carrito_repository.go` (ACTUALIZADO)
- ✅ `internal/routes/routes.go` (ACTUALIZADO)
- ✅ `internal/functional/filters.go` (COMPLETADO)
- ✅ `internal/functional/mappers.go` (COMPLETADO)
- ✅ `internal/functional/reducers.go` (COMPLETADO)
- ✅ `internal/utils/validator.go` (COMPLETADO)

---

## 🚀 Cómo Usar

### Para Administrador (rol_id = 1):
1. Hacer login y obtener token
2. Crear productos: `POST /api/productos`
3. Actualizar stock: `PUT /api/productos/{id}/stock`
4. Actualizar/Eliminar productos: `PUT/DELETE /api/productos/{id}`

### Para Usuario (rol_id = 2):
1. Hacer login y obtener token
2. Ver productos: `GET /api/productos`
3. Añadir al carrito: `POST /api/carrito`
4. Ver carrito: `GET /api/carrito`
5. Simular compra: `POST /api/pedidos` (checkout)
6. Ver mis pedidos: `GET /api/pedidos`

---

## 📝 Notas Importantes

- El token JWT debe incluir `rol_id` en los claims
- El carrito se limpia automáticamente después de crear un pedido
- El stock puede aumentar o disminuir con valores positivos o negativos
- Todos los endpoints protegidos requieren header `Authorization: Bearer {token}`
