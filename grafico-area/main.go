package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico de área para evolución acumulativa de ventas
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Ventas Acumulativas por Región",
			Subtitle: "Datos trimestrales 2025",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"América", "Europa", "Asia", "África"},
		}),
	)

	trimestres := []string{"Q1", "Q2", "Q3", "Q4"}

	america := []opts.LineData{
		{Value: 150000}, {Value: 280000}, {Value: 420000}, {Value: 580000},
	}

	europa := []opts.LineData{
		{Value: 120000}, {Value: 230000}, {Value: 340000}, {Value: 450000},
	}

	asia := []opts.LineData{
		{Value: 200000}, {Value: 380000}, {Value: 560000}, {Value: 750000},
	}

	africa := []opts.LineData{
		{Value: 50000}, {Value: 90000}, {Value: 130000}, {Value: 180000},
	}

	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Trimestre",
			Data: trimestres,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Ventas ($)",
			Type: "value",
		}),
	)

	line.AddSeries("América", america)
	line.AddSeries("Europa", europa)
	line.AddSeries("Asia", asia)
	line.AddSeries("África", africa)

	f, _ := os.Create("ventas_regiones.html")
	defer f.Close()
	line.Render(f)

	fmt.Println("✓ Gráfico de área guardado: ventas_regiones.html")
	openBrowser("ventas_regiones.html")
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
