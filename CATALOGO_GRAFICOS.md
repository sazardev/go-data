# 📊 Catálogo Completo de Gráficos - Go-Echarts v2.6.7

## 🎯 Resumen

**go-echarts** (Apache ECharts para Go) soporta **25+ tipos de gráficos** diferentes. En este proyecto hemos implementado **15 de los más populares y útiles para negocio, análisis de datos y finanzas**.

---

## ✅ GRÁFICOS IMPLEMENTADOS (15)

| #   | Tipo          | Nombre en go-echarts        | Carpeta                  | Caso de Uso                 | Estado |
| --- | ------------- | --------------------------- | ------------------------ | --------------------------- | ------ |
| 1   | Scatter       | `charts.NewScatter()`       | math-data                | Función matemática f(x)=x²  | ✅      |
| 2   | Bar           | `charts.NewBar()`           | grafico-barras           | Comparación de precios      | ✅      |
| 3   | Pie           | `charts.NewPie()`           | grafico-pastel           | Distribución de gastos      | ✅      |
| 4   | Line          | `charts.NewLine()`          | grafico-lineas-multiples | Tendencias de precios       | ✅      |
| 5   | Area          | `charts.NewLine()` + config | grafico-area             | Ventas acumulativas         | ✅      |
| 6   | Heatmap       | `charts.NewHeatMap()`       | grafico-heatmap          | Matriz de correlación       | ✅      |
| 7   | Radar         | `charts.NewRadar()`         | grafico-radar            | Desempeño multidimensional  | ✅      |
| 8   | Histogram     | `charts.NewBar()`           | grafico-histograma       | Distribución de frecuencias | ✅      |
| 9   | Gauge         | `charts.NewGauge()`         | grafico-gauge            | KPIs empresariales          | ✅      |
| 10  | Sankey        | `charts.NewSankey()`        | grafico-sankey           | Flujo financiero            | ✅      |
| 11  | Funnel        | `charts.NewFunnel()`        | grafico-funnel           | Embudo de conversión        | ✅      |
| 12  | Bar (Treemap) | `charts.NewBar()`           | grafico-treemap          | Distribución presupuestaria | ✅      |
| 13  | Bubble        | `charts.NewEffectScatter()` | grafico-burbuja          | Análisis de productos 3D    | ✅      |
| 14  | Candlestick   | `charts.NewKLine()`         | grafico-candlestick      | Análisis técnico financiero | ✅      |
| 15  | BoxPlot       | `charts.NewBoxPlot()`       | grafico-boxplot          | Distribución estadística    | ✅      |

---

## ⭕ GRÁFICOS NO IMPLEMENTADOS (10+)

Estos son los tipos disponibles en go-echarts pero aún no creados:

### 📍 Gráficos Geográficos

| Tipo            | Nombre                               | Caso de Uso                                                  | Complejidad |
| --------------- | ------------------------------------ | ------------------------------------------------------------ | ----------- |
| **Geo/Map**     | `charts.NewGeo()`, `charts.NewMap()` | Visualizar datos por ubicación geográfica (países, ciudades) | MEDIA       |
| **Leaflet Map** | Con integración OpenStreetMap        | Mapas interactivos con marcadores                            | ALTA        |

### 🔗 Gráficos de Red/Relaciones

| Tipo                | Nombre               | Caso de Uso                               | Complejidad |
| ------------------- | -------------------- | ----------------------------------------- | ----------- |
| **Graph**           | `charts.NewGraph()`  | Redes de conexión, relaciones entre nodos | MEDIA       |
| **Network Diagram** | Relaciones complejas | Org charts, dependencias, conexiones      | ALTA        |

### 🌳 Gráficos Jerárquicos

| Tipo               | Nombre                 | Caso de Uso                              | Complejidad |
| ------------------ | ---------------------- | ---------------------------------------- | ----------- |
| **Tree**           | `charts.NewTree()`     | Estructura jerárquica, árbol genealógico | MEDIA       |
| **Sunburst**       | `charts.NewSunburst()` | Jerarquía radial con proporciones        | MEDIA       |
| **Treemap (real)** | `charts.NewTreeMap()`  | Distribución jerárquica de espacios      | MEDIA       |

### 📊 Gráficos 3D

| Tipo          | Nombre                  | Caso de Uso                            | Complejidad |
| ------------- | ----------------------- | -------------------------------------- | ----------- |
| **Bar3D**     | `charts.NewBar3D()`     | Barras en 3D, volumen de datos 3D      | MEDIA       |
| **Scatter3D** | `charts.NewScatter3D()` | Scatter con 3ª dimensión (profundidad) | MEDIA       |
| **Surface3D** | `charts.NewSurface3D()` | Superficies 3D, terreno, ondas         | ALTA        |
| **Line3D**    | `charts.NewLine3D()`    | Líneas en espacio 3D                   | MEDIA       |

### 🌊 Gráficos Especializados

| Tipo              | Nombre                      | Caso de Uso                              | Complejidad |
| ----------------- | --------------------------- | ---------------------------------------- | ----------- |
| **Liquid**        | `charts.NewLiquid()`        | Gauge líquido (medidor tipo agua)        | MEDIA       |
| **WordCloud**     | `charts.NewWordCloud()`     | Nube de palabras por frecuencia          | BAJA        |
| **ThemeRiver**    | `charts.NewThemeRiver()`    | Flujo temporal con múltiples categorías  | MEDIA       |
| **Parallel**      | `charts.NewParallel()`      | Coordenadas paralelas multidimensionales | ALTA        |
| **EffectScatter** | `charts.NewEffectScatter()` | Scatter con efecto visual (onda)         | MEDIA       |

---

## 🎓 Clasificación por Caso de Uso

### 📈 Análisis de Tendencias Temporales
- ✅ **Line** - Mejor opción
- ✅ **Area** - Mostrar acumulación
- ⭕ **ThemeRiver** - Tendencias complejas

### 💼 Análisis Empresarial
- ✅ **Bar** - Comparación directa
- ✅ **Pie** - Composición
- ✅ **Sankey** - Flujos
- ✅ **Gauge** - KPIs
- ✅ **Funnel** - Conversión

### 🔬 Análisis Estadístico
- ✅ **Scatter** - Relaciones
- ✅ **Heatmap** - Correlaciones
- ✅ **BoxPlot** - Distribuciones
- ✅ **Histogram** - Frecuencias
- ⭕ **Parallel Coordinates** - Multivariante

### 💰 Análisis Financiero
- ✅ **Candlestick** - Trading
- ✅ **Line** - Series de tiempo
- ⭕ **Bar3D** - Volumen múltiple

### 🎯 Análisis Multidimensional
- ✅ **Radar** - Perfil de 5-7 dimensiones
- ✅ **Bubble** - 3 dimensiones
- ⭕ **Scatter3D** - 3 dimensiones con profundidad
- ⭕ **Parallel** - 4+ dimensiones

### 🗺️ Análisis Geográfico
- ⭕ **Geo/Map** - Por región/país
- ⭕ **Leaflet** - Ubicaciones precisas

### 📊 Información Jerárquica
- ✅ **Sankey** - Flujos
- ⭕ **Tree** - Estructura pura
- ⭕ **Sunburst** - Jerarquía radial

### 🎨 Visualización de Texto
- ⭕ **WordCloud** - Palabras frecuentes

---

## 📊 Matriz de Características

```
CARACTERÍSTICA          BAR  LINE SCATTER GAUGE RADAR SANKEY HEATMAP CANDL BUBBLE FUNNEL BOX
────────────────────────────────────────────────────────────────────────────────────────
Múltiples Series        ✓    ✓    ✓       ✓     ✓     ✓      ✓       ✓     ✓      ✗   ✓
Series Temporal         ✗    ✓    ✗       ✗     ✗     ✗      ✓       ✓     ✗      ✗   ✗
Multidimensional        ✗    ✗    ✓ (2)   ✗     ✓ (5) ✗      ✗       ✗     ✓ (3) ✗   ✗
Correlación             ✗    ✗    ✓       ✗     ✗     ✗      ✓       ✗     ✗      ✗   ✗
Distribución            ✗    ✗    ✓       ✗     ✗     ✗      ✗       ✗     ✗      ✗   ✓
Jerarquía               ✗    ✗    ✗       ✗     ✗     ✓      ✗       ✗     ✗      ✓   ✗
KPI/Medidor             ✗    ✗    ✗       ✓     ✗     ✗      ✗       ✗     ✗      ✗   ✗
Flujos                  ✗    ✗    ✗       ✗     ✗     ✓      ✗       ✗     ✗      ✗   ✗
```

---

## 🚀 Recomendaciones para Nuevos Gráficos

### Prioridad ALTA (Fáciles + Útiles)
1. **WordCloud** - Análisis de palabras clave, sentimientos
2. **ThemeRiver** - Evolución de categorías en tiempo
3. **Liquid** - Medidores visuales alternativos

### Prioridad MEDIA (Medianos + Útiles)
1. **Tree/Sunburst** - Org charts, estructuras jerárquicas
2. **Geo/Map** - Análisis por región/país
3. **Scatter3D** - Datos con 3 dimensiones reales

### Prioridad BAJA (Complejos)
1. **Bar3D/Surface3D** - Visualización 3D avanzada
2. **Parallel Coordinates** - Análisis multivariante complejo
3. **Graph/Network** - Redes complejas

---

## 💾 Estructura del Proyecto

```
/home/sazardev/Documents/golang/go-data/
├── go.mod
├── .gitignore
├── math-data/                    ✅ Scatter
│   ├── main.go
│   ├── README.md
│   └── grafico_x2.html
├── grafico-barras/               ✅ Bar
│   ├── main.go
│   ├── README.md
│   └── comparacion_precios.html
├── grafico-pastel/               ✅ Pie
├── grafico-lineas-multiples/     ✅ Line
├── grafico-area/                 ✅ Area
├── grafico-heatmap/              ✅ Heatmap
├── grafico-radar/                ✅ Radar
├── grafico-histograma/           ✅ Histogram (Bar)
├── grafico-gauge/                ✅ Gauge
├── grafico-sankey/               ✅ Sankey
├── grafico-funnel/               ✅ Funnel
├── grafico-treemap/              ✅ Bar (como treemap)
├── grafico-burbuja/              ✅ Bubble
├── grafico-candlestick/          ✅ Candlestick
└── grafico-boxplot/              ✅ BoxPlot
```

---

## 🛠️ Cómo Ejecutar Cualquier Gráfico

```bash
cd /home/sazardev/Documents/golang/go-data/[carpeta]
go run main.go
# Se generará un archivo HTML en la misma carpeta
# Se abrirá automáticamente en el navegador
```

---

## 📚 Documentación Oficial

- **GitHub**: https://github.com/go-echarts/go-echarts
- **Handbook**: https://go-echarts.github.io/go-echarts/
- **GoDoc**: https://pkg.go.dev/github.com/go-echarts/go-echarts/v2

---

## 🎯 Casos de Uso Recomendados por Industria

### 🏢 Empresa / Finanzas
✅ **Recomendado**: Bar, Line, Pie, Gauge, Sankey, Funnel
⭕ **Considerar**: Heatmap, BoxPlot, Candlestick

### 📊 Análisis de Datos / Data Science
✅ **Recomendado**: Scatter, Heatmap, BoxPlot, Radar, Bubble
⭕ **Considerar**: 3D, Parallel, Sunburst

### 🛍️ E-commerce / Marketing
✅ **Recomendado**: Funnel, Bar, Pie, Line, Gauge
⭕ **Considerar**: Geo (por región), WordCloud

### 🏥 Medicina / Salud
✅ **Recomendado**: BoxPlot, Heatmap, Scatter, Histogram
⭕ **Considerar**: 3D, Parallel

### 🗺️ Geografía / Demografía
✅ **Recomendado**: Geo/Map, Scatter
⭕ **Considerar**: Bubble, Heatmap

---

## 📋 Checklist de Características

Cada gráfico implementado incluye:
- ✅ Código Go limpio y comentado
- ✅ Auto-apertura en navegador (cross-platform)
- ✅ HTML interactivo con zoom, pan, tooltip
- ✅ Archivo README.md detallado con:
  - Cuándo usar
  - Interpretación de datos
  - Casos de uso reales
  - Análisis e insights
  - Recomendaciones

---

## 🎓 Para Aprender Más

Cada carpeta contiene:
1. **main.go** - Código fuente educativo
2. **README.md** - Guía completa de uso
3. **HTML generado** - Gráfico interactivo

Explora los archivos README para:
- Entender cuándo usar cada gráfico
- Ver interpretación de datos reales
- Aplicar a tus casos de uso específicos

---

## 📞 Soporte

Para agregar nuevos gráficos o modificar existentes:
1. Consulta la documentación oficial de go-echarts
2. Revisa ejemplos en el repositorio oficial
3. Adapta el patrón usado en otros gráficos

---

**Última Actualización**: Diciembre 2025  
**Versión go-echarts**: v2.6.7  
**Go Version**: 1.25.5
