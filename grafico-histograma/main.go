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
			Title:    "Histograma - Distribución de Tiempos de Respuesta",
			Subtitle: "Servidor de aplicación (ms)",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
	)

	rangos := []string{"0-50ms", "50-100ms", "100-150ms", "150-200ms", "200-250ms", "250-300ms"}
	frecuencias := []opts.BarData{
		{Value: 450}, {Value: 320}, {Value: 180}, {Value: 95}, {Value: 45}, {Value: 10},
	}

	bar.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Rango de Tiempos",
			Data: rangos,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Frecuencia",
			Type: "value",
		}),
	)

	bar.AddSeries("Frecuencia", frecuencias)

	f, _ := os.Create("histograma_tiempos.html")
	defer f.Close()
	bar.Render(f)

	fmt.Println("✓ Histograma guardado: histograma_tiempos.html")
	openBrowser("histograma_tiempos.html")
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
