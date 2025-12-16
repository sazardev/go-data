# 📊 Gráfico de Área - Ventas Acumulativas por Región

## 📋 Información General

| Propiedad                | Descripción                                  |
| ------------------------ | -------------------------------------------- |
| **Tipo de Gráfico**      | Area Chart (Gráfico de Área)                 |
| **Nombre en go-echarts** | `charts.NewLine()` con configuración de área |
| **Origen**               | Análisis de Series Temporales                |
| **Archivo de Salida**    | `ventas_regiones.html`                       |
| **Variante**             | Múltiples áreas apiladas                     |

## 🎯 ¿Cuándo Usar?

Los gráficos de área son ideales para:

- **Acumulación**: Mostrar cómo se acumula un valor total en el tiempo
- **Composición Temporal**: Visualizar contribución de cada serie al total
- **Cambio de Proporción**: Cómo cambia la importancia relativa de cada serie
- **Volumen Total**: Enfatizar magnitud además de tendencia
- **Comparación Apilada**: Múltiples series que suman un total

## 🔍 Características

- Área debajo de la línea rellena con color
- Múltiples áreas apiladas para mostrar total
- Excelente para visualizar proporciones cambiantes
- Enfatiza magnitud además de tendencia
- Cada color representa una serie/categoría

## 📊 Caso de Uso: Ventas Acumulativas por Región

Este ejemplo muestra **ventas trimestrales** de **4 regiones geográficas** durante el año:

```
Regiones:
  1. América:  $150K → $580K (Q1 → Q4)
  2. Europa:   $120K → $450K
  3. Asia:     $200K → $750K (Mayor crecimiento)
  4. África:   $50K  → $180K (Menor mercado)

Período: 4 trimestres (Q1, Q2, Q3, Q4)
Total Inicial (Q1): $520K
Total Final (Q4):   $1,960K (Crecimiento 277%)
```

## 💡 Aplicaciones Reales

| Sector         | Aplicación                                         |
| -------------- | -------------------------------------------------- |
| **Ventas**     | Ingresos acumulativos por región/país              |
| **Producción** | Volumen de producción por línea en el tiempo       |
| **Tráfico**    | Usuarios por fuente de tráfico (web, mobile, api)  |
| **Energía**    | Consumo energético por tipo (solar, eólica, fósil) |
| **Población**  | Crecimiento demográfico por grupo etario           |
| **Finanzas**   | Composición de portafolio a lo largo del tiempo    |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom temporal, seleccionar series, tooltip con desglose

## 🚀 Ejecución

```bash
cd grafico-area
go run main.go
# Genera: ventas_regiones.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.LineData{
    {Value: 150000},  // Valores acumulativos (se grafican como área)
    {Value: 280000},
    // ...
}
```

## ⚙️ Configuración

- **Eje X**: Trimestres (Q1, Q2, Q3, Q4)
- **Eje Y**: Ventas en USD
- **Número de Series**: 4 regiones
- **Puntos de Datos**: 4 por serie
- **Apilamiento**: Sí (áreas se superponen)

## 🎨 Variantes

### Área Apilada al 100%
Muestra proporción relativa (porcentajes) en lugar de valores absolutos

### Área No Apilada
Cada área transparente para ver sobreposiciones

### Área Suave
Bordes curvos en lugar de rectos

### Área Escalonada
Cambios abruptos entre trimestres

## 📈 Interpretación

```
Área de Color:
  - Tamaño: Volumen de ventas de esa región
  - Pendiente: Velocidad de crecimiento
  - Espesor Total: Total de ventas en ese período

Asia (rojo/naranja):
  - Área más grande: Mayor mercado
  - Crecimiento constante: Mercado estable

África (color más claro):
  - Área muy pequeña: Mercado emergente
  - Potencial de crecimiento: Curva menos pronunciada
```

## 💡 Insights del Ejemplo

```
Trimestre 1 (Q1):  Total $520K
  - América: 29%
  - Europa:  23%
  - Asia:    38% (Líder)
  - África:  10%

Trimestre 4 (Q4):  Total $1,960K (Crecimiento 277%)
  - América: 30% (Estable)
  - Europa:  23% (Estable)
  - Asia:    38% (Mantiene liderazgo)
  - África:  9%  (Ligeramente menor%)
```

## 📊 Comparación de Formatos

| Aspecto          | Línea             | Área                   | Comparación               |
| ---------------- | ----------------- | ---------------------- | ------------------------- |
| Énfasis          | Tendencia         | Magnitud y composición | Area > Línea para totales |
| Múltiples Series | Bien              | Mejor (se ve el stack) | Area mejora visualización |
| Valores Exactos  | Fácil leer puntos | Menos exacto           | Línea > Área              |
| Impacto Visual   | Moderado          | Alto                   | Area más impactante       |

## ⚠️ Limitaciones

- Difícil leer valores exactos (principalmente para series inferiores)
- No recomendado para más de 5-6 series
- Áreas pequeñas pueden no ser visibles
- Requiere orden de series consistente

## 🔧 Mejoras Posibles

- Permitir seleccionar serie específica para enfoque
- Mostrar tabla de datos junto al gráfico
- Incluir línea de tendencia general
