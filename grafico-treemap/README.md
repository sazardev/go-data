# 📊 Gráfico de Presupuesto - Distribución por Departamento

## 📋 Información General

| Propiedad                | Descripción                               |
| ------------------------ | ----------------------------------------- |
| **Tipo de Gráfico**      | Bar Chart / Treemap (Desglose Jerárquico) |
| **Nombre en go-echarts** | `charts.NewBar()`                         |
| **Origen**               | Análisis Presupuestario                   |
| **Archivo de Salida**    | `distribucion_presupuesto.html`           |
| **Departamentos**        | 6 áreas de la empresa                     |

## 🎯 ¿Cuándo Usar?

Este gráfico es ideal para:

- **Presupuesto**: Distribución de fondos por departamento
- **Proporción**: Mostrar importancia relativa de cada área
- **Asignación de Recursos**: Visualizar inversión por área
- **Comparación Departamental**: Entender prioridades
- **Justificación**: Explicar decisiones de presupuesto

## 🔍 Características

- Barras horizontales o verticales
- Cada barra representa un departamento
- Altura/ancho proporcional a presupuesto
- Fácil comparación visual
- Ideal para 4-8 departamentos

## 📊 Caso de Uso: Presupuesto Empresarial Anual

Este ejemplo muestra cómo **$770,000** de presupuesto total se distribuye entre **6 departamentos**:

```
DISTRIBUCIÓN DE PRESUPUESTO:
1. Ventas:       $250,000 (32.5%)  ← Mayor asignación
2. Desarrollo:   $180,000 (23.4%)
3. Marketing:    $120,000 (15.6%)
4. RH:           $80,000  (10.4%)
5. Admin:        $90,000  (11.7%)
6. Legal:        $50,000  (6.5%)   ← Menor asignación
──────────────────────────────────────────
Total:           $770,000 (100%)
```

## 💡 Aplicaciones Reales

| Sector          | Aplicación                         |
| --------------- | ---------------------------------- |
| **Empresa**     | Presupuesto anual por departamento |
| **Gobierno**    | Gasto público por ministerio       |
| **Universidad** | Fondos por facultad                |
| **Hospital**    | Presupuesto por especialidad       |
| **ONG**         | Fondo por programa o proyecto      |
| **Proyecto**    | Inversión por componente           |

## 🛠️ Tecnología

- **Biblioteca**: go-echarts/v2
- **Backend**: Go
- **Frontend**: ECharts.js (Apache)
- **Interactividad**: Tooltip con valores, tooltip con porcentajes

## 🚀 Ejecución

```bash
cd grafico-treemap
go run main.go
# Genera: distribucion_presupuesto.html (se abre automáticamente en navegador)
```

## 📝 Estructura de Datos

```go
valores := []opts.BarData{
    {Value: 250000},  // Presupuesto Ventas
    {Value: 180000},  // Presupuesto Desarrollo
    // ...
}
```

## ⚙️ Configuración

- **Eje X**: Nombres de departamentos
- **Eje Y**: Monto presupuestario
- **Número de Barras**: 6
- **Escala**: $0 a $300,000
- **Moneda**: USD

## 📊 Desglose Detallado

### 1. Ventas: $250,000 (32.5%) 🥇
```
Justificación:
✓ Mayor generador de ingresos
✓ Inversión en equipos, herramientas
✓ Viajes, conferencias, eventos
✓ Comisiones y bonificaciones

Categorías de Gasto:
├─ Equipo de ventas: $120,000 (48%)
├─ Viajes y eventos: $80,000 (32%)
└─ Herramientas CRM: $50,000 (20%)

ROI Esperado: 5-10x (por cada $1, retorna $5-10)
```

### 2. Desarrollo: $180,000 (23.4%) 🥈
```
Justificación:
✓ Crear/mantener productos
✓ Inversión en talento técnico
✓ Infraestructura y servidores
✓ Formación continua

Categorías de Gasto:
├─ Salarios ingenieros: $100,000 (56%)
├─ Infraestructura: $50,000 (28%)
└─ Herramientas/licencias: $30,000 (16%)

ROI Esperado: 8-12x (base productiva)
```

### 3. Marketing: $120,000 (15.6%) 🥉
```
Justificación:
✓ Generar demanda
✓ Publicidad digital
✓ Contenido marketing
✓ Reputación de marca

Categorías de Gasto:
├─ Publicidad digital: $60,000 (50%)
├─ Contenido/SEO: $35,000 (29%)
└─ Eventos: $25,000 (21%)

ROI Esperado: 3-5x
```

### 4. Admin: $90,000 (11.7%)
```
Justificación:
✓ Gestión administrativa
✓ Finanzas y contabilidad
✓ Legal y compliance
✓ Oficina y operaciones

Categorías de Gasto:
├─ Salarios: $55,000 (61%)
├─ Oficina: $20,000 (22%)
└─ Sistemas: $15,000 (17%)

ROI Esperado: Necesario (soporte)
```

### 5. RH: $80,000 (10.4%)
```
Justificación:
✓ Reclutamiento y selección
✓ Formación de empleados
✓ Relaciones laborales
✓ Beneficios

Categorías de Gasto:
├─ Reclutamiento: $30,000 (37%)
├─ Formación: $25,000 (31%)
└─ Beneficios: $25,000 (31%)

ROI Esperado: Retencion de talento (indirecto)
```

### 6. Legal: $50,000 (6.5%)
```
Justificación:
✓ Asesoría legal
✓ Cumplimiento normativo
✓ Protección de IP
✓ Seguros

Categorías de Gasto:
├─ Asesoría externa: $30,000 (60%)
├─ Seguros: $15,000 (30%)
└─ Compliance: $5,000 (10%)

ROI Esperado: Protección (riesgo mitigation)
```

## 📈 Análisis de Distribución

### Categorías Estratégicas

```
INVERSIÓN POR TIPO:
─────────────────────────────────────────
Generador de Ingresos:
  └─ Ventas + Marketing: $370,000 (48%)

Creación de Valor:
  └─ Desarrollo: $180,000 (23%)

Soporte Operativo:
  └─ Admin + RH + Legal: $220,000 (29%)

BALANCE: Buen equilibrio entre crecimiento y operación
```

### Comparación con Estándares

```
Departamento    Nuestro %    Promedio Industria    Variación
───────────────────────────────────────────────────────────
Ventas          32.5%        25-30%               +2.5% ✓
Desarrollo      23.4%        15-20%               +3-8% (Tecnológico)
Marketing       15.6%        10-15%               +0-5.6% ✓
Admin           11.7%        10-12%               +0-1.7% ✓
RH              10.4%        8-12%                -1.6 a +2.4% ✓
Legal           6.5%         3-5%                 +1.5-3.5% ✓
```

## 🎯 Escenarios Alternativos

### Escenario 1: Enfoque en Crecimiento (+Tech)
```
Ventas:         $200,000 (-20K)
Desarrollo:     $220,000 (+40K)
Marketing:      $140,000 (+20K)
RH:             $70,000  (-10K)
Admin:          $70,000  (-20K)
Legal:          $70,000  (+20K)
Total:          $770,000

Efecto: Más innovación y escalabilidad
```

### Escenario 2: Enfoque en Operaciones (Estable)
```
Ventas:         $230,000 (-20K)
Desarrollo:     $150,000 (-30K)
Marketing:      $120,000 (Igual)
RH:             $100,000 (+20K)
Admin:          $110,000 (+20K)
Legal:          $60,000  (+10K)
Total:          $770,000

Efecto: Menos riesgo, operación más robusta
```

## 💡 Recomendaciones

### ✓ Fortalezas
- Inversión fuerte en ventas → generación de ingresos
- Inversión adecuada en desarrollo → competitividad
- Cobertura legal y compliance → protección

### ⚠️ Oportunidades
- Considerar aumentar Marketing (15% es bajo)
- RH podría crecer (retención de talento es clave)
- Tech podría aumentar si hay capacidad

### 🔴 Alertas
- Si Ventas < 25%, revisar estrategia de ingresos
- Si Desarrollo < 15%, posible obsolescencia de producto

## 📊 Proyección Anual

```
PRESUPUESTO ANUAL PROYECTADO:
══════════════════════════════════════════

INGRESOS ESPERADOS: $3,850,000
  (Basado en ROI de Ventas + Desarrollo)

COSTO TOTAL: $770,000

UTILIDAD BRUTA: $3,080,000 (80% margen)

Después de:
- Impuestos (25%): $770,000
- Costos variables (50%): $1,925,000
─────────────────────────────────
UTILIDAD NETA ESTIMADA: $385,000 (10%)
```

## 🔧 Mejoras Posibles

- Desglose jerárquico (Ventas → Equipo A, Equipo B)
- Series temporales (comparar presupuesto vs real)
- Proyección presupuestaria para próximo año
- Análisis de ROI por departamento
- Presupuesto variable vs fijo

## ⚠️ Limitaciones

- No muestra cómo se gasta el presupuesto internamente
- No captura valor generado (solo inversión)
- Necesita comparación histórica para contexto
- No incluye beneficios indirectos

## 📝 Conclusión

La distribución presupuestaria de **$770,000** refleja una empresa **orientada al crecimiento** con equilibrio operativo. El 48% en Ventas + Marketing asegura generación de ingresos, mientras que 23% en Desarrollo mantiene competitividad técnica.
