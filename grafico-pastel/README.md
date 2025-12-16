# 📊 Gráfico de Pastel - Distribución de Gastos

## 📋 Información General

| Propiedad                | Descripción                  |
| ------------------------ | ---------------------------- |
| **Tipo de Gráfico**      | Pie Chart (Gráfico Circular) |
| **Nombre en go-echarts** | `charts.NewPie()`            |
| **Origen**               | Estadística Descriptiva      |
| **Archivo de Salida**    | `distribucion_gastos.html`   |
| **Variante**             | Donut Chart (Anillo)         |

## 🎯 ¿Cuándo Usar?

Los gráficos de pastel son ideales para:

- **Composición**: Mostrar cómo se divide un total en partes
- **Proporciones**: Visualizar porcentajes respecto al total
- **Presupuestos**: Desglose de presupuesto por categoría
- **Distribución**: Mostrar distribución de recursos o gastos
- **Marcas de Mercado**: Cuota de mercado entre competidores

## 🔍 Características

- Cada porción representa una categoría
- El tamaño de la porción es proporcional al valor
- Se pueden mostrar porcentajes y valores absolutos
- Ideal para mostrar el "todo" en partes
- Máximo 6-7 categorías recomendado (evita confusión visual)

## 📊 Caso de Uso: Distribución de Gastos Mensuales

Este ejemplo muestra el **desglose de gastos familiares** en 8 categorías:

```
Categorías:
  1. Vivienda:         $1,200 (33.6%)
  2. Alimentación:     $600   (16.8%)
  3. Transporte:       $300   (8.4%)
  4. Servicios:        $250   (7.0%)
  5. Entretenimiento:  $200   (5.6%)
  6. Educación:        $400   (11.2%)
  7. Ahorros:          $500   (14.0%)
  8. Otros:            $150   (4.2%)

Total: $3,600
```

## 💡 Aplicaciones Reales

| Sector                | Aplicación                            |
| --------------------- | ------------------------------------- |
| **Finanzas Personal** | Desglose de gastos mensuales          |
| **Empresa**           | Ingresos por línea de negocio         |
| **Ventas**            | Cuota de mercado por competidor       |
| **Presupuesto**       | Asignación de fondos por departamento |
| **Tiempo**            | Cómo se distribuye el tiempo laboral  |
| **Población**         | Distribución demográfica              |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Seleccionar porciones, tooltip con valores, leyenda clickeable

## 🚀 Ejecución

```bash
cd grafico-pastel
go run main.go
# Genera: distribucion_gastos.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.PieData{
    {Name: "Categoría", Value: 1200},
    {Name: "Otra", Value: 600},
    // ...
}
```

## ⚙️ Configuración

- **Radio Interior**: 30% (para efecto donut)
- **Radio Exterior**: 75%
- **Número de Categorías**: 8
- **Total**: $3,600
- **Leyenda**: Posicionada a la izquierda

## 🎨 Variantes

### Gráfico de Pastel Simple
Sin el anillo interior (radio interior = 0%)

### Donut Chart
Con radio interior > 0% (este ejemplo)

### Pastel Anidado
Múltiples niveles de categorización

### Pastel Explosivo
Algunas porciones destacadas (separadas)

## 📈 Interpretación

- **Porciones Grandes**: Gastos principales (Vivienda, Educación)
- **Porciones Pequeñas**: Gastos menores (Otros, Servicios)
- **Leyenda Lateral**: Facilita la referencia de categorías
- **Colores**: Codificación automática para diferenciación

## ⚠️ Limitaciones

- Difícil comparar valores similares
- No recomendado para más de 7-8 categorías
- Los valores pequeños son difíciles de leer
- Mejor para mostrar "composición" que "comparación"
