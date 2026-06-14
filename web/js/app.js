fetch('/api/productos')
  .then(res => res.json())
  .then(data => {
    const container = document.getElementById('productos');
    data.forEach(p => {
      container.innerHTML += `
        <div class="producto">
          <h3>${p.nombre}</h3>
          <p>$${p.precio}</p>
          <button onclick="addToCart(${p.id})">Añadir al carrito</button>
        </div>
      `;
    });
  });