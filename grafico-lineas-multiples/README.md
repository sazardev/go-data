# 📊 Gráfico de Líneas Múltiples - Tendencias de Precios

## 📋 Información General

| Propiedad                | Descripción                      |
| ------------------------ | -------------------------------- |
| **Tipo de Gráfico**      | Line Chart (Gráfico de Líneas)   |
| **Nombre en go-echarts** | `charts.NewLine()`               |
| **Origen**               | Análisis de Series Temporales    |
| **Archivo de Salida**    | `tendencia_precios.html`         |
| **Variante**             | Múltiples series con comparación |

## 🎯 ¿Cuándo Usar?

Los gráficos de líneas son ideales para:

- **Series Temporales**: Mostrar cambios a lo largo del tiempo
- **Tendencias**: Visualizar dirección y patrón de datos
- **Comparación Temporal**: Comparar evolución de múltiples variables
- **Pronóstico**: Identificar tendencias para predicciones
- **Monitoreo**: Seguimiento continuo de métricas

## 🔍 Características

- Puntos conectados por líneas
- Excelente para mostrar tendencias
- Múltiples series para comparación temporal
- Fácil identificar picos y valles
- Ideal para datos contínuos en el tiempo

## 📊 Caso de Uso: Evolución de Precios de Criptomonedas

Este ejemplo muestra la **evolución de 3 criptomonedas** durante **12 meses**:

```
Monedas:
  1. Bitcoin:   $42,000 → $55,000
  2. Ethereum:  $2,500  → $3,500
  3. Solana:    $140    → $260

Período: 12 meses (Ene a Dic)
Escala: USD (diferentes rangos)
Patrones: Tendencias alcistas con fluctuaciones
```

## 💡 Aplicaciones Reales

| Sector       | Aplicación                                      |
| ------------ | ----------------------------------------------- |
| **Finanzas** | Evolución de precios de acciones, criptomonedas |
| **Análisis** | Tendencias de tráfico web, usuarios activos     |
| **Clima**    | Temperatura, precipitación a lo largo del año   |
| **Salud**    | Evolución de síntomas, métricas de salud        |
| **Ventas**   | Ingresos mensuales, conversión en el tiempo     |
| **IoT**      | Datos de sensores en tiempo real                |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom temporal, pan, tooltip, leyenda interactiva

## 🚀 Ejecución

```bash
cd grafico-lineas-multiples
go run main.go
# Genera: tendencia_precios.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.LineData{
    {Value: 42000},  // Valor en el mes
    {Value: 45000},
    // ...
}
```

## ⚙️ Configuración

- **Eje X**: Meses (categoría)
- **Eje Y**: Precio USD (numérico)
- **Número de Series**: 3 (Bitcoin, Ethereum, Solana)
- **Puntos de Datos**: 12 por serie
- **Período**: 12 meses

## 🎨 Variantes

### Línea Suave
Líneas curvadas en lugar de rectas

### Gráfico de Área
Área bajo la línea rellena (efecto acumulativo)

### Línea con Símbolos
Marcadores en cada punto de dato

### Línea Escalonada
Cambios abruptos entre valores

## 📈 Interpretación

- **Tendencia Ascendente**: Mercado en alza, crecimiento positivo
- **Fluctuaciones**: Volatilidad, cambios rápidos de precio
- **Intersecciones**: Puntos donde dos variables tienen igual valor
- **Pendiente**: Velocidad de cambio (mayor pendiente = cambio más rápido)

## 💡 Insight del Ejemplo

```
Bitcoin:    Tendencia alcista consistente (+31% en el período)
Ethereum:   Variabilidad moderada (+40% en el período)
Solana:     Crecimiento alto pero volátil (+86% en el período)
```

## ⚠️ Limitaciones

- Difícil comparar series con escalas muy diferentes
- Muchas series (>5) puede hacer la visualización confusa
- Requiere datos ordenados temporalmente
- Cambios abruptos pueden desdibujar la tendencia

## 🔧 Mejoras Posibles

- Normalizar escalas para mejor comparación
- Agregar promedio móvil para suavizar tendencias
- Incluir bandas de confianza o desviación estándar
