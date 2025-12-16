# 📊 Gráfico Sankey - Flujo de Ingresos y Gastos

## 📋 Información General

| Propiedad                | Descripción                        |
| ------------------------ | ---------------------------------- |
| **Tipo de Gráfico**      | Sankey Diagram (Diagrama de Flujo) |
| **Nombre en go-echarts** | `charts.NewSankey()`               |
| **Origen**               | Análisis de Flujos                 |
| **Archivo de Salida**    | `flujo_financiero.html`            |
| **Nodos**                | 7 (Origen y destino del dinero)    |

## 🎯 ¿Cuándo Usar?

Los diagramas Sankey son ideales para:

- **Flujos**: Mostrar cómo fluye material, energía o dinero
- **Distribución**: Visualizar cómo se distribuye un recurso
- **Pérdidas**: Identificar dónde se "pierde" el flujo
- **Transformación**: Mostrar cambios de estado o forma
- **Trazabilidad**: Seguir el camino de un elemento

## 🔍 Características

- Nodos (rectángulos) representan puntos
- Flujos (líneas) muestran conexiones y magnitud
- Ancho del flujo proporcional al valor
- Permite múltiples niveles y ramificaciones
- Ideal para mostrar "dónde va el dinero"

## 📊 Caso de Uso: Análisis Financiero Mensual

Este ejemplo muestra cómo **$155,000 de ingresos** se distribuyen en gastos e impuestos:

```
INGRESOS TOTALES: $155,000

Distribución Inicial:
├─ Impuestos:     $25,000 (16%)
├─ Salarios:      $60,000 (39%)
├─ Operativo:     $35,000 (23%)
├─ Marketing:     $20,000 (13%)
└─ R&D:           $15,000 (10%)

NETO RETENIDO (Ganancia):
  Salarios →           $60,000
  Operativo →          $35,000
  Marketing →          $20,000
  R&D →                $15,000
  ─────────────────────────
  Total:              $130,000 (84%)
```

## 💡 Aplicaciones Reales

| Sector         | Aplicación                                         |
| -------------- | -------------------------------------------------- |
| **Finanzas**   | Flujo de efectivo, presupuesto, desglose de gastos |
| **Energía**    | Producción → distribución → consumo                |
| **Agua**       | Afluentes → tratamiento → distribución → uso       |
| **Logística**  | Origen → centros → distribuidores → clientes       |
| **Datos**      | Fuente de datos → procesamiento → análisis         |
| **Producción** | Materia prima → manufactura → venta                |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Hover en flujos, tooltip con valores exactos

## 🚀 Ejecución

```bash
cd grafico-sankey
go run main.go
# Genera: flujo_financiero.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
nodes := []opts.SankeyNode{
    {Name: "Ingresos Totales"},
    {Name: "Impuestos"},
    // ...
}

links := []opts.SankeyLink{
    {Source: "Ingresos Totales", Target: "Impuestos", Value: 25000},
    // ...
}
```

## ⚙️ Configuración

- **Nodos**: 7 puntos de origen/destino
- **Flujos**: 9 conexiones
- **Nivel 1 (Ingreso)**: 1 nodo (Ingresos Totales)
- **Nivel 2 (Gastos)**: 5 nodos (donde va el dinero)
- **Nivel 3 (Retención)**: 1 nodo (Neto Retenido)

## 📊 Visualización del Flujo

```
                            IMPUESTOS
                           ($25,000)
                                ▲
                                │
                                │
INGRESOS TOTALES          SALARIOS
($155,000) ───────────→ ($60,000) ──┐
    │                       ▲       │
    │                       │       │
    │                  OPERATIVO    │
    │                  ($35,000)    │
    ├──────────────────────▲────────│
    │                      │        │
    │                 MARKETING    │
    │                 ($20,000)    │
    ├────────────────────▲─────────│
    │                    │        │
    │                   R&D      │
    │                  ($15,000) │
    └──────────────────▲────────┘
                       │
                       └────────→ NETO RETENIDO
                                 ($130,000)
```

## 📈 Análisis Financiero

### Estructura de Costos

```
Item            Monto      % del Total    % del Ingreso
─────────────────────────────────────────────────────
Salarios        $60,000    46%           39%
Operativo       $35,000    27%           23%
Marketing       $20,000    15%           13%
R&D             $15,000    12%           10%
─────────────────────────────────────────────────────
Total Gastos    $130,000   100%          84%
Impuestos       $25,000    ─             16%
─────────────────────────────────────────────────────
Ingreso Neto    $155,000   ─             100%
Ganancia Neta   $0         ─             0%
```

### Observaciones

⚠️ **Problema Identificado**: 
- Los gastos totales ($130,000) igualan el neto retenido
- No hay ganancia neta en este ejemplo
- Esto indica que TODA la ganancia se destina a operación

### Escenario Alterno (Saludable)

```
Ingresos Totales:      $200,000
Gastos Operativos:     $130,000
Impuestos:             $30,000
─────────────────────────────
Ganancia Neta:         $40,000 (20%)
```

## 🎨 Interpretación Visual

- **Ancho del flujo**: Proporcional al monto
- **Flujos más gruesos**: Mayor asignación presupuestaria
- **Ramificaciones**: Distribución de recursos
- **Convergencia**: Agregación en neto retenido

## 💡 Insights del Ejemplo

### Gastos Principales
1. **Salarios ($60K)**: Mayor proporción - costo humano
2. **Operativo ($35K)**: Segundo gasto - infraestructura
3. **Marketing ($20K)**: Inversión en crecimiento

### Distribución Estratégica
- **39% en Salarios**: Inversión en talento (típico)
- **23% en Operativo**: Necesario para funcionar
- **13% en Marketing**: Aceptable para crecimiento
- **10% en R&D**: Bajo para innovación (considerar aumento)

### Alertas
⚠️ R&D muy bajo (10%) - empresa puede perder competitividad
⚠️ Sin ganancia neta - revisar si esto es sostenible

## 🔧 Mejoras Posibles

- Visualizar múltiples meses para ver tendencia
- Comparar presupuestado vs realizado
- Drill-down en categorías para más detalle
- Añadir nodos intermedios (p.ej., "Gastos Fijos" vs "Variables")
- Mostrar proyecciones futuras

## 📊 Comparativa: Antes vs Después de Optimización

### Escenario Actual
```
Ingresos:    $155,000 → Gastos: $130,000 → Neto: $25,000 (impuestos)
```

### Escenario Optimizado (+30% ingresos)
```
Ingresos:    $201,500 → Gastos: $155,000 → Neto: $46,500 (ganancia)
                                           → Impuestos: $32,500
```

## ⚠️ Limitaciones

- Difícil de leer con muchos nodos (>10)
- Requiere datos precisos de flujos
- No muestra causas de flujos
- Orden de nodos afecta visualización

## 📝 Recomendaciones

1. **Corto Plazo**: 
   - Reducir gastos operativos 5-10%
   - Revisar eficiencia de marketing

2. **Mediano Plazo**:
   - Aumentar ingresos 15-20%
   - Mantener control de costos

3. **Largo Plazo**:
   - Generar ganancia neta del 15-20%
   - Reinvertir 50% en R&D y crecimiento
