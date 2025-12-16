# 📊 Gráfico BoxPlot - Distribución Estadística

## 📋 Información General

| Propiedad                | Descripción                          |
| ------------------------ | ------------------------------------ |
| **Tipo de Gráfico**      | Box Plot (Diagrama de Caja)          |
| **Nombre en go-echarts** | `charts.NewBoxPlot()`                |
| **Origen**               | Estadística Descriptiva (John Tukey) |
| **Archivo de Salida**    | `boxplot_estadistico.html`           |
| **Series**               | 4 servidores con distribuciones      |

## 🎯 ¿Cuándo Usar?

Los box plots son ideales para:

- **Distribución**: Mostrar resumen estadístico en forma visual
- **Comparación**: Comparar distribuciones entre grupos
- **Outliers**: Identificar valores atípicos
- **Cuartiles**: Ver cómo se distribuyen los datos en rangos
- **Simetría**: Detectar si la distribución es simétrica

## 🔍 Características

- Línea central = Mediana
- Caja = Rango intercuartil (25%-75%)
- Líneas (whiskers) = Mín y máx (dentro de 1.5×IQR)
- Puntos = Outliers
- Compacto pero informativo

## 📊 Caso de Uso: Rendimiento de Servidores

Este ejemplo compara la **distribución de tiempos de respuesta** de **4 servidores**:

```
TIEMPOS DE RESPUESTA (ms):
Servidor 1: [25, 45, 75, 10, 95]
Servidor 2: [30, 50, 80, 15, 100]
Servidor 3: [20, 40, 70, 8, 90]
Servidor 4: [35, 55, 85, 20, 105]

ESTADÍSTICAS POR SERVIDOR:
               Min    Q1   Mediana   Q3   Max    IQR
Servidor 1     10     35     50      80    95    45
Servidor 2     15     40     55      85   100    45
Servidor 3      8     30     45      75    90    45
Servidor 4     20     45     60      90   105    45
```

## 💡 Aplicaciones Reales

| Sector        | Aplicación                                 |
| ------------- | ------------------------------------------ |
| **IT**        | Latencia, tiempo de respuesta por servidor |
| **Calidad**   | Distribución de medidas en manufactura     |
| **Medicina**  | Distribución de presión arterial, peso     |
| **Educación** | Calificaciones por grupo o materia         |
| **Economía**  | Distribución de ingresos por región        |
| **Logística** | Tiempo de entrega por ruta                 |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Hover detallado, tooltip con estadísticas

## 🚀 Ejecución

```bash
cd grafico-boxplot
go run main.go
# Genera: boxplot_estadistico.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.BoxPlotData{
    {Value: []interface{}{25, 45, 75, 10, 95}},
    // Valores: Min, Q1, Mediana, Max, IQR
}
```

## ⚙️ Configuración

- **Eje X**: 4 servidores
- **Eje Y**: Tiempo en milisegundos (0-105 ms)
- **Series**: 4 cajas (una por servidor)
- **Datos**: 5 valores estadísticos por caja

## 📊 Anatomía del Box Plot

```
                    Máximo (95ms)
                         ●
                         │
                    ╭────┴────╮  ← Whisker (bigote)
                    │          │
          ╭─────────┤    ●     ├────────╮  ← Cuartiles
          │    ┌────┤  75      ├────┐   │
  Q1 (25) │    │ ┌──┤ Mediana ├──┐ │   │  Q3 (75)
          │    │ │  │   50    │  │ │   │
          │    │ │  └──────────  │ │   │
          │    └────┤           ├────┘   │
          ╰────────┤     ●      ├────────╯
                    │          │
                    │ Outlier? │ ← Si está fuera del rango
                    │          │
                    ╰────┬─────╯
                         │
                    Mínimo (10ms)
```

## 🎓 Componentes Estadísticos

### Mediana (Q2 o 50º percentil)
```
Definición: Valor central cuando datos están ordenados
Servidor 1: 50ms (mitad arriba, mitad abajo)
Interpretación: Rendimiento "típico"
```

### Cuartiles
```
Q1 (25º percentil): 25% de datos bajo este valor
- Servidor 1: 25ms (25% de respuestas < 25ms)

Q3 (75º percentil): 75% de datos bajo este valor
- Servidor 1: 75ms (75% de respuestas < 75ms)

IQR (Rango Intercuartil) = Q3 - Q1
- Servidor 1: 75 - 25 = 50ms (50% de datos central)
```

### Whiskers (Bigotes)
```
Límite Inferior = Q1 - 1.5 × IQR
- Servidor 1: 25 - (1.5 × 50) = -50 (mínimo es 10, ok)

Límite Superior = Q3 + 1.5 × IQR
- Servidor 1: 75 + (1.5 × 50) = 150 (máximo es 95, ok)

Rango: 10 a 95 (sin outliers)
```

### Outliers
```
Cualquier valor:
- Menor a Q1 - 1.5×IQR = Outlier bajo
- Mayor a Q3 + 1.5×IQR = Outlier alto

En el ejemplo: NO hay outliers
```

## 📈 Análisis Comparativo

### Servidor 1: Rendimiento Intermedio
```
┌─────────────────────────────┐
│ 10  35  50  80  95          │
└─────────────────────────────┘
Rango: 10-95 (85ms span)
Mediana: 50ms
Dispersión: Moderada
Simetría: Relativamente simétrico
Evaluación: ACEPTABLE
```

### Servidor 2: Rendimiento Más Lento
```
┌─────────────────────────────┐
│ 15  40  55  85 100          │
└─────────────────────────────┘
Rango: 15-100 (85ms span)
Mediana: 55ms (+5ms vs S1)
Dispersión: Moderada
Simetría: Relativamente simétrico
Evaluación: LENTO (15-20ms más lento)
```

### Servidor 3: Rendimiento MEJOR
```
┌─────────────────────────────┐
│ 8   30  45  75  90          │
└─────────────────────────────┘
Rango: 8-90 (82ms span)
Mediana: 45ms (-5ms vs S1)
Dispersión: Moderada
Simetría: Relativamente simétrico
Evaluación: EXCELENTE (más rápido)
```

### Servidor 4: Rendimiento Más Lento
```
┌─────────────────────────────┐
│ 20  45  60  90 105          │
└─────────────────────────────┘
Rango: 20-105 (85ms span)
Mediana: 60ms (+10ms vs S1)
Dispersión: Moderada
Simetría: Relativamente simétrico
Evaluación: LENTO (peor rendimiento)
```

## 🏆 Ranking de Servidores

```
Posición    Servidor    Mediana    Rango    Evaluación
─────────────────────────────────────────────────────
1º lugar    S3          45ms       8-90     ⭐⭐⭐ EXCELENTE
2º lugar    S1          50ms       10-95    ⭐⭐ BUENO
3º lugar    S2          55ms       15-100   ⭐ ACEPTABLE
4º lugar    S4          60ms       20-105   ✗ LENTO
```

## 💡 Interpretación de Simetría

### Box Plot Simétrico
```
        ║
    ╭───╫───╮
    │   ║   │  ← Mediana en el centro
    ╰───╫───╯
        ║
```
Significado: Distribución normal/equilibrada

### Box Plot Sesgado a Derecha
```
        ║
    ╭───╫──────╮
    │   ║      │  ← Mediana a la izquierda
    ╰───╫──────╯
        ║
```
Significado: Hay valores altos pero raros (outliers)

### Box Plot Sesgado a Izquierda
```
        ║
    ╭──────╫───╮
    │      ║   │  ← Mediana a la derecha
    ╭──────╫───╯
        ║
```
Significado: Hay valores bajos pero raros

## 📊 Análisis de Variabilidad

```
Servidor    IQR (25%-75%)    Variabilidad    Consistencia
──────────────────────────────────────────────────────────
S1          50ms             Moderada        Media
S2          45ms             Moderada        Media-Alta
S3          45ms             Moderada        Media-Alta
S4          45ms             Moderada        Media

IQR pequeño = Más consistente
IQR grande = Más variable

→ Todos tienen variabilidad similar
→ Diferencia es en la mediana (ubicación)
```

## 🔧 Decisiones Operacionales

### Recomendación 1: Usar S3 como Referencia
```
✓ Mejor rendimiento promedio
✓ Rango aceptable
→ Estudiar configuración de S3
→ Aplicar a otros servidores
```

### Recomendación 2: Investigar S4
```
⚠️ Peor rendimiento
⚠️ Máximo alcanza 105ms (inaceptable)
→ Diagnóstico: ¿Hardware viejo? ¿Carga alta?
→ Acción: Actualizar hardware o redistribuir carga
```

### Recomendación 3: Meta SLA
```
Establecer SLA (Service Level Agreement):
- 50% de respuestas < 45ms (S3)
- 95% de respuestas < 80ms (Q3)
- <1% sobre 100ms (no outliers graves)

Cumplimiento Actual: ✓ Se cumplen
```

## ⚠️ Limitaciones

- No muestra forma exacta de distribución
- Necesita muchos datos para ser confiable
- Difficult interpretar con pocos datos
- No muestra patrón temporal

## 🔧 Mejoras Posibles

- Agregar violín plot (muestra distribución)
- Superponer puntos individuales
- Series temporales de box plots
- Benchmarks de industria como línea
- Alertas si mediana excede umbral

## 📝 Conclusión

**Resumen del Análisis**:
- **Servidor 3**: Mejor rendimiento (45ms mediana)
- **Servidor 4**: Requiere atención (60ms mediana)
- **Variabilidad**: Similar en todos (IQR ~45ms)
- **Acción**: Migrar carga de S4 a S3, investigar causa

**SLA Actual**: ✓ Se cumplen especificaciones
**Mejora Potencial**: +15% en rendimiento si se optimiza S4
