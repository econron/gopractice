package main

import (
	"log/slog"
	"os"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"context"
)

func main() {
	name := "Gopher"
	value := 100
	slog.Info("Hello, ", name)
	slog.Error("This is an error message")

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	logger.Info("fruit", "name", name, "value", value)
	logger.Info("This is info log", slog.Group("attr",
		slog.String("name", "banana"),
		slog.Int("count", 1),
	))

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		logger.Error("client error", slog.Any("error", err))
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	model.SetTemperature(0.0)
	model.SafetySettings = []*genai.SafetySetting{
		{
		  Category:  genai.HarmCategoryHarassment,
		  Threshold: genai.HarmBlockOnlyHigh,
		},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(`二条城　541 Nijojocho, Nakagyo Ward, Kyoto, 604-8301
法観寺　〒605-0862 Kyoto, Higashiyama Ward, 清水八坂上町388
ジャズクラブ　〒604-8082 Kyoto, Nakagyo Ward, 御幸町西入弁慶石町48

ここを巡るための最短ルートを教えていただけますか？`))
	if err != nil {
		logger.Error("model error", slog.Any("error", err))
	}

	logger.Info("response", slog.Any("resp", resp))

	resp2, err := model.CountTokens(ctx, genai.Text(`二条城　541 Nijojocho, Nakagyo Ward, Kyoto, 604-8301
法観寺　〒605-0862 Kyoto, Higashiyama Ward, 清水八坂上町388
ジャズクラブ　〒604-8082 Kyoto, Nakagyo Ward, 御幸町西入弁慶石町48

ここを巡るための最短ルートを教えていただけますか？`))
	if err != nil {
		logger.Error("count error", slog.Any("error", err))
	}
	logger.Info("count", slog.Any("resp", resp2))
}