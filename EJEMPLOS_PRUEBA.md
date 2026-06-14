# Ejemplos de Prueba - Sistema de E-Commerce con Roles

## 1. Registro de Usuario

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Admin User",
    "email": "admin@example.com",
    "password": "password123",
    "rol_id": 1
  }'
```

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Normal User",
    "email": "user@example.com",
    "password": "password123",
    "rol_id": 2
  }'
```

---

## 2. Login y Obtener Token

### Admin Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

### User Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**Respuesta esperada:**
```json
{
  "message": "Login exitoso",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## 3. Operaciones ADMIN - CRUD de Productos

### 3.1 Crear Producto (Admin)
```bash
ADMIN_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

curl -X POST http://localhost:8080/api/productos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "nombre": "Laptop Dell XPS 13",
    "descripcion": "Laptop ultraportátil",
    "precio": 1299.99,
    "stock": 50,
    "categoria_id": 1,
    "imagen": "https://example.com/laptop.jpg"
  }'
```

### 3.2 Listar Todos los Productos (Público)
```bash
curl -X GET http://localhost:8080/api/productos
```

### 3.3 Obtener Producto por ID (Público)
```bash
curl -X GET http://localhost:8080/api/productos/1
```

### 3.4 Actualizar Producto (Admin)
```bash
curl -X PUT http://localhost:8080/api/productos/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "nombre": "Laptop Dell XPS 13 - Actualizado",
    "descripcion": "Laptop ultraportátil mejorada",
    "precio": 1199.99,
    "stock": 45,
    "categoria_id": 1,
    "imagen": "https://example.com/laptop-new.jpg"
  }'
```

### 3.5 Actualizar Stock (Admin - Incrementar)
```bash
curl -X PUT http://localhost:8080/api/productos/1/stock \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "cantidad": 20
  }'
```

### 3.6 Actualizar Stock (Admin - Disminuir)
```bash
curl -X PUT http://localhost:8080/api/productos/1/stock \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d '{
    "cantidad": -10
  }'
```

### 3.7 Eliminar Producto (Admin)
```bash
curl -X DELETE http://localhost:8080/api/productos/1 \
  -H "Authorization: Bearer $ADMIN_TOKEN"
```

---

## 4. Operaciones USER - Carrito y Compras

### 4.1 Obtener Carrito Actual
```bash
USER_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

curl -X GET http://localhost:8080/api/carrito \
  -H "Authorization: Bearer $USER_TOKEN"
```

### 4.2 Añadir Producto al Carrito
```bash
curl -X POST http://localhost:8080/api/carrito \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $USER_TOKEN" \
  -d '{
    "producto_id": 1,
    "cantidad": 2
  }'
```

### 4.3 Remover Producto del Carrito
```bash
curl -X DELETE http://localhost:8080/api/carrito/1 \
  -H "Authorization: Bearer $USER_TOKEN"
```

### 4.4 Vaciar Carrito Completamente
```bash
curl -X POST http://localhost:8080/api/carrito/clear \
  -H "Authorization: Bearer $USER_TOKEN"
```

### 4.5 Simular Compra (Crear Pedido - Checkout)
```bash
curl -X POST http://localhost:8080/api/pedidos \
  -H "Authorization: Bearer $USER_TOKEN"
```

**Respuesta esperada:**
```json
{
  "message": "Pedido creado exitosamente",
  "pedido": {
    "id": 1,
    "usuario_id": 2,
    "total": 2599.98,
    "estado": "PENDIENTE",
    "fecha": "2026-06-14T01:45:57Z",
    "detalles": [
      {
        "pedido_id": 1,
        "producto_id": 1,
        "cantidad": 2,
        "precio_unitario": 1299.99,
        "subtotal": 2599.98
      }
    ]
  }
}
```

### 4.6 Ver Mis Pedidos Realizados
```bash
curl -X GET http://localhost:8080/api/pedidos \
  -H "Authorization: Bearer $USER_TOKEN"
```

**Respuesta esperada:**
```json
{
  "message": "Lista de pedidos",
  "pedidos": [
    {
      "id": 1,
      "usuario_id": 2,
      "total": 2599.98,
      "estado": "PENDIENTE",
      "fecha": "2026-06-14T01:45:57Z"
    }
  ]
}
```

---

## 5. Códigos de Error Esperados

### Error 400 - Bad Request
```json
{
  "error": "Formato de datos inválido"
}
```

### Error 401 - Unauthorized
```json
{
  "error": "Token requerido"
}
```

### Error 403 - Forbidden (Sin permiso de rol)
```json
{
  "error": "No tienes permiso para acceder a este recurso"
}
```

### Error 404 - Not Found
```json
{
  "error": "Producto no encontrado"
}
```

### Error 500 - Server Error
```json
{
  "error": "Descripción del error"
}
```

---

## 6. Flujo Completo de Usuario

1. **Registrarse** → `POST /api/register`
2. **Iniciar sesión** → `POST /api/login` (obtener token)
3. **Ver productos** → `GET /api/productos` (sin token)
4. **Añadir al carrito** → `POST /api/carrito` (con token)
5. **Ver carrito** → `GET /api/carrito` (con token)
6. **Finalizar compra** → `POST /api/pedidos` (con token)
7. **Ver mis pedidos** → `GET /api/pedidos` (con token)

---

## 7. Flujo Completo de Admin

1. **Registrarse como Admin** → `POST /api/register` (rol_id: 1)
2. **Iniciar sesión** → `POST /api/login` (obtener token)
3. **Crear producto** → `POST /api/productos` (con token admin)
4. **Actualizar stock** → `PUT /api/productos/{id}/stock` (con token admin)
5. **Ver todos los productos** → `GET /api/productos`
6. **Actualizar información de producto** → `PUT /api/productos/{id}` (con token admin)
7. **Eliminar producto** → `DELETE /api/productos/{id}` (con token admin)

---

## ⚠️ Notas Importantes

- **rol_id = 1**: Administrador (acceso a CRUD completo de productos)
- **rol_id = 2**: Usuario normal (solo puede comprar)
- El **carrito se limpia automáticamente** después de crear un pedido
- El **stock puede aumentar o disminuir** con valores positivos o negativos
- El token debe estar en formato `Bearer {token}` en el header
