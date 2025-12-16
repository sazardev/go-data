# 📊 Gráfico de Burbujas - Análisis de Productos

## 📋 Información General

| Propiedad                | Descripción                                        |
| ------------------------ | -------------------------------------------------- |
| **Tipo de Gráfico**      | Bubble Chart / EffectScatter (Gráfico de Burbujas) |
| **Nombre en go-echarts** | `charts.NewEffectScatter()`                        |
| **Origen**               | Análisis Multivariante                             |
| **Archivo de Salida**    | `analisis_productos.html`                          |
| **Productos**            | 4 líneas con múltiples variantes                   |

## 🎯 ¿Cuándo Usar?

Los gráficos de burbujas son ideales para:

- **3 Variables**: Mostrar relación entre 3 o 4 variables simultáneamente
- **Proporción**: Tamaño de burbuja representa una 3ª dimensión
- **Cartera de Productos**: Analizar productos por múltiples métricas
- **Matriz Estratégica**: Posicionar items en 2D (p.ej. crecimiento vs participación)
- **Comparación Compleja**: Ver múltiples atributos de forma compacta

## 🔍 Características

- Eje X y Y representan dos variables
- Tamaño de burbuja representa tercera variable
- Color indica categoría/serie
- Permite visualizar hasta 4 dimensiones
- Ideal para 5-20 elementos

## 📊 Caso de Uso: Análisis de Productos

Este ejemplo analiza **4 líneas de productos** por **3 métricas**:

```
VARIABLES:
- Eje X: Precio del producto ($)
- Eje Y: Demanda (unidades vendidas)
- Tamaño Burbuja: Rentabilidad (%)

PRODUCTOS:
Producto A: 3 variantes (100-200 precio)
Producto B: 3 variantes (80-160 precio)
Producto C: 2 variantes (250-280 precio)
Producto D: 2 variantes (50-75 precio)

Total: 10 puntos de datos
```

## 💡 Aplicaciones Reales

| Sector         | Aplicación                                  |
| -------------- | ------------------------------------------- |
| **Retail**     | Precio vs demanda vs margen                 |
| **Portfolio**  | Riesgo vs retorno vs volatilidad            |
| **Publicidad** | Costo vs impresiones vs conversión          |
| **Salud**      | Dosis vs efectividad vs efectos secundarios |
| **Educación**  | Precio vs demanda vs empleabilidad          |
| **Energía**    | Capacidad vs costo vs emisiones             |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Hover en burbujas, tooltip con 3 dimensiones

## 🚀 Ejecución

```bash
cd grafico-burbuja
go run main.go
# Genera: analisis_productos.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
[]opts.EffectScatterData{
    {Value: []interface{}{precio, demanda, rentabilidad}},
    // Ejemplo: {Value: []interface{}{100, 500, 30}}
}
```

## ⚙️ Configuración

- **Eje X**: Precio (rango 50-280)
- **Eje Y**: Demanda (rango 300-900)
- **Tamaño Burbuja**: Rentabilidad (escala visual)
- **Número de Series**: 4 productos
- **Total de Puntos**: 10 burbujas

## 📊 Matriz de Datos Detallada

### Producto A: Premium-Alto Margen

```
Variante    Precio   Demanda   Rentabilidad
───────────────────────────────────────────
A1          $100     500       30%
A2          $150     450       45%
A3          $200     400       50%

Análisis:
✓ Margen crece con precio
⚠ Demanda baja con precio (elasticidad)
→ Estrategia: Premium con margen optimizado
```

### Producto B: Estándar-Volumen

```
Variante    Precio   Demanda   Rentabilidad
───────────────────────────────────────────
B1          $80      700       25%
B2          $120     650       40%
B3          $160     600       55%

Análisis:
✓ Gran demanda a bajo precio
✓ Margen mejora a precio más alto
→ Estrategia: Escalada de precio progresiva
```

### Producto C: Lujo-Alto Precio

```
Variante    Precio   Demanda   Rentabilidad
───────────────────────────────────────────
C1          $250     300       60%
C2          $280     350       70%

Análisis:
✓ Muy alto margen (60-70%)
⚠ Demanda baja (segmento pequeño)
→ Estrategia: Nicho de lujo/premium
```

### Producto D: Entrada-Volumen

```
Variante    Precio   Demanda   Rentabilidad
───────────────────────────────────────────
D1          $50      900       15%
D2          $75      850       28%

Análisis:
✓ Máxima demanda (entrada de mercado)
⚠ Margen bajo (competencia por precio)
→ Estrategia: Volumen de base para clientes
```

## 🎨 Interpretación Visual

```
MAPA ESTRATÉGICO DE PRODUCTOS:
         Demanda Alta
             ▲
             │    
      D1 ●      B1 ●   
             │       A1 ●  C2 ●
             │    B2 ●  A2  C1
             │   ● ●   ●
             │    B3  A3
             │  ●      ●
             │─────────────→ Precio Alto
      
Tamaño de Burbuja = Rentabilidad
```

## 💡 Análisis Competitivo

### Cuadrante 1: Bajo Precio, Alta Demanda
**Producto D**: Entrada de mercado
- Fortaleza: Mayor volumen
- Debilidad: Margen bajo
- Estrategia: Perderlos a versiones premium

### Cuadrante 2: Precio Medio, Alta Demanda
**Producto B**: Mejor relación calidad-precio
- Fortaleza: Volumen + margen equilibrado
- Oportunidad: Escalada de precios
- Estrategia: Core del negocio

### Cuadrante 3: Precio Alto, Baja Demanda
**Producto C**: Lujo/Especializado
- Fortaleza: Máximo margen
- Debilidad: Volumen bajo
- Estrategia: Diferenciación premium

### Cuadrante 4: Precio Alto, Baja Demanda
**Producto A**: Intermedio (transición)
- Situación: Entre B y C
- Oportunidad: Mejorar posicionamiento
- Estrategia: Diferenciación o reduccionar precio

## 📈 Recomendaciones Estratégicas

### 🥇 Producto B: MANTENER Y EXPANDIR
```
Perfil: Líder de mercado (mejor equilibrio)
Acción: 
✓ Aumentar inversión en marketing
✓ Expandir capacidad de producción
✓ Estudiar escalada de precio (D→B)

Pronóstico: Crecimiento de 15-20% anual
```

### 🥈 Producto A: OPTIMIZAR
```
Perfil: Potencial pero subutilizado
Acción:
✓ Mejorar propuesta de valor vs C
✓ Buscar diferenciador técnico
✓ O reducir precio para competir con B

Pronóstico: Consolidarse en segmento
```

### 🥉 Producto C: ESPECIALIZARSE
```
Perfil: Segmento de nicho rentable
Acción:
✓ Marketing enfocado a luxury
✓ Mejorar experiencia de cliente
✓ Crear comunidad de marca

Pronóstico: Estable, crecimiento lento
```

### Producto D: TRANSICIÓN
```
Perfil: Puerta de entrada
Acción:
✓ Usar como anzuelo para subir a B
✓ Programa de lealtad
✓ Bundling con productos más caros

Pronóstico: Retención del 30-40% hacia B
```

## 📊 Simulación de Cambios

### Escenario 1: Aumento de Precio en D
```
Precio D: $50 → $65 (+30%)
Demanda esperada: 900 → 750 (-17%)
Rentabilidad: 15% → 35% (+133%)

Impacto en Ventas:
Actual: 1,750 unidades × promedio $60 = $105K
Futuro: 1,600 unidades × promedio $70 = $112K
Beneficio: +$7K en ventas, +20K en ganancias
```

### Escenario 2: Lanzar Producto E (Intermedio)
```
Posición: Entre B y C ($150-200, demanda 400)
Rentabilidad: 50%

Efecto:
✓ Capturar clientes de B que quieren premium
✓ Generar $20-30K en ventas nuevas
✓ No canibaliza líneas existentes significativamente
```

## 🔧 Mejoras Posibles

- Filtrar por rango de precios
- Ver trayectoria temporal de cada producto
- Proyectar posición futura
- Comparar vs competencia
- Drill-down en variantes individuales

## ⚠️ Limitaciones

- Difícil ver valores exactos (solo aproximado)
- Muchos puntos pueden sobreponerse
- Necesita leyenda clara para colores
- Requiere interactividad para precisión

## 📝 Conclusión

**Portfolio Balanceado**: Combinación exitosa de:
- **Volumen** (Productos B y D): 65% de demanda
- **Margen** (Productos A y C): 40-70% rentabilidad
- **Nicho** (Producto C): Diferenciación premium

**Recomendación**: Invertir en escalada de demanda de B manteniendo márgenes, usar D como embudo hacia B y C, especializar A.
