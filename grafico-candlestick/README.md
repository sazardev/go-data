# 📊 Gráfico Candlestick - Análisis Técnico de Acciones

## 📋 Información General

| Propiedad                | Descripción                 |
| ------------------------ | --------------------------- |
| **Tipo de Gráfico**      | Candlestick/KLine (Velas)   |
| **Nombre en go-echarts** | `charts.NewKLine()`         |
| **Origen**               | Análisis Técnico Financiero |
| **Archivo de Salida**    | `analisis_tecnico.html`     |
| **Sesiones**             | 10 días de cotización       |

## 🎯 ¿Cuándo Usar?

Los gráficos candlestick son ideales para:

- **Mercados Financieros**: Precio de acciones, criptomonedas, forex
- **Análisis Técnico**: Identificar patrones de precios
- **Trading**: Tomar decisiones de compra/venta
- **Tendencias**: Ver movimiento de precios en periodos
- **Volatilidad**: Analizar amplitud de oscilaciones

## 🔍 Características

- Cada vela representa un período (día, hora, minuto)
- 4 valores: Apertura, Cierre, Máximo, Mínimo
- Vela verde = Cierre > Apertura (alcista)
- Vela roja = Cierre < Apertura (bajista)
- Mecha = Rango máximo a mínimo
- Cuerpo = Rango apertura a cierre

## 📊 Caso de Uso: Cotización de ACME Corp

Este ejemplo muestra la **evolución del precio de una acción** en **10 sesiones de trading**:

```
VELAS (Cada una representa 1 día):
───────────────────────────────────────────
Día  Apertura  Mínimo  Máximo  Cierre  Tipo
───────────────────────────────────────────
1    $100      $95     $110    $105    🟢 Alcista
2    $105      $100    $108    $103    🔴 Bajista
3    $103      $102    $115    $108    🟢 Alcista
4    $108      $105    $118    $112    🟢 Alcista
5    $112      $108    $120    $110    🔴 Bajista
6    $110      $109    $122    $115    🟢 Alcista
7    $115      $112    $125    $118    🟢 Alcista
8    $118      $115    $128    $120    🟢 Alcista
9    $120      $118    $130    $122    🟢 Alcista
10   $122      $120    $132    $125    🟢 Alcista

Tendencia General: ALCISTA
Rango: $95 - $132 (variación de 39%)
Cierre Final: $125 (+25% desde inicio)
```

## 💡 Aplicaciones Reales

| Sector            | Aplicación                  |
| ----------------- | --------------------------- |
| **Finanzas**      | Precio de acciones, bonos   |
| **Criptomonedas** | BTC, ETH, altcoins          |
| **Forex**         | Pares de divisas            |
| **Commodities**   | Oro, petróleo, agriculturas |
| **Índices**       | S&P 500, IBEX, DAX          |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Zoom temporal, tooltip con OCHL

## 🚀 Ejecución

```bash
cd grafico-candlestick
go run main.go
# Genera: analisis_tecnico.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.KlineData{
    {Value: []interface{}{apertura, cierre, mínimo, máximo}},
    // Ejemplo: {Value: []interface{}{100, 105, 95, 110}}
}
```

## ⚙️ Configuración

- **Eje X**: 10 días de cotización
- **Eje Y**: Precio en USD
- **Rango de Precios**: $95 a $132
- **Formato**: Velas tradicionales
- **Colores**: Verde (alcista), Rojo (bajista)

## 📊 Anatomía de una Vela

```
                  Máximo ($110)
                     │
                ┌────┘
                │
Alcista       ┌─┴─┐ ← Cierre ($105)
Día 1         │   │
              │   │
              └─┬─┘ ← Apertura ($100)
                │
                └────┐
                     │
                  Mínimo ($95)

Interpretación:
✓ Cuerpo pequeño: Poco movimiento
✓ Mechas largas: Alta volatilidad
✓ Verde: Compradores ganaron
```

## 🎨 Patrones Candlestick Comunes

### Patrones Alcistas (Señal de Compra)

**1. Hammer (Martillo)**
```
        ┌────┐
        │    │  (Cuerpo pequeño en top)
        │    │
        │    │
        └────┘
           │
        (Mecha larga)
Significado: Rechazo de bajos, probable rebote
```

**2. Three White Soldiers (Tres Soldados Blancos)**
```
┌─┐  ┌─┐  ┌─┐
│ │  │ │  │ │  (Tres velas verdes consecutivas)
│ │  │ │  │ │
└─┘  └─┘  └─┘
Significado: Tendencia alcista fuerte
```

**3. Morning Star (Estrella Matutina)**
```
 ┌─┐          ┌─────┐
 │ │  ┌────┐  │     │  (Patrón de 3 velas)
 │ │  │    │  │     │
 └─┘  └────┘  └─────┘
Significado: Reversión de tendencia bajista
```

### Patrones Bajistas (Señal de Venta)

**1. Shooting Star (Estrella Fugaz)**
```
        ┌────┐
        │    │  (Mecha larga arriba)
        │    │
        │    │
        └────┘
Significado: Rechazo de altos
```

**2. Three Black Crows (Tres Cuervos Negros)**
```
┌─┐  ┌─┐  ┌─┐
│ │  │ │  │ │  (Tres velas rojas consecutivas)
│ │  │ │  │ │
└─┘  └─┘  └─┘
Significado: Tendencia bajista fuerte
```

## 📈 Análisis Técnico del Ejemplo

### Estructura de Tendencia

```
FASE 1: Acumulación (Días 1-2)
├─ Día 1: Alcista pero consolidación
├─ Día 2: Bajista, posible retroceso
└─ Rango: $100-$108

FASE 2: Rompimiento (Día 3-4)
├─ Día 3: Alcista con cierre fuerte
├─ Día 4: Alcista continuación
└─ Break sobre $110

FASE 3: Tendencia (Días 5-10)
├─ 7 de 8 días alcistas (88% alcista!)
├─ Retroceso mínimo en Día 5
└─ Máximo histórico en Día 10

CONCLUSIÓN: Tendencia ALCISTA FUERTE
```

### Niveles Clave

```
Resistencia: $125 (máximo actual, zona de oferta)
Soporte: $120 (retroceso probable)
Media Móvil 10: ~$114 (promedio en período)

Volatilidad: Media (mechas moderadas)
Tendencia: MUY ALCISTA (88% días verdes)
```

## 💡 Decisiones de Trading

### Para Compradores
```
✓ SEÑAL COMPRA FUERTE
  - Tendencia alcista establecida
  - Rompe resistencia
  - Volumen en alza
  
Objetivo: $132 - $140
Stop Loss: $118
Risk/Reward: 1:1.5 (aceptable)
```

### Para Vendedores
```
⚠️ ADVERTENCIA
  - Tendencia muy alcista
  - No hay señal bajista clara
  - Esperar retroceso o reversión
  
Esperar: Vela bajista fuerte o doble techo
Nivel: $130-135
```

## 📊 Métricas Técnicas

```
Rendimiento:
├─ Total: +25% ($100 → $125)
├─ Máximo Intraday: +32% ($100 → $132)
├─ Drawdown: -3% (caída máxima desde último high)
└─ Ratio Win/Loss: 8:1 (8 verdes vs 1 roja)

Volatilidad:
├─ Rango Promedio: $5.3 por día
├─ Amplitud Máxima: $10 (Día 4)
└─ Amplitud Mínima: $2 (Día 2)
```

## 🎯 Estrategias Posibles

### Estrategia 1: Seguir Tendencia (Trend Following)
```
Lógica: La tendencia es tu amigo
Entrada: Confirmación de alza (cierre > apertura)
Salida: Nivel de resistencia o reversión
Riesgo: Bajo-Medio

En Nuestro Caso: MUY FAVORABLE
✓ Tendencia clara
✓ Múltiples confirmaciones
✓ Objetivo: $140+
```

### Estrategia 2: Retroceso (Pullback)
```
Lógica: Esperar retracción para mejor entrada
Entrada: En soporte ($118-120) con confirmación
Salida: Nuevo máximo
Riesgo: Bajo

En Nuestro Caso: ESPERAR
- Día 5 fue pequeño retroceso
- Próximo: $120-122 es soporte
```

### Estrategia 3: Breakout
```
Lógica: Comprar resistencias rotas
Entrada: Cierre > $120
Salida: Nuevo máximo + expansión
Riesgo: Medio

En Nuestro Caso: YA SUCEDIÓ
✓ Rompió $110 (Día 3)
✓ Rompió $120 (Día 7)
✓ Próxima: $130-132
```

## ⚠️ Riesgos y Consideraciones

### 1. Reversión de Tendencia
```
Señales de Alerta:
⚠️ Vela bajista fuerte (body grande)
⚠️ Cierre bajo vs apertura
⚠️ Volumen en baja
⚠️ Rechazo en resistencia

Acción: Establecer stop loss en $120
```

### 2. Volatilidad Inesperada
```
Causas Posibles:
- Resultados de empresa
- Noticias macroeconómicas
- Cambio en sentimiento
- Liquidez limitada

Defensa: Stop loss + posición reducida
```

### 3. Falsa Ruptura (Fakeout)
```
Patrón: Rompe resistencia pero retrocede
Riesgo: Trampa para traders
Defensa: Esperar confirmación (2-3 velas)
```

## 🔧 Mejoras Posibles

- Agregar bandas de Bollinger (volatilidad)
- Volumen de trading por vela
- Promedio móvil (SMA/EMA)
- RSI (Relative Strength Index)
- MACD (Moving Average Convergence Divergence)
- Soporte/resistencia automático

## 📝 Conclusión Técnica

**COMPRA CONFIRMADA**:
- Tendencia alcista establecida (88% días verdes)
- Rompimientos de resistencia confirmados
- Objetivo técnico: $130-140
- Stop Loss recomendado: $118
- Riesgo/Recompensa: 1:1.5 (favorable)

**Nota**: Este es análisis técnico. Las decisiones reales deben considerar también fundamentos, sentimiento de mercado y tolerancia al riesgo.
