package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Análisis Técnico - Cotización de Acciones",
			Subtitle: "ACME Corp - Últimas 10 sesiones",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithLegendOpts(opts.Legend{
			Data: []string{"Precio"},
		}),
	)

	dias := []string{"Día 1", "Día 2", "Día 3", "Día 4", "Día 5",
		"Día 6", "Día 7", "Día 8", "Día 9", "Día 10"}

	precios := []opts.KlineData{
		{Value: []interface{}{100, 105, 95, 110}},
		{Value: []interface{}{105, 103, 100, 108}},
		{Value: []interface{}{103, 108, 102, 115}},
		{Value: []interface{}{108, 112, 105, 118}},
		{Value: []interface{}{112, 110, 108, 120}},
		{Value: []interface{}{110, 115, 109, 122}},
		{Value: []interface{}{115, 118, 112, 125}},
		{Value: []interface{}{118, 120, 115, 128}},
		{Value: []interface{}{120, 122, 118, 130}},
		{Value: []interface{}{122, 125, 120, 132}},
	}

	kline.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Fecha",
			Type: "category",
			Data: dias,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Precio ($)",
			Type: "value",
		}),
	)

	kline.AddSeries("Precio", precios)

	f, _ := os.Create("analisis_tecnico.html")
	defer f.Close()
	kline.Render(f)

	fmt.Println("✓ Gráfico candlestick guardado: analisis_tecnico.html")
	openBrowser("analisis_tecnico.html")
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
