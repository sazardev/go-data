package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico gauge para KPIs
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Indicadores Clave de Rendimiento (KPIs)",
			Subtitle: "Estado actual del negocio",
		}),
	)

	kpis := []string{"Satisfacción", "Rentabilidad", "Crecimiento", "Eficiencia"}
	valores := []interface{}{87.5, 72.3, 91.2, 68.8}

	for i, kpi := range kpis {
		gauge.AddSeries(kpi, []opts.GaugeData{
			{Value: valores[i], Name: kpi},
		})
	}

	f, _ := os.Create("kpis_negocio.html")
	defer f.Close()
	gauge.Render(f)

	fmt.Println("✓ Gráfico gauge guardado: kpis_negocio.html")
	openBrowser("kpis_negocio.html")
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
