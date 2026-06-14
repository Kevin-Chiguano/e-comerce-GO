# Guía de HTMLs para Sistema de E-Commerce con Roles

## 📄 HTMLs Existentes vs Requeridos

### ✅ HTMLs que YA EXISTEN:
- `login.html` - Formulario de login
- `registro.html` - Formulario de registro
- `index.html` - Página de inicio
- `productos.html` - Página de productos (usuario)
- `carrito.html` - Página del carrito de compras
- `pedidos.html` - Página de mis pedidos realizados
- `admin_productos.html` - Panel de listado de productos (admin)
- `admin_nuevo_producto.html` - Formulario para crear nuevo producto (admin)
- `admin_editar_producto.html` - Formulario para editar producto (admin)

### ✅ HTMLs RECOMENDADOS (OPCIONAL - Si quieres mejorar la interfaz):

---

## 📋 Contenido Recomendado para Cada HTML

### 1. **productos.html** - Página de Productos (Usuario)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Productos - E-Commerce</title>
  <link rel="stylesheet" href="/css/styles.css">
</head>
<body>
  <nav><!-- Navegación --></nav>
  
  <div class="productos-container">
    <h1>Nuestros Productos</h1>
    <div id="productos-list"></div>
  </div>

  <script>
    // GET /api/productos - Listar todos
    // Permitir añadir al carrito (POST /api/carrito)
  </script>
</body>
</html>
```

**Funcionalidades:**
- Listar productos desde `GET /api/productos`
- Mostrar nombre, descripción, precio, imagen
- Botón "Añadir al Carrito" → `POST /api/carrito`
- Ver stock disponible

---

### 2. **carrito.html** - Carrito de Compras (Usuario)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Carrito - E-Commerce</title>
</head>
<body>
  <h1>Mi Carrito</h1>
  
  <table id="carrito-items">
    <!-- Mostrar items del carrito -->
  </table>
  
  <div>
    <p>Total: <strong id="total">$0.00</strong></p>
    <button onclick="checkout()">Finalizar Compra</button>
    <button onclick="clearCart()">Vaciar Carrito</button>
  </div>

  <script>
    // GET /api/carrito - Obtener carrito
    // DELETE /api/carrito/{producto_id} - Remover producto
    // POST /api/carrito/clear - Vaciar carrito
    // POST /api/pedidos - Crear pedido (checkout)
  </script>
</body>
</html>
```

**Funcionalidades:**
- `GET /api/carrito` - Mostrar items del carrito
- Mostrar cantidad y precio de cada producto
- Botón para remover items → `DELETE /api/carrito/{id}`
- Botón para vaciar carrito → `POST /api/carrito/clear`
- Botón para finalizar compra → `POST /api/pedidos`
- Mostrar total a pagar

---

### 3. **pedidos.html** - Mis Pedidos (Usuario)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Mis Pedidos - E-Commerce</title>
</head>
<body>
  <h1>Mis Pedidos</h1>
  
  <table id="pedidos-list">
    <!-- Mostrar pedidos realizados -->
  </table>

  <script>
    // GET /api/pedidos - Obtener pedidos del usuario
    // Mostrar estado del pedido, fecha, total
  </script>
</body>
</html>
```

**Funcionalidades:**
- `GET /api/pedidos` - Obtener todos los pedidos del usuario
- Mostrar ID, fecha, estado, total
- Ver detalles del pedido
- Mostrar historial completo de compras

---

### 4. **admin_productos.html** - Panel de Productos (Admin)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Gestión de Productos - Admin</title>
</head>
<body>
  <h1>Productos - Panel de Administración</h1>
  
  <a href="/admin/productos/nuevo" class="btn">Crear Nuevo Producto</a>
  
  <table id="productos-tabla">
    <thead>
      <tr>
        <th>ID</th>
        <th>Nombre</th>
        <th>Precio</th>
        <th>Stock</th>
        <th>Acciones</th>
      </tr>
    </thead>
    <tbody id="productos-body">
      <!-- Mostrar productos -->
    </tbody>
  </table>

  <script>
    // GET /api/productos - Listar todos
    // DELETE /api/productos/{id} - Eliminar
    // PUT /api/productos/{id}/stock - Actualizar stock rápido
  </script>
</body>
</html>
```

**Funcionalidades:**
- `GET /api/productos` - Listar todos los productos
- Link a "Nuevo Producto"
- Link a "Editar" para cada producto
- Botón "Eliminar" → `DELETE /api/productos/{id}`
- Mostrar cantidad actual de stock
- Botón para aumentar/disminuir stock → `PUT /api/productos/{id}/stock`

---

### 5. **admin_nuevo_producto.html** - Crear Producto (Admin)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Nuevo Producto - Admin</title>
</head>
<body>
  <h1>Crear Nuevo Producto</h1>
  
  <form onsubmit="crearProducto(event)">
    <input type="text" id="nombre" placeholder="Nombre" required>
    <textarea id="descripcion" placeholder="Descripción" required></textarea>
    <input type="number" id="precio" placeholder="Precio" step="0.01" required>
    <input type="number" id="stock" placeholder="Stock" required>
    <input type="number" id="categoria_id" placeholder="Categoría ID" required>
    <input type="url" id="imagen" placeholder="URL de imagen">
    <button type="submit">Guardar Producto</button>
  </form>

  <script>
    // POST /api/productos - Crear producto
  </script>
</body>
</html>
```

**Funcionalidades:**
- Formulario para crear nuevo producto
- Campos: nombre, descripción, precio, stock, categoría, imagen
- `POST /api/productos` - Guardar nuevo producto
- Validación de campos
- Redireccionar a admin_productos.html después de guardar

---

### 6. **admin_editar_producto.html** - Editar Producto (Admin)
```html
<!DOCTYPE html>
<html>
<head>
  <title>Editar Producto - Admin</title>
</head>
<body>
  <h1>Editar Producto</h1>
  
  <form onsubmit="actualizarProducto(event)">
    <input type="text" id="nombre" placeholder="Nombre" required>
    <textarea id="descripcion" placeholder="Descripción" required></textarea>
    <input type="number" id="precio" placeholder="Precio" step="0.01" required>
    <input type="number" id="stock" placeholder="Stock" required>
    <input type="number" id="categoria_id" placeholder="Categoría ID" required>
    <input type="url" id="imagen" placeholder="URL de imagen">
    <button type="submit">Actualizar Producto</button>
    <button onclick="aumentarStock(10)">+10 Stock</button>
    <button onclick="disminuirStock(10)">-10 Stock</button>
  </form>

  <script>
    // GET /api/productos/{id} - Obtener datos del producto
    // PUT /api/productos/{id} - Actualizar producto
    // PUT /api/productos/{id}/stock - Actualizar stock
  </script>
</body>
</html>
```

**Funcionalidades:**
- Obtener datos del producto (ID desde parámetro URL)
- `GET /api/productos/{id}` - Cargar datos iniciales
- Formulario editable para todos los campos
- `PUT /api/productos/{id}` - Guardar cambios
- `PUT /api/productos/{id}/stock` - Botones rápidos para actualizar stock
- Validación de campos

---

### 7. **login.html** - Formulario de Login
```html
<!DOCTYPE html>
<html>
<head>
  <title>Login - E-Commerce</title>
</head>
<body>
  <h1>Iniciar Sesión</h1>
  
  <form onsubmit="login(event)">
    <input type="email" id="email" placeholder="Email" required>
    <input type="password" id="password" placeholder="Contraseña" required>
    <button type="submit">Iniciar Sesión</button>
  </form>

  <script>
    // POST /api/login - Obtener token JWT
    // Guardar token en localStorage
    // Redirigir según rol
  </script>
</body>
</html>
```

**Funcionalidades:**
- `POST /api/login` - Autenticar usuario
- Guardar token JWT en localStorage
- Redirigir a `/productos` para usuarios (rol_id=2)
- Redirigir a `/admin/productos` para admin (rol_id=1)

---

### 8. **registro.html** - Formulario de Registro
```html
<!DOCTYPE html>
<html>
<head>
  <title>Registro - E-Commerce</title>
</head>
<body>
  <h1>Crear Cuenta</h1>
  
  <form onsubmit="registro(event)">
    <input type="text" id="nombre" placeholder="Nombre completo" required>
    <input type="email" id="email" placeholder="Email" required>
    <input type="password" id="password" placeholder="Contraseña" required>
    <input type="hidden" id="rol_id" value="2"> <!-- Usuario normal -->
    <button type="submit">Registrarse</button>
  </form>

  <script>
    // POST /api/register - Crear cuenta
    // rol_id = 2 para usuarios normales
  </script>
</body>
</html>
```

**Funcionalidades:**
- `POST /api/register` - Crear nuevo usuario
- Por defecto rol_id=2 (usuario normal)
- Validación de email
- Redirigir a login después de registrarse

---

## 🔑 Variables de Almacenamiento (localStorage)

```javascript
// Guardar después de login
localStorage.setItem('token', 'eyJhbGciOi...');
localStorage.setItem('user_id', '1');
localStorage.setItem('rol_id', '2'); // 1=Admin, 2=User
localStorage.setItem('email', 'user@example.com');

// Obtener token cuando sea necesario
const token = localStorage.getItem('token');
const headers = {
  'Authorization': `Bearer ${token}`,
  'Content-Type': 'application/json'
};
```

---

## 🛠️ Función Helper JavaScript

```javascript
// Función para hacer peticiones con token
async function apiCall(endpoint, method = 'GET', data = null) {
  const token = localStorage.getItem('token');
  const options = {
    method,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    }
  };
  
  if (data) {
    options.body = JSON.stringify(data);
  }
  
  try {
    const response = await fetch(`/api${endpoint}`, options);
    if (!response.ok && response.status === 401) {
      // Token expirado, redirigir a login
      localStorage.clear();
      window.location.href = '/login';
    }
    return await response.json();
  } catch (error) {
    console.error('Error:', error);
    return { error: 'Error en la solicitud' };
  }
}

// Uso:
// const productos = await apiCall('/productos');
// const result = await apiCall('/carrito', 'POST', { producto_id: 1, cantidad: 2 });
```

---

## 📝 Resumen de Cambios

| HTML | Usuario | Admin | Debe Contener |
|------|---------|-------|----------------|
| `login.html` | ✅ | ✅ | Form login, guardador de token |
| `registro.html` | ✅ | ✅ | Form registro, rol_id=2 por defecto |
| `index.html` | ✅ | ✅ | Nav con enlaces, resumen |
| `productos.html` | ✅ | ❌ | Lista productos, btn carrito |
| `carrito.html` | ✅ | ❌ | Items, total, checkout, clear |
| `pedidos.html` | ✅ | ❌ | Historial de pedidos |
| `admin_productos.html` | ❌ | ✅ | Lista con CRUD, botones acciones |
| `admin_nuevo_producto.html` | ❌ | ✅ | Formulario crear producto |
| `admin_editar_producto.html` | ❌ | ✅ | Formulario editar, botones stock |

---

## ✨ Recomendación Final

**¿Necesitas implementar HTMLs?**
- ✅ **SÍ**, si quieres una interfaz de usuario funcional
- ❌ **NO**, si solo necesitas la API REST (ya está lista 100%)

**Los endpoints ya están listos y funcionando**, solo falta conectarlos desde JavaScript en los HTMLs.
