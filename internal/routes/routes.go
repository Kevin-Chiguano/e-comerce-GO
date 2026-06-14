package routes

import (
	"database/sql"
	"ecommerce-manager/internal/handlers"
	"ecommerce-manager/internal/repositories"
	"ecommerce-manager/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// Repositorios
	userRepo := repositories.NewUsuarioRepository(db)
	prodRepo := repositories.NewProductoRepository(db)
	carritoRepo := repositories.NewCarritoRepository(db)
	pedidoRepo := repositories.NewPedidoRepository(db)

	// Servicios
	authService := services.NewAuthService(userRepo)
	productoService := services.NewProductoService(prodRepo)
	carritoService := services.NewCarritoService(carritoRepo)
	pedidoService := services.NewPedidoService(pedidoRepo, carritoRepo)
	usuarioService := services.NewUsuarioService(userRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	productoHandler := handlers.NewProductoHandler(productoService)
	carritoHandler := handlers.NewCarritoHandler(carritoService)
	pedidoHandler := handlers.NewPedidoHandler(pedidoService)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService, authService)

	// ==================== RUTAS HTML ====================
	r.GET("/", func(c *gin.Context) {
		c.File("./web/templates/index.html")
	})
	r.GET("/login", func(c *gin.Context) {
		c.File("./web/templates/login.html")
	})
	r.GET("/productos", func(c *gin.Context) {
		c.File("./web/templates/productos.html")
	})
	r.GET("/carrito", func(c *gin.Context) {
		c.File("./web/templates/carrito.html")
	})
	r.GET("/pedidos", func(c *gin.Context) {
		c.File("./web/templates/pedidos.html")
	})

	// ==================== API Routes ====================
	api := r.Group("/api")
	{
		api.POST("/register", usuarioHandler.Register)
		api.POST("/login", authHandler.Login)

		api.GET("/productos", productoHandler.GetAll)
		api.GET("/productos/:id", productoHandler.GetByID)

		api.GET("/carrito", carritoHandler.GetCarrito)
		api.POST("/carrito", carritoHandler.AddItem)

		api.POST("/pedidos", pedidoHandler.CrearPedido)
		api.GET("/pedidos", pedidoHandler.GetMisPedidos)
	}
}