# 📊 Gráfico Radar - Análisis de Desempeño de Equipos

## 📋 Información General

| Propiedad                | Descripción                    |
| ------------------------ | ------------------------------ |
| **Tipo de Gráfico**      | Radar Chart (Gráfico de Araña) |
| **Nombre en go-echarts** | `charts.NewRadar()`            |
| **Origen**               | Análisis Multidimensional      |
| **Archivo de Salida**    | `desempenio_equipos.html`      |
| **Dimensiones**          | 5 indicadores por equipo       |

## 🎯 ¿Cuándo Usar?

Los gráficos radar son ideales para:

- **Evaluación Multidimensional**: Comparar múltiples atributos simultáneamente
- **Perfiles**: Mostrar fortalezas y debilidades de forma visual
- **Competencia**: Comparar desempeño de equipos/productos en varias dimensiones
- **Balance**: Identificar áreas desequilibradas que necesitan atención
- **Simetría**: Detectar qué es simétrico vs qué necesita mejora

## 🔍 Características

- Múltiples ejes radiantes desde el centro
- Cada eje representa una dimensión/métrica
- Polígono conecta todos los puntos de un conjunto
- Múltiples polígonos pueden superponerse
- Ideal para 3-5 dimensiones (máximo 7-8)

## 📊 Caso de Uso: Desempeño de Equipos

Este ejemplo compara el **desempeño de 3 equipos** en **5 dimensiones**:

```
Indicadores (escala 0-100):
  1. Productividad    - Velocidad de entrega
  2. Calidad          - Precisión y cero defectos
  3. Innovación       - Generación de ideas nuevas
  4. Comunicación     - Colaboración interna
  5. Cumplimiento     - Adhesión a procesos

Equipos:
  - Equipo A: Líder en Productividad (92) e Innovación (78)
  - Equipo B: Equilibrado, fuerte en Calidad (92)
  - Equipo C: Fuerte en Innovación (92) y Comunicación (90)
```

## 💡 Aplicaciones Reales

| Sector        | Aplicación                                            |
| ------------- | ----------------------------------------------------- |
| **RRHH**      | Evaluación de competencias de empleados               |
| **Deporte**   | Análisis de habilidades de jugadores                  |
| **Producto**  | Comparación de características entre modelos          |
| **Finanzas**  | Evaluación de riesgo (riesgo país, rating crediticio) |
| **Educación** | Desempeño estudiantil en diferentes materias          |
| **Empresa**   | Salud organizacional (cultura, productividad, etc.)   |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom, seleccionar serie, tooltip detallado

## 🚀 Ejecución

```bash
cd grafico-radar
go run main.go
# Genera: desempenio_equipos.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
radarIndicators := []*opts.Indicator{
    {Name: "Productividad", Max: 100},
    {Name: "Calidad", Max: 100},
    // ...
}

[]opts.RadarData{
    {Value: []interface{}{92, 85, 78, 88, 95}},
}
```

## ⚙️ Configuración

- **Indicadores**: 5 dimensiones
- **Escala**: 0-100 por indicador
- **Series**: 3 equipos
- **Forma**: Pentágono (5 lados)
- **Interpolación**: Líneas conectan puntos

## 📊 Datos del Ejemplo

```
                 Prod  Calid  Innov  Comun  Cumpl
Equipo A:        92    85     78     88     95     (Excelente Cumplimiento)
Equipo B:        88    92     85     82     90     (Excelente Calidad)
Equipo C:        80    88     92     90     85     (Excelente Innovación)

Promedio:        87    88     85     87     90
```

## 🎨 Variantes

### Radar Relleno
Áreas dentro del polígono con color (mayor impacto visual)

### Radar de Líneas
Solo líneas, sin relleno

### Radar Escalonado
Cambios abruptos entre puntos

### Múltiples Escalas
Cada indicador con su propio máximo

## 💡 Interpretación

### Equipo A
```
Fortalezas: Cumplimiento (95), Productividad (92)
Debilidades: Innovación (78)
Perfil: Equipo confiable pero poco innovador
```

### Equipo B
```
Fortalezas: Calidad (92), Innovación (85)
Debilidades: Comunicación (82)
Perfil: Equipo técnico pero menos colaborativo
```

### Equipo C
```
Fortalezas: Innovación (92), Comunicación (90)
Debilidades: Productividad (80)
Perfil: Equipo creativo y colaborativo, mejora en velocidad
```

## 📈 Análisis Comparativo

| Aspecto             | Ganador | Valor | Diferencia |
| ------------------- | ------- | ----- | ---------- |
| Mejor Productividad | A       | 92    | +12 vs C   |
| Mejor Calidad       | B       | 92    | +7 vs C    |
| Mejor Innovación    | C       | 92    | +14 vs A   |
| Mejor Comunicación  | C       | 90    | +8 vs B    |
| Mejor Cumplimiento  | A       | 95    | +5 vs B    |

## 🔧 Mejoras Posibles

- Agregar meta/objetivo (línea de referencia)
- Mostrar tendencia histórica
- Incluir información adicional en tooltip
- Permitir seleccionar qué equipos comparar

## ⚠️ Consideraciones

- Máximo 5-6 dimensiones para claridad
- Seleccionar dimensiones comparables
- Escala debe ser consistente
- Evitar sesgos visuales (orden de indicadores)

## 📊 Ejemplo de Mejora Recomendada

Para Equipo A:
- **Innovación**: 78 → 88 (+10 puntos)
  - Implementar sesiones brainstorming semanal
  - Dedicar 20% del tiempo a proyectos innovadores

Para Equipo B:
- **Comunicación**: 82 → 90 (+8 puntos)
  - Reuniones diarias standup
  - Herramientas de colaboración mejoradas

Para Equipo C:
- **Productividad**: 80 → 88 (+8 puntos)
  - Optimizar procesos repetitivos
  - Reducir reuniones innecesarias
