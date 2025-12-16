package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	bubble := charts.NewEffectScatter()
	bubble.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Análisis de Rendimiento de Productos",
			Subtitle: "Precio vs Demanda vs Rentabilidad",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"Producto A", "Producto B", "Producto C", "Producto D"},
		}),
	)

	productoA := []opts.EffectScatterData{
		{Value: []interface{}{100, 500, 30}},
		{Value: []interface{}{150, 450, 45}},
		{Value: []interface{}{200, 400, 50}},
	}

	productoB := []opts.EffectScatterData{
		{Value: []interface{}{80, 700, 25}},
		{Value: []interface{}{120, 650, 40}},
		{Value: []interface{}{160, 600, 55}},
	}

	productoC := []opts.EffectScatterData{
		{Value: []interface{}{250, 300, 60}},
		{Value: []interface{}{280, 350, 70}},
	}

	productoD := []opts.EffectScatterData{
		{Value: []interface{}{50, 900, 15}},
		{Value: []interface{}{75, 850, 28}},
	}

	bubble.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Precio ($)",
			Type: "value",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Demanda (unidades)",
			Type: "value",
		}),
	)

	bubble.AddSeries("Producto A", productoA)
	bubble.AddSeries("Producto B", productoB)
	bubble.AddSeries("Producto C", productoC)
	bubble.AddSeries("Producto D", productoD)

	f, _ := os.Create("analisis_productos.html")
	defer f.Close()
	bubble.Render(f)

	fmt.Println("✓ Gráfico de burbujas guardado: analisis_productos.html")
	openBrowser("analisis_productos.html")
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
