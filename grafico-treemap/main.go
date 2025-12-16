package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Distribución de Presupuesto Empresarial",
			Subtitle: "Desglose jerárquico por departamento",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
	)

	nombres := []string{"Ventas", "Desarrollo", "Marketing", "RH", "Legal", "Admin"}
	valores := []opts.BarData{
		{Value: 250000},
		{Value: 180000},
		{Value: 120000},
		{Value: 80000},
		{Value: 50000},
		{Value: 90000},
	}

	bar.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Type: "category",
			Data: nombres,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
		}),
	)

	bar.AddSeries("Presupuesto", valores)

	f, _ := os.Create("distribucion_presupuesto.html")
	defer f.Close()
	bar.Render(f)

	fmt.Println("✓ Distribución de presupuesto guardada: distribucion_presupuesto.html")
	openBrowser("distribucion_presupuesto.html")
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
