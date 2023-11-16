package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"sync"
)

// featureFlags armazena o estado de várias feature flags.
var (
	featureFlags      = make(map[string]bool)
	featureFlagsMutex sync.Mutex
)

// main é a função principal que inicia o servidor Fiber.
func main() {
	// Parsing de flags de linha de comando
	flag.Parse()

	// Inicializa o Fiber
	app := fiber.New()

	// Rota para obter o estado atual de todas as feature flags
	app.Get("/feature-status", getFeatureStatus)

	// Rota para atualizar o estado de uma feature flag específica
	app.Post("/update-feature/:flagName", updateFeature)

	// Endpoint protegido pela feature flag "exampleFlag1"
	app.Get("/seu-endpoint-1", FeatureFlagMiddleware("exampleFlag1"), yourHandler)

	// Endpoint protegido pela feature flag "exampleFlag2"
	app.Get("/seu-endpoint-2", FeatureFlagMiddleware("exampleFlag2"), yourHandler)

	// Inicia o servidor na porta 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

// FeatureFlagMiddleware retorna um middleware para verificar o estado de uma feature flag específica.
func FeatureFlagMiddleware(flagName string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Verifica o estado da feature flag específica
		if isEnabled := getFeatureFlag(flagName); isEnabled {
			// Se a feature está habilitada, continua para o próximo middleware ou manipulador
			return c.Next()
		} else {
			// Se a feature está desabilitada, retorna um status HTTP 403 Forbidden
			return c.Status(http.StatusForbidden).SendString(fmt.Sprintf("Feature %s está desabilitada.", flagName))
		}
	}
}

// yourHandler é o manipulador principal que é chamado quando a feature flag está habilitada.
func yourHandler(c *fiber.Ctx) error {
	// Lógica quando a feature está habilitada (retorna código 200)
	return c.SendStatus(http.StatusOK)
}

// getFeatureStatus retorna o estado atual de todas as feature flags.
func getFeatureStatus(c *fiber.Ctx) error {
	featureFlagsMutex.Lock()
	defer featureFlagsMutex.Unlock()

	return c.JSON(featureFlags)
}

// updateFeature atualiza o estado de uma feature flag específica.
func updateFeature(c *fiber.Ctx) error {
	flagName := c.Params("flagName")

	var requestBody struct {
		EnableFeature bool `json:"enableFeature"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Erro ao analisar a solicitação.")
	}

	setFeatureFlag(flagName, requestBody.EnableFeature)

	return c.SendString(fmt.Sprintf("Feature %s atualizada para: %v", flagName, getFeatureFlag(flagName)))
}

// setFeatureFlag define o estado de uma feature flag.
func setFeatureFlag(flagName string, value bool) {
	featureFlagsMutex.Lock()
	defer featureFlagsMutex.Unlock()

	featureFlags[flagName] = value
}

// getFeatureFlag retorna o estado de uma feature flag.
func getFeatureFlag(flagName string) bool {
	featureFlagsMutex.Lock()
	defer featureFlagsMutex.Unlock()

	return featureFlags[flagName]
}
