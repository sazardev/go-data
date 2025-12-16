package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Matriz de Correlación - Variables Financieras",
			Subtitle: "Análisis de relaciones entre variables",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
	)

	vars := []string{"Precio", "Volumen", "ROI", "Riesgo", "Liquidez", "Rentabilidad"}

	items := make([]opts.HeatMapData, 0)
	correlations := [][]float32{
		{1.0, 0.85, 0.92, -0.65, 0.78, 0.88},
		{0.85, 1.0, 0.71, -0.45, 0.82, 0.73},
		{0.92, 0.71, 1.0, -0.58, 0.69, 0.95},
		{-0.65, -0.45, -0.58, 1.0, -0.62, -0.61},
		{0.78, 0.82, 0.69, -0.62, 1.0, 0.70},
		{0.88, 0.73, 0.95, -0.61, 0.70, 1.0},
	}

	for i := 0; i < len(vars); i++ {
		for j := 0; j < len(vars); j++ {
			items = append(items, opts.HeatMapData{
				Name:  "",
				Value: []interface{}{i, j, correlations[i][j]},
			})
		}
	}

	hm.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Variables",
			Type: "category",
			Data: vars,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Variables",
			Type: "category",
			Data: vars,
		}),
	)

	hm.AddSeries("correlación", items)

	f, _ := os.Create("matriz_correlacion.html")
	defer f.Close()
	hm.Render(f)

	fmt.Println("✓ Gráfico de calor guardado: matriz_correlacion.html")
	openBrowser("matriz_correlacion.html")
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
