package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico de líneas múltiples para tendencias de precios
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Evolución de Precios - Últimos 12 Meses",
			Subtitle: "Comparativa de criptomonedas",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"Bitcoin", "Ethereum", "Solana"},
		}),
	)

	meses := []string{"Ene", "Feb", "Mar", "Abr", "May", "Jun",
		"Jul", "Ago", "Sep", "Oct", "Nov", "Dic"}

	bitcoin := []opts.LineData{
		{Value: 42000}, {Value: 45000}, {Value: 43000}, {Value: 46000},
		{Value: 48000}, {Value: 50000}, {Value: 52000}, {Value: 51000},
		{Value: 49000}, {Value: 51000}, {Value: 53000}, {Value: 55000},
	}

	ethereum := []opts.LineData{
		{Value: 2500}, {Value: 2700}, {Value: 2600}, {Value: 2900},
		{Value: 3000}, {Value: 3100}, {Value: 3200}, {Value: 3100},
		{Value: 3000}, {Value: 3200}, {Value: 3300}, {Value: 3500},
	}

	solana := []opts.LineData{
		{Value: 140}, {Value: 160}, {Value: 150}, {Value: 180},
		{Value: 200}, {Value: 210}, {Value: 220}, {Value: 210},
		{Value: 200}, {Value: 220}, {Value: 240}, {Value: 260},
	}

	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Mes",
			Data: meses,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Precio (USD)",
			Type: "value",
		}),
	)

	line.AddSeries("Bitcoin", bitcoin)
	line.AddSeries("Ethereum", ethereum)
	line.AddSeries("Solana", solana)

	f, _ := os.Create("tendencia_precios.html")
	defer f.Close()
	line.Render(f)

	fmt.Println("✓ Gráfico de líneas múltiples guardado: tendencia_precios.html")
	openBrowser("tendencia_precios.html")
}

func openBrowser(archivo string) {
	rutaAbs, _ := os.Getwd()
	rutaCompleta := rutaAbs + "/" + archivo

	var cmd *exec.Cmd
	switch {
	case exec.Command("xdg-open", "--version").Run() == nil:
		cmd = exec.Command("xdg-open", rutaCompleta)
	case exec.Command("open", "--version").Run() == nil:
		cmd = exec.Command("open", rutaCompleta)
	default:
		cmd = exec.Command("cmd", "/c", "start", rutaCompleta)
	}
	cmd.Run()
}
