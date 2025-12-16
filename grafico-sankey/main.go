package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Flujo de Ingresos y Gastos - Análisis Financiero",
			Subtitle: "Desglose presupuestario mensual",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
	)

	nodes := []opts.SankeyNode{
		{Name: "Ingresos Totales"},
		{Name: "Impuestos"},
		{Name: "Salarios"},
		{Name: "Operativo"},
		{Name: "Marketing"},
		{Name: "R&D"},
		{Name: "Neto Retenido"},
	}

	links := []opts.SankeyLink{
		{Source: "Ingresos Totales", Target: "Impuestos", Value: 25000},
		{Source: "Ingresos Totales", Target: "Salarios", Value: 60000},
		{Source: "Ingresos Totales", Target: "Operativo", Value: 35000},
		{Source: "Ingresos Totales", Target: "Marketing", Value: 20000},
		{Source: "Ingresos Totales", Target: "R&D", Value: 15000},
		{Source: "Salarios", Target: "Neto Retenido", Value: 60000},
		{Source: "Operativo", Target: "Neto Retenido", Value: 35000},
		{Source: "Marketing", Target: "Neto Retenido", Value: 20000},
		{Source: "R&D", Target: "Neto Retenido", Value: 15000},
	}

	sankey.AddSeries("sankey", nodes, links)

	f, _ := os.Create("flujo_financiero.html")
	defer f.Close()
	sankey.Render(f)

	fmt.Println("✓ Diagrama Sankey guardado: flujo_financiero.html")
	openBrowser("flujo_financiero.html")
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
