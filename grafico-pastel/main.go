package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico de pastel para distribución de gastos
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Distribución de Gastos Mensuales",
			Subtitle: "Presupuesto familiar - Diciembre 2025",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
		charts.WithLegendOpts(opts.Legend{
			Orient: "vertical",
			Left:   "left",
		}),
	)

	// Datos de gastos
	items := []opts.PieData{
		{Name: "Vivienda", Value: 1200},
		{Name: "Alimentación", Value: 600},
		{Name: "Transporte", Value: 300},
		{Name: "Servicios", Value: 250},
		{Name: "Entretenimiento", Value: 200},
		{Name: "Educación", Value: 400},
		{Name: "Ahorros", Value: 500},
		{Name: "Otros", Value: 150},
	}

	pie.AddSeries("gastos", items,
		charts.WithPieChartOpts(opts.PieChart{
			Radius: []string{"30%", "75%"},
		}),
	)

	f, _ := os.Create("distribucion_gastos.html")
	defer f.Close()
	pie.Render(f)

	fmt.Println("✓ Gráfico de pastel guardado: distribucion_gastos.html")
	openBrowser("distribucion_gastos.html")
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
