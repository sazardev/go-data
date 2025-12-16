package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	radar := charts.NewRadar()

	// Definir indicadores del radar
	radarIndicators := []*opts.Indicator{
		{Name: "Productividad", Max: 100},
		{Name: "Calidad", Max: 100},
		{Name: "Innovación", Max: 100},
		{Name: "Comunicación", Max: 100},
		{Name: "Cumplimiento", Max: 100},
	}

	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Análisis de Desempeño de Equipos",
			Subtitle: "Evaluación multidimensional Q4 2025",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "item",
		}),
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"Equipo A", "Equipo B", "Equipo C"},
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator: radarIndicators,
		}),
	)

	radar.AddSeries("Equipo A", []opts.RadarData{
		{Value: []interface{}{92, 85, 78, 88, 95}},
	}).
		AddSeries("Equipo B", []opts.RadarData{
			{Value: []interface{}{88, 92, 85, 82, 90}},
		}).
		AddSeries("Equipo C", []opts.RadarData{
			{Value: []interface{}{80, 88, 92, 90, 85}},
		})

	f, _ := os.Create("desempenio_equipos.html")
	defer f.Close()
	radar.Render(f)

	fmt.Println("✓ Gráfico radar guardado: desempenio_equipos.html")
	openBrowser("desempenio_equipos.html")
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
