// Cargar productos
async function cargarProductos() {
    const res = await fetch('/api/productos');
    const productos = await res.json();
    const container = document.getElementById('productos-list');

    if (!container) return;

    container.innerHTML = '';
    productos.forEach(p => {
        container.innerHTML += `
            <div class="producto">
                <h3>${p.nombre}</h3>
                <p>${p.descripcion || ''}</p>
                <p><strong>$${p.precio}</strong></p>
                <button onclick="agregarAlCarrito(${p.id})">Añadir al Carrito</button>
            </div>
        `;
    });
}

// Agregar al carrito
async function agregarAlCarrito(productoId) {
    await fetch('/api/carrito', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ producto_id: productoId, cantidad: 1 })
    });
    alert('Producto añadido al carrito!');
}

// Cargar carrito
async function cargarCarrito() {
    const res = await fetch('/api/carrito');
    const carrito = await res.json();
    const container = document.getElementById('carrito-list');
    let total = 0;

    if (!container) return;

    container.innerHTML = '';
    carrito.detalles.forEach(item => {
        const subtotal = item.cantidad * item.producto.precio;
        total += subtotal;
        container.innerHTML += `
            <div class="producto">
                <h4>${item.producto.nombre}</h4>
                <p>Cantidad: ${item.cantidad} × $${item.producto.precio}</p>
                <p>Subtotal: $${subtotal}</p>
            </div>
        `;
    });

    document.getElementById('total').innerHTML = `<h2>Total: $${total}</h2>`;
}

// Crear Pedido
async function crearPedido() {
    const res = await fetch('/api/pedidos', { method: 'POST' });
    if (res.ok) {
        alert('¡Pedido realizado con éxito!');
        window.location.href = '/pedidos';
    } else {
        alert('Error al crear el pedido');
    }
}

// Cargar pedidos
async function cargarPedidos() {
    // Por ahora solo placeholder
    document.getElementById('pedidos-list').innerHTML = '<p>Funcionalidad de pedidos en desarrollo...</p>';
}

// Login (básico)
document.getElementById('loginForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    const res = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });

    if (res.ok) {
        alert('Login exitoso!');
        window.location.href = '/productos';
    } else {
        alert('Credenciales incorrectas');
    }
});

// Ejecutar según la página
if (window.location.pathname === '/productos') cargarProductos();
if (window.location.pathname === '/carrito') cargarCarrito();
if (window.location.pathname === '/pedidos') cargarPedidos();

// Registro
document.getElementById('registerForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const user = {
        nombre: document.getElementById('regNombre').value,
        email: document.getElementById('regEmail').value,
        password: document.getElementById('regPassword').value
    };

    const res = await fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user)
    });

    if (res.ok) {
        alert('¡Registro exitoso! Ahora puedes iniciar sesión.');
    } else {
        alert('Error al registrarse');
    }
});

// Login mejorado
document.getElementById('loginForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    const res = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });

    if (res.ok) {
        alert('Login exitoso!');
        window.location.href = '/productos';
    } else {
        alert('Credenciales incorrectas');
    }
});