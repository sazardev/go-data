# 📊 Proyecto Go-Echarts: Visualización de Datos en Go

## 🎯 Proyecto Completado

Este proyecto es una **colección completa de 15 ejemplos de gráficos interactivos** implementados en Go utilizando la librería **go-echarts/v2**.

**Estado**: ✅ **COMPLETADO Y VALIDADO** - Todos los 15 proyectos compilados y ejecutados exitosamente sin errores.

---

## 📋 Documentación

Comienza por uno de estos documentos:

### 📚 Documentos Principales

1. **[GUIA_RAPIDA.md](GUIA_RAPIDA.md)** 🚀
   - Guía de inicio rápido
   - Cómo ejecutar cada gráfico
   - Cómo elegir el gráfico adecuado
   - Primeros pasos

2. **[CATALOGO_GRAFICOS.md](CATALOGO_GRAFICOS.md)** 📖
   - Inventario completo de los 25+ tipos disponibles
   - Cuáles están implementados (15)
   - Cuáles están disponibles pero no implementados (10+)
   - Clasificación por caso de uso
   - Recomendaciones de prioridad

### 📄 README Individual (En cada carpeta)
Cada carpeta de gráfico contiene un `README.md` detallado:
- Cuándo usar el gráfico
- Interpretación de datos
- Aplicaciones reales
- Análisis del caso específico
- Recomendaciones

---

## ✅ Gráficos Implementados (15)

### 📊 Análisis de Datos

| #   | Gráfico     | Carpeta                  | Salida                   | Descripción            |
| --- | ----------- | ------------------------ | ------------------------ | ---------------------- |
| 1   | **Scatter** | math-data                | grafico_x2.html          | Función f(x)=x²        |
| 2   | **Bar**     | grafico-barras           | comparacion_precios.html | Comparación de precios |
| 3   | **Pie**     | grafico-pastel           | distribucion_gastos.html | Distribución de gastos |
| 4   | **Line**    | grafico-lineas-multiples | tendencia_precios.html   | Tendencias de precios  |
| 5   | **Area**    | grafico-area             | ventas_regiones.html     | Ventas acumulativas    |

### 🔬 Análisis Estadístico

| #   | Gráfico       | Carpeta            | Salida                   | Descripción                 |
| --- | ------------- | ------------------ | ------------------------ | --------------------------- |
| 6   | **Heatmap**   | grafico-heatmap    | matriz_correlacion.html  | Correlación de variables    |
| 7   | **Radar**     | grafico-radar      | desempenio_equipos.html  | Análisis multidimensional   |
| 8   | **Histogram** | grafico-histograma | histograma_tiempos.html  | Distribución de frecuencias |
| 9   | **BoxPlot**   | grafico-boxplot    | boxplot_estadistico.html | Distribución estadística    |
| 10  | **Bubble**    | grafico-burbuja    | analisis_productos.html  | Análisis 3D de productos    |

### 💼 Análisis Empresarial

| #   | Gráfico           | Carpeta             | Salida                        | Descripción          |
| --- | ----------------- | ------------------- | ----------------------------- | -------------------- |
| 11  | **Gauge**         | grafico-gauge       | kpis_negocio.html             | KPIs empresariales   |
| 12  | **Sankey**        | grafico-sankey      | flujo_financiero.html         | Flujo de dinero      |
| 13  | **Funnel**        | grafico-funnel      | embudo_conversion.html        | Embudo de conversión |
| 14  | **Bar (Treemap)** | grafico-treemap     | distribucion_presupuesto.html | Presupuesto          |
| 15  | **Candlestick**   | grafico-candlestick | analisis_tecnico.html         | Análisis técnico     |

---

## 🚀 Cómo Empezar

### Opción 1: Ejecutar Todos los Gráficos

```bash
cd /home/sazardev/Documents/golang/go-data

# Script para ejecutar todos
for dir in grafico-* math-data; do
  echo "═══════════════════════════════════════"
  echo "▶ Ejecutando $dir..."
  echo "═══════════════════════════════════════"
  cd "$dir"
  timeout 5 go run main.go 2>&1 | head -3
  cd ..
done
```

### Opción 2: Ejecutar un Gráfico Específico

```bash
# Ejemplo: Gráfico de barras
cd grafico-barras
go run main.go
# Se abre automáticamente en el navegador

# Ejemplo: Gauge (KPIs)
cd ../grafico-gauge
go run main.go

# Ejemplo: Scatter (Función matemática)
cd ../math-data
go run main.go
```

### Opción 3: Verificar Compilación de Todos

```bash
cd /home/sazardev/Documents/golang/go-data

for dir in grafico-* math-data; do
  cd "$dir"
  echo -n "$dir: "
  go build -o /tmp/test.bin main.go 2>&1 && echo "✓ OK" || echo "✗ ERROR"
  cd ..
done
```

---

## 📂 Estructura del Proyecto

```
go-data/
├── README.md                    ← Este archivo
├── GUIA_RAPIDA.md              ← Inicio rápido
├── CATALOGO_GRAFICOS.md        ← Inventario completo
├── go.mod                      ← Dependencias Go
├── .gitignore                  ← Archivos ignorados
│
├── math-data/
│   ├── main.go
│   ├── README.md
│   └── grafico_x2.html
│
├── grafico-barras/
│   ├── main.go
│   ├── README.md
│   └── comparacion_precios.html
│
├── grafico-pastel/
├── grafico-lineas-multiples/
├── grafico-area/
├── grafico-heatmap/
├── grafico-radar/
├── grafico-histograma/
├── grafico-gauge/
├── grafico-sankey/
├── grafico-funnel/
├── grafico-treemap/
├── grafico-burbuja/
├── grafico-candlestick/
└── grafico-boxplot/
    (Cada uno con: main.go, README.md, archivo.html)
```

---

## 🎯 Guía Rápida de Selección

### ¿Cuál gráfico usar?

**Mostrar composición:**
→ Pie (grafico-pastel) o Bar (grafico-barras)

**Mostrar tendencias temporales:**
→ Line (grafico-lineas-multiples) o Area (grafico-area)

**Comparar valores:**
→ Bar (grafico-barras) o Radar (grafico-radar)

**Mostrar relaciones entre variables:**
→ Scatter (math-data) o Heatmap (grafico-heatmap)

**KPIs o estado actual:**
→ Gauge (grafico-gauge)

**Flujos financieros:**
→ Sankey (grafico-sankey)

**Embudo de conversión:**
→ Funnel (grafico-funnel)

**Análisis técnico (trading):**
→ Candlestick (grafico-candlestick)

**Distribuciones estadísticas:**
→ BoxPlot (grafico-boxplot) o Histogram (grafico-histograma)

**Análisis multidimensional:**
→ Radar (grafico-radar) o Bubble (grafico-burbuja)

---

## 📊 Características Comunes

Todos los gráficos incluyen:

✅ **Interactividad**
- Zoom y pan
- Hover tooltips
- Leyenda interactiva
- Descarga como imagen

✅ **Código Clean**
- Comentarios explicativos
- Estructura modular
- Sigue mejores prácticas Go

✅ **Documentación Completa**
- README en cada carpeta
- Explicación de casos de uso
- Interpretación de datos
- Aplicaciones reales

✅ **Auto-apertura en Navegador**
- Soporta Linux (xdg-open)
- Soporta macOS (open)
- Soporta Windows (cmd)

---

## 🛠️ Tecnología

- **Lenguaje**: Go 1.25.5
- **Librería Principal**: github.com/go-echarts/go-echarts/v2 (v2.6.7)
- **Backend JS**: Apache ECharts (JavaScript)
- **Salida**: HTML interactivo
- **Navegadores**: Chrome, Firefox, Safari, Edge (HTML5)

---

## 📚 Lectura Recomendada

### Para Principiantes
1. Comienza con [GUIA_RAPIDA.md](GUIA_RAPIDA.md)
2. Ejecuta un gráfico simple (ej: grafico-barras)
3. Lee su README.md
4. Modifica el código y experimenta

### Para Expertos
1. Consulta [CATALOGO_GRAFICOS.md](CATALOGO_GRAFICOS.md)
2. Revisa el código fuente en cada carpeta
3. Agrega nuevos gráficos (Tree, Sunburst, Geo, 3D, etc.)
4. Crea dashboards combinando múltiples gráficos

### Para Aprender Go-Echarts
1. Ve a https://github.com/go-echarts/go-echarts
2. Consulta la documentación oficial
3. Adapta ejemplos de este proyecto
4. Crea tus propios casos de uso

---

## 🎓 Casos de Uso por Industria

### 🏢 Empresa / Finanzas
**Mejor para**: Bar, Line, Pie, Gauge, Sankey, Funnel

### 📊 Análisis de Datos
**Mejor para**: Scatter, Heatmap, BoxPlot, Radar, Bubble

### 🛍️ E-commerce
**Mejor para**: Funnel, Bar, Pie, Line

### 💻 Tecnología/IT
**Mejor para**: Line, BoxPlot, Heatmap, Scatter

### 🏥 Medicina/Salud
**Mejor para**: BoxPlot, Heatmap, Scatter

---

## ✅ Validación

**Última Validación: 16 Diciembre 2025**

Todos los 15 proyectos han sido:
- ✅ Compilados exitosamente
- ✅ Ejecutados sin errores
- ✅ Generan HTML interactivo
- ✅ Se abren automáticamente en navegador
- ✅ Documentados completamente

---

## 📞 Recursos Externos

- 🌐 **Go-Echarts GitHub**: https://github.com/go-echarts/go-echarts
- 📖 **Go-Echarts Handbook**: https://go-echarts.github.io/go-echarts/
- 📚 **Go-Echarts GoDoc**: https://pkg.go.dev/github.com/go-echarts/go-echarts/v2
- 🎨 **Apache ECharts**: https://echarts.apache.org/
- 🚀 **Go Official**: https://golang.org

---

## 💡 Tips Profesionales

1. **Elige el gráfico correcto** - Uno por pregunta
2. **Simplifica** - Menos es más en visualización
3. **Contextualiza** - Incluye metas, benchmarks, tendencia
4. **Interactividad** - Aprovecha hover, filtros, leyenda
5. **Storytelling** - Conecta múltiples gráficos en narrativa

---

## 🚀 Próximos Pasos Sugeridos

1. **Explora** los 15 gráficos ejecutándolos
2. **Lee** los README de los que más te interesen
3. **Modifica** datos y vuelve a compilar
4. **Adapta** a tus propios datos/casos
5. **Combina** múltiples gráficos en dashboards
6. **Implementa** gráficos adicionales (Tree, Sunburst, 3D, Geo)

---

## 📝 Resumen

Este proyecto proporciona:
- ✅ 15 ejemplos funcionales y documentados
- ✅ Código Go limpio y reutilizable
- ✅ Documentación completa en Markdown
- ✅ Guías de cuándo usar cada gráfico
- ✅ Aplicaciones reales del mundo

**¡Listo para visualizar tus datos! 🎨📊**

---

**Creado**: Diciembre 2025  
**Go Version**: 1.25.5  
**go-echarts**: v2.6.7  
**Estado**: ✅ Completo y Validado
