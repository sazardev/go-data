package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico de barras para comparación de precios
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Comparación de Precios por Producto",
			Subtitle: "Análisis de mercado 2025",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithLegendOpts(opts.Legend{
			Left: "center",
		}),
	)

	// Datos de productos y precios en diferentes tiendas
	productos := []string{"Laptop", "Monitor", "Teclado", "Mouse", "Auriculares"}
	tienda1 := []opts.BarData{
		{Value: 1200}, {Value: 350}, {Value: 120}, {Value: 35}, {Value: 85},
	}
	tienda2 := []opts.BarData{
		{Value: 1150}, {Value: 380}, {Value: 110}, {Value: 40}, {Value: 95},
	}
	tienda3 := []opts.BarData{
		{Value: 1300}, {Value: 320}, {Value: 130}, {Value: 30}, {Value: 75},
	}

	bar.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Productos",
			Data: productos,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Precio ($)",
			Type: "value",
		}),
	)

	bar.AddSeries("TechStore", tienda1)
	bar.AddSeries("ElectroMart", tienda2)
	bar.AddSeries("CompuWorld", tienda3)

	f, _ := os.Create("comparacion_precios.html")
	defer f.Close()
	bar.Render(f)

	fmt.Println("✓ Gráfico de barras guardado: comparacion_precios.html")
	openBrowser("comparacion_precios.html")
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
