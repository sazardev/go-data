package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	// Crear gráfico de línea
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Gráfico de f(x) = x²",
			Subtitle: "Parábola cuadrática",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "x",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "f(x)",
		}),
	)

	// Generar datos para x² desde -10 a 10
	var xAxis []string
	var yAxis []opts.LineData

	for x := -10.0; x <= 10.0; x += 0.5 {
		xAxis = append(xAxis, fmt.Sprintf("%.1f", x))
		y := x * x
		yAxis = append(yAxis, opts.LineData{Value: y})
	}

	// Configurar eje X y agregar datos
	line.SetGlobalOptions(
		charts.WithXAxisOpts(opts.XAxis{
			Name: "x",
			Data: xAxis,
		}),
	)

	// Agregar los datos al gráfico
	line.AddSeries("f(x) = x²", yAxis)

	// Guardar como HTML
	archivo := "grafico_x2.html"
	f, err := os.Create(archivo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	if err := line.Render(f); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Gráfico interactivo guardado como '%s'\n", archivo)

	// Abrir en navegador automáticamente
	openBrowser(archivo)
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

	if err := cmd.Run(); err != nil {
		fmt.Printf("Abre manualmente: %s\n", rutaCompleta)
	}
}
