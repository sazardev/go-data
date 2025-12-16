# 📊 Gráfico de Barras - Comparación de Precios

## 📋 Información General

| Propiedad                | Descripción                   |
| ------------------------ | ----------------------------- |
| **Tipo de Gráfico**      | Bar Chart (Gráfico de Barras) |
| **Nombre en go-echarts** | `charts.NewBar()`             |
| **Origen**               | Estadística Descriptiva       |
| **Archivo de Salida**    | `comparacion_precios.html`    |
| **Comparación**          | Múltiples series en grupos    |

## 🎯 ¿Cuándo Usar?

Los gráficos de barras son ideales para:

- **Comparación de Categorías**: Comparar valores entre diferentes categorías
- **Series Múltiples**: Mostrar varias series de datos lado a lado
- **Valores Discretos**: Representar datos categóricos o discretos
- **Análisis Competitivo**: Comparar precios, rendimiento, etc. entre competidores
- **Datos Presupuestarios**: Mostrar asignaciones de presupuesto por departamento

## 🔍 Características

- Barras verticales u horizontales
- Múltiples series pueden compararse lado a lado o apiladas
- Excelente para datos categóricos
- Fácil lectura de valores exactos
- Ideal para auditorios no técnicos

## 📊 Caso de Uso: Comparación de Precios

Este ejemplo compara precios de **5 productos** en **3 tiendas diferentes**:

```
Productos: Laptop, Monitor, Teclado, Mouse, Auriculares
Tiendas: 
  - TechStore
  - ElectroMart
  - CompuWorld

Rango de Precios: $30 a $1,300
Formato: Barras agrupadas por producto
```

## 💡 Aplicaciones Reales

| Sector        | Aplicación                              |
| ------------- | --------------------------------------- |
| **Retail**    | Comparación de precios entre tiendas    |
| **Finanzas**  | Comparación de tasas entre bancos       |
| **Marketing** | Presupuesto por canal publicitario      |
| **RH**        | Salarios por departamento o puesto      |
| **Educación** | Calificaciones por estudiante o materia |
| **Ventas**    | Ingresos por vendedor o región          |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom, filtro de leyenda, descarga de imagen

## 🚀 Ejecución

```bash
cd grafico-barras
go run main.go
# Genera: comparacion_precios.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.BarData{
    {Value: 1200},  // Valor de la barra
    {Value: 350},
    // ...
}
```

## ⚙️ Configuración

- **Categorías (Eje X)**: 5 productos
- **Series**: 3 tiendas
- **Tipo de Eje X**: category
- **Tipo de Eje Y**: value (numérico)
- **Rango de Precios**: $30 a $1,300

## 🎨 Variantes

### Barras Horizontales
Intercambiar ejes para mejorar legibilidad en nombres largos

### Barras Apiladas
Mostrar proporción de cada serie respecto al total

### Barras 100% Apiladas
Mostrar composición relativa (porcentajes)

## 📈 Interpretación

- **Barras del mismo color**: Representan el mismo producto en diferentes tiendas
- **Barras agrupadas**: Facilitan la comparación entre tiendas para cada producto
- **Altura de barras**: Proporcional al precio del producto
