package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico funnel para embudo de conversión
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Embudo de Conversión - E-commerce",
			Subtitle: "Análisis de caída de usuarios en el proceso de compra",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
		charts.WithLegendOpts(opts.Legend{
			Orient: "vertical",
			Left:   "left",
		}),
	)

	items := []opts.FunnelData{
		{Name: "Visitas Totales", Value: 50000},
		{Name: "Productos Vistos", Value: 35000},
		{Name: "Carrito Agregado", Value: 18000},
		{Name: "Checkout Iniciado", Value: 8500},
		{Name: "Pago Completado", Value: 4200},
	}

	funnel.AddSeries("conversión", items)

	f, _ := os.Create("embudo_conversion.html")
	defer f.Close()
	funnel.Render(f)

	fmt.Println("✓ Gráfico funnel guardado: embudo_conversion.html")
	openBrowser("embudo_conversion.html")
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
