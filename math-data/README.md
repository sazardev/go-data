# 📊 Gráfico Scatter - Función Matemática f(x) = x²

## 📋 Información General

| Propiedad                | Descripción                          |
| ------------------------ | ------------------------------------ |
| **Tipo de Gráfico**      | Scatter Plot (Gráfico de Dispersión) |
| **Nombre en go-echarts** | `charts.NewScatter()`                |
| **Origen**               | Matemáticas y Análisis de Datos      |
| **Archivo de Salida**    | `grafico_x2.html`                    |

## 🎯 ¿Cuándo Usar?

Los gráficos Scatter son ideales para:

- **Análisis de Relaciones**: Mostrar la relación entre dos variables numéricas
- **Distribuciones**: Visualizar cómo se distribuyen los datos en un espacio 2D o 3D
- **Funciones Matemáticas**: Representar funciones y ecuaciones como f(x)
- **Detección de Patrones**: Identificar clusters, tendencias y outliers
- **Correlación**: Analizar si existe correlación entre dos variables

## 🔍 Características

- Cada punto representa un par de valores (x, y)
- Se puede mostrar la relación entre dos variables contínuas
- Ideal para grandes volúmenes de datos puntuales
- Permite identificar patrones, clusters y outliers
- Compatible con análisis de regresión

## 📊 Caso de Uso: f(x) = x²

Este ejemplo muestra una **parábola cuadrática** (función matemática) mediante puntos en el plano:

```
Eje X: Valores de x desde -10 hasta 10
Eje Y: Valores de f(x) = x²

Rango de Y: 0 a 100
Forma: Parábola descendente-ascendente (forma de U)
```

## 💡 Aplicaciones Reales

| Sector        | Aplicación                                           |
| ------------- | ---------------------------------------------------- |
| **Ciencia**   | Análisis experimental, relaciones físicas            |
| **Medicina**  | Correlación dosis-respuesta, BMI vs presión arterial |
| **Negocio**   | Análisis de inversión vs retorno, precio vs demanda  |
| **Finanzas**  | Riesgo vs rendimiento de activos                     |
| **Marketing** | Gasto publicitario vs conversiones                   |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom, pan, tooltip, descarga de imagen

## 🚀 Ejecución

```bash
cd math-data
go run main.go
# Genera: grafico_x2.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.ScatterData{
    {Value: []interface{}{x, y}},
    // x: valor en el eje horizontal
    // y: valor en el eje vertical (resultado de la función)
}
```

## ⚙️ Configuración

- **Rango X**: -10 a 10 (incremento de 0.1)
- **Cálculo Y**: y = x * x
- **Total de Puntos**: 201 puntos
- **Eje X**: Tipo "value" (numérico)
- **Eje Y**: Tipo "value" (numérico)
