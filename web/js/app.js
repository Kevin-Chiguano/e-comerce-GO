// ==================== CONFIGURACIÓN JWT ====================
function getToken() {
    return localStorage.getItem('token');
}

async function fetchWithAuth(url, options = {}) {
    const token = getToken();
    if (!token && (url.includes('/api/carrito') || url.includes('/api/pedidos'))) {
        alert('Debes iniciar sesión primero');
        window.location.href = '/login';
        return;
    }

    options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
    };

    return fetch(url, options);
}

// ==================== PRODUCTOS ====================
async function cargarProductos() {
    const res = await fetch('/api/productos');
    if (!res.ok) return;
    
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

async function agregarAlCarrito(productoId) {
    const res = await fetchWithAuth('/api/carrito', {
        method: 'POST',
        body: JSON.stringify({ producto_id: productoId, cantidad: 1 })
    });

    if (!res) return;

    const data = await res.json().catch(() => null);
    if (res.ok) {
        alert('✅ Producto añadido al carrito!');
        return;
    }

    alert('Error al añadir al carrito: ' + (data?.error || res.statusText));
}

// ==================== CARRITO ====================
async function cargarCarrito() {
    const res = await fetchWithAuth('/api/carrito');
    if (!res || !res.ok) return;

    const carrito = await res.json();
    const container = document.getElementById('carrito-list');
    let total = 0;

    if (!container) return;

    container.innerHTML = '';
    if (carrito.detalles && carrito.detalles.length > 0) {
        carrito.detalles.forEach(item => {
            const subtotal = item.cantidad * item.producto.precio;
            total += subtotal;
            container.innerHTML += `
                <div class="producto">
                    <h4>${item.producto.nombre}</h4>
                    <p>Cantidad: ${item.cantidad} × $${item.producto.precio}</p>
                    <p><strong>Subtotal: $${subtotal}</strong></p>
                </div>
            `;
        });
    } else {
        container.innerHTML = '<p>Tu carrito está vacío</p>';
    }

    const totalElement = document.getElementById('total');
    if (totalElement) {
        totalElement.innerHTML = `<h2>Total: $${total.toFixed(2)}</h2>`;
    }
}

async function crearPedido() {
    const res = await fetchWithAuth('/api/pedidos', { method: 'POST' });
    if (res && res.ok) {
        alert('🎉 Pedido realizado con éxito!');
        window.location.href = '/pedidos';
    } else {
        alert('Error al crear el pedido');
    }
}

// ==================== PEDIDOS ====================
async function cargarPedidos() {
    const container = document.getElementById('pedidos-list');
    if (!container) return;
    container.innerHTML = '<p>Cargando pedidos...</p>';
    
    // Implementación básica por ahora
    container.innerHTML = `
        <p>Funcionalidad de "Mis Pedidos" en desarrollo.</p>
        <p>Próximamente podrás ver tu historial aquí.</p>
    `;
}

// ==================== AUTH ====================
// === LOGIN ===
document.getElementById('loginForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const email = document.getElementById('loginEmail').value.trim();
    const password = document.getElementById('loginPassword').value.trim();

    if (!email || !password) {
        alert("Por favor completa todos los campos");
        return;
    }

    try {
        const res = await fetch('/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await res.json();

        if (res.ok) {
            localStorage.setItem('token', data.token);
            alert('✅ Login exitoso!');
            window.location.href = '/productos';
        } else {
            console.error("Error login:", data);
            alert('❌ Credenciales incorrectas: ' + (data.error || ''));
        }
    } catch (err) {
        console.error(err);
        alert('Error de conexión');
    }
});

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
        alert('✅ Registro exitoso! Ahora inicia sesión.');
        window.location.href = '/login';
    } else {
        alert('❌ Error al registrarse');
    }
});

// ==================== EJECUCIÓN AUTOMÁTICA ====================
if (window.location.pathname === '/productos') cargarProductos();
if (window.location.pathname === '/carrito') cargarCarrito();
if (window.location.pathname === '/pedidos') cargarPedidos();