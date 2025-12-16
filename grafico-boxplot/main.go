package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	boxplot := charts.NewBoxPlot()
	boxplot.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Análisis de Distribución - Tiempos de Procesamiento",
			Subtitle: "Box plot por servidor",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
	)

	servidores := []string{"Servidor 1", "Servidor 2", "Servidor 3", "Servidor 4"}

	datosServidores := [][]interface{}{
		{25, 45, 75, 10, 95},
		{30, 50, 80, 15, 100},
		{20, 40, 70, 8, 90},
		{35, 55, 85, 20, 105},
	}

	for i, servidor := range servidores {
		boxplot.AddSeries(servidor, []opts.BoxPlotData{
			{Value: datosServidores[i]},
		})
	}

	boxplot.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Servidores",
			Type: "category",
			Data: servidores,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Tiempo (ms)",
			Type: "value",
		}),
	)

	f, _ := os.Create("boxplot_estadistico.html")
	defer f.Close()
	boxplot.Render(f)

	fmt.Println("✓ Box plot guardado: boxplot_estadistico.html")
	openBrowser("boxplot_estadistico.html")
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
