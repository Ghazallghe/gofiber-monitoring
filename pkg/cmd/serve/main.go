package serve

import (
	"os"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/configs"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func main() {
	app := fiber.New()
	initialization()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetUpRoutes(app)

	port := os.Getenv("SERVER_PORT")
	connection := ":" + port
	app.Listen(connection)
}

func initialization() {
	configs.SetUpConfigs()
	db.SetUpDB()
}

func New() *cobra.Command {
	// nolint: exhaustruct
	return &cobra.Command{
		Use:   "serve",
		Short: "runs http server for HTTP monitoring api",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
}
