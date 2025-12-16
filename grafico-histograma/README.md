# 📊 Gráfico Histograma - Distribución de Tiempos de Respuesta

## 📋 Información General

| Propiedad                | Descripción                              |
| ------------------------ | ---------------------------------------- |
| **Tipo de Gráfico**      | Histogram (Histograma)                   |
| **Nombre en go-echarts** | `charts.NewBar()` (agrupación de rangos) |
| **Origen**               | Estadística Descriptiva                  |
| **Archivo de Salida**    | `histograma_tiempos.html`                |
| **Datos**                | 6 rangos de frecuencias                  |

## 🎯 ¿Cuándo Usar?

Los histogramas son ideales para:

- **Distribución**: Visualizar cómo se distribuyen los datos contínuos
- **Frecuencia**: Mostrar cantidad de observaciones en cada rango
- **Normalidad**: Detectar si los datos siguen distribución normal
- **Asimetría**: Identificar sesgos en los datos
- **Valores Atípicos**: Encontrar outliers fuera de la distribución principal

## 🔍 Características

- Eje X: Rangos o bins de valores
- Eje Y: Frecuencia (cantidad de observaciones)
- Barras contiguas (sin espacios)
- Cada barra representa un intervalo
- Permite ver la forma de la distribución

## 📊 Caso de Uso: Distribución de Tiempos de Respuesta

Este ejemplo muestra cómo se distribuyen los **tiempos de respuesta del servidor**:

```
Rango de Tiempos:
  0-50ms:    450 respuestas  (40%)
  50-100ms:  320 respuestas  (29%)
  100-150ms: 180 respuestas  (16%)
  150-200ms: 95 respuestas   (9%)
  200-250ms: 45 respuestas   (4%)
  250-300ms: 10 respuestas   (1%)

Total: 1,100 respuestas
Moda: 0-50ms (valor más frecuente)
Sesgo: Derecha (cola larga hacia derecha)
```

## 💡 Aplicaciones Reales

| Sector          | Aplicación                                            |
| --------------- | ----------------------------------------------------- |
| **Informática** | Latencia de red, tiempo de carga, tiempo de respuesta |
| **Calidad**     | Distribución de peso, tamaño, duración                |
| **Salud**       | Distribución de presión arterial, peso de pacientes   |
| **Educación**   | Calificaciones de estudiantes, tiempo de examen       |
| **Finanzas**    | Distribución de retornos, montos de transacciones     |
| **Biología**    | Longitud de genes, expresión de proteínas             |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom, filtro por rango, tooltip con estadísticas

## 🚀 Ejecución

```bash
cd grafico-histograma
go run main.go
# Genera: histograma_tiempos.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.BarData{
    {Value: 450},  // Frecuencia del rango
    {Value: 320},
    // ...
}
```

## ⚙️ Configuración

- **Eje X**: 6 rangos de 50ms cada uno
- **Eje Y**: Frecuencia (cantidad)
- **Rango Total**: 0-300ms
- **Bins**: 6 intervalos contiguos
- **Ancho de Barra**: Proporcional al rango

## 📊 Análisis Estadístico

```
Estadísticas de Tiempos de Respuesta:
- Moda (más frecuente):     0-50ms (450)
- Mediana (50%):           ~75ms
- Media Aproximada:        ~90ms
- Máximo:                  300ms
- Mínimo:                  0ms

Percentiles Estimados:
- 10%:                     ~20ms
- 25%:                     ~40ms
- 50%:                     ~75ms
- 75%:                     ~130ms
- 90%:                     ~190ms
```

## 📈 Interpretación

```
Forma de la Distribución:
    │
    │     ┌─┐
    │     │ │
    │     │ │  ┌─┐
    │     │ │  │ │
    │     │ │  │ │  ┌─┐
    │ ┌─┐ │ │  │ │  │ │  ┌─┐ ┌─┐
    ├─┼─┼─┼─┼──┼─┼──┼─┼──┼─┼─┼─┤
    0  1  2  3  4   5  6  7  8  9 (×50ms)

Características:
✓ Moda clara en el inicio
✓ Sesgo hacia la derecha (cola larga)
✓ Mayoría de datos concentrados en 0-50ms
✓ Cola pequeña pero visible en 250-300ms
```

## 🎨 Variantes

### Histograma Normalizado
Mostrar densidad o proporción en lugar de frecuencia

### Histograma Acumulado
Frecuencia acumulada (% acumulativo)

### Curva de Densidad
Superponer curva normal o kernel density estimate

### Histogramas Superpuestos
Comparar dos distribuciones

## 💡 Insights

### Rendimiento del Servidor
- **Excelente**: 40% responden en <50ms
- **Bueno**: 69% responden en <100ms
- **Aceptable**: 85% responden en <150ms
- **Alerta**: 15% tardandan >150ms

### Recomendaciones
1. **Metas de SLA**:
   - 95% < 100ms (¡Lo cumplen!)
   - 99% < 200ms (¡Lo cumplen!)

2. **Optimización**:
   - Investigar las 10 respuestas >250ms
   - Mejorar casos de 200-250ms

## ⚠️ Limitaciones

- Depende mucho del tamaño y número de bins
- Diferente número de bins puede cambiar la interpretación
- No muestra valores individuales exactos
- Puede ocultar bimodalidad si bins no son adecuados

## 🔧 Mejoras Posibles

- Permitir ajustar número de bins
- Superponer normal distribution (campana de Gauss)
- Mostrar media y desviación estándar
- Filtrar outliers manualmente
- Incluir tabla estadística junto al gráfico
