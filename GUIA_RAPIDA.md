# 🎨 Guía Rápida - Gráficos Go-Echarts

## 📂 Carpetas Disponibles (15 Ejemplos)

### 1️⃣ math-data - Scatter (Función Matemática)
```bash
cd math-data && go run main.go
# Genera: grafico_x2.html
# Uso: Visualizar función f(x) = x²
```
📖 [Ver README](math-data/README.md)

---

### 2️⃣ grafico-barras - Bar Chart (Comparación)
```bash
cd grafico-barras && go run main.go
# Genera: comparacion_precios.html
# Uso: Comparar precios entre tiendas
```
📖 [Ver README](grafico-barras/README.md)

---

### 3️⃣ grafico-pastel - Pie Chart (Composición)
```bash
cd grafico-pastel && go run main.go
# Genera: distribucion_gastos.html
# Uso: Mostrar composición de gastos
```
📖 [Ver README](grafico-pastel/README.md)

---

### 4️⃣ grafico-lineas-multiples - Line Chart (Tendencias)
```bash
cd grafico-lineas-multiples && go run main.go
# Genera: tendencia_precios.html
# Uso: Comparar evolución de precios de criptomonedas
```
📖 [Ver README](grafico-lineas-multiples/README.md)

---

### 5️⃣ grafico-area - Area Chart (Acumulación)
```bash
cd grafico-area && go run main.go
# Genera: ventas_regiones.html
# Uso: Mostrar ventas acumulativas por región
```
📖 [Ver README](grafico-area/README.md)

---

### 6️⃣ grafico-heatmap - Heatmap (Correlación)
```bash
cd grafico-heatmap && go run main.go
# Genera: matriz_correlacion.html
# Uso: Mostrar correlación entre variables financieras
```
📖 [Ver README](grafico-heatmap/README.md)

---

### 7️⃣ grafico-radar - Radar Chart (Multidimensional)
```bash
cd grafico-radar && go run main.go
# Genera: desempenio_equipos.html
# Uso: Comparar desempeño de equipos en 5 dimensiones
```
📖 [Ver README](grafico-radar/README.md)

---

### 8️⃣ grafico-histograma - Histogram (Distribución)
```bash
cd grafico-histograma && go run main.go
# Genera: histograma_tiempos.html
# Uso: Mostrar distribución de tiempos de respuesta
```
📖 [Ver README](grafico-histograma/README.md)

---

### 9️⃣ grafico-gauge - Gauge Chart (KPIs)
```bash
cd grafico-gauge && go run main.go
# Genera: kpis_negocio.html
# Uso: Mostrar indicadores empresariales (tipo velocímetro)
```
📖 [Ver README](grafico-gauge/README.md)

---

### 🔟 grafico-sankey - Sankey Diagram (Flujos)
```bash
cd grafico-sankey && go run main.go
# Genera: flujo_financiero.html
# Uso: Mostrar cómo fluye el dinero (ingresos → gastos)
```
📖 [Ver README](grafico-sankey/README.md)

---

### 1️⃣1️⃣ grafico-funnel - Funnel Chart (Conversión)
```bash
cd grafico-funnel && go run main.go
# Genera: embudo_conversion.html
# Uso: Mostrar embudo de conversión e-commerce
```
📖 [Ver README](grafico-funnel/README.md)

---

### 1️⃣2️⃣ grafico-treemap - Bar Chart (Presupuesto)
```bash
cd grafico-treemap && go run main.go
# Genera: distribucion_presupuesto.html
# Uso: Mostrar distribución de presupuesto por departamento
```
📖 [Ver README](grafico-treemap/README.md)

---

### 1️⃣3️⃣ grafico-burbuja - Bubble Chart (3D Analysis)
```bash
cd grafico-burbuja && go run main.go
# Genera: analisis_productos.html
# Uso: Análisis de productos (precio vs demanda vs margen)
```
📖 [Ver README](grafico-burbuja/README.md)

---

### 1️⃣4️⃣ grafico-candlestick - Candlestick Chart (Trading)
```bash
cd grafico-candlestick && go run main.go
# Genera: analisis_tecnico.html
# Uso: Análisis técnico de precios (velas)
```
📖 [Ver README](grafico-candlestick/README.md)

---

### 1️⃣5️⃣ grafico-boxplot - Box Plot (Estadística)
```bash
cd grafico-boxplot && go run main.go
# Genera: boxplot_estadistico.html
# Uso: Comparar distribuciones estadísticas de servidores
```
📖 [Ver README](grafico-boxplot/README.md)

---

## 🎯 Elegir el Gráfico Adecuado

### ❓ ¿Quiero mostrar **composición** (partes de un total)?
→ Usa: **Pie** (grafico-pastel) o **Bar**

### ❓ ¿Quiero mostrar **tendencias** en el tiempo?
→ Usa: **Line** (grafico-lineas-multiples) o **Area**

### ❓ ¿Quiero **comparar** valores entre categorías?
→ Usa: **Bar** (grafico-barras) o **Radar**

### ❓ ¿Quiero mostrar **relación** entre 2 variables?
→ Usa: **Scatter** (math-data) o **Heatmap**

### ❓ ¿Quiero mostrar **KPI o estado actual**?
→ Usa: **Gauge** (grafico-gauge)

### ❓ ¿Quiero mostrar **flujo de dinero o recursos**?
→ Usa: **Sankey** (grafico-sankey)

### ❓ ¿Quiero mostrar **embudo de conversión**?
→ Usa: **Funnel** (grafico-funnel)

### ❓ ¿Quiero mostrar **análisis técnico de precios**?
→ Usa: **Candlestick** (grafico-candlestick)

### ❓ ¿Quiero mostrar **distribución estadística**?
→ Usa: **BoxPlot** (grafico-boxplot) o **Histogram**

### ❓ ¿Quiero mostrar **múltiples dimensiones** de forma compacta?
→ Usa: **Radar** (grafico-radar) o **Bubble** (grafico-burbuja)

---

## 📚 Documentación

### Documento Principal
📄 [CATALOGO_GRAFICOS.md](CATALOGO_GRAFICOS.md) - Guía completa con:
- Todos los tipos disponibles en go-echarts
- Cuáles están implementados
- Cuáles aún no (y por qué)
- Matriz de características

### README por Tipo
Cada carpeta contiene `README.md` con:
- ✅ Cuándo usar este gráfico
- ✅ Cómo funciona
- ✅ Interpretación de datos
- ✅ Aplicaciones reales
- ✅ Análisis del caso específico
- ✅ Recomendaciones de mejora

---

## 🚀 Primeros Pasos

### 1. Ejecutar todos los gráficos
```bash
cd /home/sazardev/Documents/golang/go-data

# Ejecutar uno a uno y explorar
for dir in grafico-* math-data; do
  cd "$dir"
  echo "Ejecutando $dir..."
  go run main.go 2>&1 | head -1
  cd ..
  sleep 1
done
```

### 2. Abrir en navegador
Cada gráfico genera automáticamente un archivo HTML que se abre en tu navegador:
- Interactivo (zoom, pan, tooltip)
- Descargable como imagen
- Leyenda interactiva

### 3. Estudiar el código
```bash
# Ejemplo: Ver estructura de Bar Chart
cat grafico-barras/main.go
```

### 4. Leer la documentación
```bash
# Ejemplo: Entender cuándo usar cada gráfico
cat grafico-barras/README.md
```

---

## 🎨 Estructura Común del Código

Todos los gráficos siguen este patrón:

```go
package main

import (
    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
    "os"
)

func main() {
    // 1. Crear gráfico
    chart := charts.NewXXX()
    
    // 2. Configurar opciones globales
    chart.SetGlobalOptions(
        charts.WithTitleOpts(opts.Title{...}),
        charts.WithTooltipOpts(opts.Tooltip{...}),
        // ...
    )
    
    // 3. Preparar datos
    items := []opts.XxxData{...}
    
    // 4. Agregar series
    chart.AddSeries("nombre", items)
    
    // 5. Guardar a archivo
    f, _ := os.Create("salida.html")
    defer f.Close()
    chart.Render(f)
    
    // 6. Abrir en navegador
    openBrowser("salida.html")
}

func openBrowser(archivo string) {
    // Código para abrir en navegador (cross-platform)
}
```

---

## 📊 Tabla de Referencia Rápida

| Gráfico     | Mejor Para        | Ejemplo de Datos | Típico % de Casos |
| ----------- | ----------------- | ---------------- | ----------------- |
| **Bar**     | Comparación       | Precios, ventas  | 20%               |
| **Line**    | Tendencias        | Series temporal  | 25%               |
| **Pie**     | Composición       | Porcentajes      | 15%               |
| **Scatter** | Relaciones        | Correlación      | 10%               |
| **Gauge**   | KPIs              | Métricas         | 10%               |
| **Funnel**  | Conversión        | Embudo           | 5%                |
| **Otros**   | Casos específicos | Variado          | 15%               |

---

## 🔧 Personalización

Cada gráfico es editable. Ejemplos:

### Cambiar Colores
```go
chart.SetGlobalOptions(
    charts.WithColorsOpts(opts.Colors{
        "#FF6B6B", "#4ECDC4", "#45B7D1",
    }),
)
```

### Cambiar Título
```go
charts.WithTitleOpts(opts.Title{
    Title:    "Mi Nuevo Título",
    Subtitle: "Subtítulo",
})
```

### Cambiar Datos
Edita los arrays de datos en `main.go` y vuelve a ejecutar:
```bash
go run main.go
```

---

## ✨ Casos de Uso Rápidos

### Dashboard de Ventas
Usa: **Line** (ingresos en tiempo) + **Bar** (por región)

### Análisis de Producto
Usa: **Bubble** (precio vs demanda vs margen)

### Reporte Financiero
Usa: **Sankey** (ingresos → gastos) + **Gauge** (KPIs)

### Análisis de Calidad
Usa: **BoxPlot** (distribución) + **Heatmap** (correlación)

### Estrategia de Marketing
Usa: **Funnel** (conversión) + **Line** (tendencia)

---

## 💡 Tips Profesionales

1. **Elige un gráfico para cada pregunta** - No sobrecargues
2. **Contextualiza con datos** - Incluye benchmarks o metas
3. **Usa interactividad** - Aprovecha el hover y leyenda
4. **Crea historias** - Múltiples gráficos que se conecten
5. **Simplifica** - Menos es más en visualización

---

## 📞 Recursos

- 🌐 **Sitio oficial**: https://go-echarts.github.io/go-echarts/
- 📚 **GitHub**: https://github.com/go-echarts/go-echarts
- 📖 **Go Docs**: https://pkg.go.dev/github.com/go-echarts/go-echarts/v2
- 🎓 **ECharts Original**: https://echarts.apache.org/

---

## 🎯 Próximos Pasos

1. Explora cada gráfico ejecutándolos
2. Lee el README de cada uno
3. Modifica datos y vuelve a ejecutar
4. Adapta a tus propios datos
5. Crea dashboards combinando múltiples gráficos

---

**¡Listo para empezar! Elige un gráfico y explora.** 🚀
