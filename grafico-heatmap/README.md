# 📊 Gráfico Heatmap - Matriz de Correlación Financiera

## 📋 Información General

| Propiedad                | Descripción                        |
| ------------------------ | ---------------------------------- |
| **Tipo de Gráfico**      | Heatmap (Mapa de Calor)            |
| **Nombre en go-echarts** | `charts.NewHeatMap()`              |
| **Origen**               | Análisis Estadístico Multivariante |
| **Archivo de Salida**    | `matriz_correlacion.html`          |
| **Datos**                | 6 variables × 6 variables          |

## 🎯 ¿Cuándo Usar?

Los heatmaps son ideales para:

- **Correlación**: Mostrar relaciones entre múltiples variables
- **Densidad**: Visualizar concentración de datos en 2D
- **Matrices**: Representar datos en formato matricial
- **Detección de Patrones**: Identificar agrupaciones y anomalías
- **Análisis Temporal**: Ver patrones que varían en tiempo y espacio

## 🔍 Características

- Usa color para representar valores numéricos
- Colores cálidos (rojo) = valores altos/fuertes
- Colores fríos (azul) = valores bajos/débiles
- Matriz de celdas codificadas por color
- Ideal para grandes matrices de datos

## 📊 Caso de Uso: Matriz de Correlación Financiera

Este ejemplo muestra la **correlación entre 6 variables financieras**:

```
Variables:
  1. Precio          - Valor de mercado del activo
  2. Volumen         - Cantidad de transacciones
  3. ROI             - Retorno sobre inversión
  4. Riesgo          - Volatilidad/desviación estándar
  5. Liquidez        - Facilidad para convertir a efectivo
  6. Rentabilidad    - Ganancia generada

Rango de Correlación: -1 a +1
  -1.0  → Correlación perfecta negativa (inversa)
   0.0  → Sin correlación
  +1.0  → Correlación perfecta positiva
```

## 💡 Aplicaciones Reales

| Sector            | Aplicación                                       |
| ----------------- | ------------------------------------------------ |
| **Finanzas**      | Correlación entre activos, variables económicas  |
| **Biología**      | Genes vs fenotipo, proteína vs función           |
| **Clima**         | Temperatura, humedad, presión por ubicación/hora |
| **Web Analytics** | Evento vs fuente de tráfico                      |
| **Redes**         | Matriz de distancia/conectividad entre nodos     |
| **Manufactura**   | Parámetros de proceso vs calidad                 |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom, tooltip con valor exacto, escala de colores

## 🚀 Ejecución

```bash
cd grafico-heatmap
go run main.go
# Genera: matriz_correlacion.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.HeatMapData{
    {
        Value: []interface{}{i, j, correlacion},
        // i: índice fila (variable Y)
        // j: índice columna (variable X)
        // correlacion: valor de -1 a 1
    },
}
```

## ⚙️ Configuración

- **Dimensión**: 6 × 6 (36 celdas)
- **Eje X y Y**: Mismas 6 variables
- **Rango de Valores**: -1.0 a +1.0
- **Codificación**: Color gradiente

## 📊 Matriz de Datos

```
                Precio  Volumen  ROI    Riesgo Liquidez Rentab
Precio          1.00    0.85     0.92   -0.65  0.78     0.88
Volumen         0.85    1.00     0.71   -0.45  0.82     0.73
ROI             0.92    0.71     1.00   -0.58  0.69     0.95
Riesgo         -0.65   -0.45    -0.58    1.00  -0.62    -0.61
Liquidez        0.78    0.82     0.69   -0.62  1.00     0.70
Rentabilidad    0.88    0.73     0.95   -0.61  0.70     1.00
```

## 🎨 Escala de Colores

```
Valor         Color          Interpretación
+1.0         Rojo Intenso    Correlación perfecta positiva
+0.7 a +0.9  Rojo            Correlación fuerte positiva
+0.3 a +0.7  Naranja         Correlación moderada positiva
 0.0         Blanco/Gris     Sin correlación
-0.3 a -0.7  Cyan            Correlación moderada negativa
-0.7 a -0.9  Azul            Correlación fuerte negativa
-1.0         Azul Intenso    Correlación perfecta negativa
```

## 💡 Interpretación

### Correlaciones Positivas Fuertes (>0.8)
- **Precio ↔ ROI (0.92)**: Precios altos asociados con mejor retorno
- **ROI ↔ Rentabilidad (0.95)**: Muy relacionado, casi la misma métrica
- **Precio ↔ Rentabilidad (0.88)**: Precios altos correlacionan con rentabilidad

### Correlaciones Negativas Fuertes (<-0.6)
- **Precio ↔ Riesgo (-0.65)**: Mayor precio, menor riesgo
- **ROI ↔ Riesgo (-0.58)**: ROI alto asociado con menor volatilidad
- **Rentabilidad ↔ Riesgo (-0.61)**: Rentabilidad alta con bajo riesgo

### Sin Correlación (~0)
- No hay pares con correlación cercana a 0

## 📈 Insights del Ejemplo

1. **Triángulo Superior Izquierda**: Variables de retorno (Precio, ROI, Rentabilidad) altamente correlacionadas
2. **Diagonal Principal**: Siempre 1.0 (auto-correlación perfecta)
3. **Riesgo Aislado**: Tiene correlación opuesta con otras variables
4. **Liquidez**: Variable intermedia, correlaciona moderadamente con todas

## ⚠️ Limitaciones

- Difícil leer valores exactos de grandes matrices
- No muestra causalidad, solo correlación
- Orden de filas/columnas afecta la interpretación visual
- Puede mostrar correlaciones espurias

## 🔧 Mejoras Posibles

- Ordenar matriz por similitud (clustering)
- Agregar significancia estadística (p-values)
- Incluir tabla de valores exactos
- Permitir filtrar por rango de correlación
