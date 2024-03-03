package handlers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/Tracking-Detector/td-model-utils/env"
	"github.com/gofiber/fiber/v2"
)

type ModelHandler struct {
	app *fiber.App
}

func NewModelHandler(app *fiber.App) *ModelHandler {
	return &ModelHandler{app: app}
}

func (mh *ModelHandler) PostModel(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	defer form.RemoveAll()

	modelName := c.Params("modelName")

	modelDir := filepath.Join(env.GetModelPath(), modelName)
	if _, err := os.Stat(modelDir); os.IsNotExist(err) {
		os.MkdirAll(modelDir, os.ModePerm)
	}

	modelJSONFile := filepath.Join(modelDir, "model.json")
	modelJSON, err := form.File["model.json"][0].Open()
	if err != nil {
		return err
	}
	defer modelJSON.Close()

	modelJSONContent, err := io.ReadAll(modelJSON)
	if err != nil {
		return err
	}
	if err := os.WriteFile(modelJSONFile, modelJSONContent, 0644); err != nil {
		return err
	}
	weightsFile := filepath.Join(modelDir, "model.weights.bin")
	weights, err := form.File["model.weights.bin"][0].Open()
	if err != nil {
		return err
	}
	defer weights.Close()

	// Create the output file
	outputFile, err := os.Create(weightsFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Copy the contents of the uploaded file to the output file
	_, err = io.Copy(outputFile, weights)
	if err != nil {
		return err
	}

	return c.SendString("Model saved successfully")
}

func (mh *ModelHandler) RegisterRoutes() {
	mh.app.Static("/models", env.GetModelPath())
	mh.app.Post("/models/:modelName", mh.PostModel)
}
