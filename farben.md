hy


# Farb- & Design-System


> Meine persönliche Designsprache für Websites, Apps, Logos und Icons.
> Basis: Dunkles Theme mit Türkis als Leitfarbe.

---

## Farbpalette

| Rolle        | Name        | HEX       | RGB                    | Verwendung                        |
|--------------|-------------|-----------|------------------------|-----------------------------------|
| Background   | Schwarz     | `#0A0A0A` | `rgb(10, 10, 10)`      | Seitenhintergrund, Canvas         |
| Text         | Eisweiß     | `#F0FAFA` | `rgb(240, 250, 250)`   | Fließtext, Überschriften          |
| Primary      | Türkis      | `#37B5BE` | `rgb(55, 181, 190)`    | Buttons, Links, Akzente, Logos    |
| Secondary    | Lila        | `#8B3D8B` | `rgb(139, 61, 139)`    | Badges, Tags, zweite Kategorie    |
| Accent       | Orange      | `#FF8C42` | `rgb(255, 140, 66)`    | Warnungen, CTAs, Highlights       |

### Erweiterte Töne

| Name              | HEX       | Verwendung                                    |
|-------------------|-----------|-----------------------------------------------|
| Türkis gedämpft   | `#A8C8CA` | Sekundärtext, Timestamps, Placeholders        |
| Türkis dunkel     | `#1E5F63` | Hover-States, Card-Borders, subtile Flächen   |
| Hintergrund hell  | `#141414` | Card-Hintergründe, Modal-Overlays             |
| Hintergrund mittel| `#1C1C1C` | Sidebar, Nav-Hintergrund, Code-Blöcke         |

---

## CSS Custom Properties

```css
:root {
  /* Kernfarben */
  --color-bg:         #0A0A0A;
  --color-bg-card:    #141414;
  --color-bg-muted:   #1C1C1C;

  --color-text:       #F0FAFA;
  --color-text-muted: #A8C8CA;

  --color-primary:    #37B5BE;
  --color-primary-dk: #1E5F63;
  --color-secondary:  #8B3D8B;
  --color-accent:     #FF8C42;

  /* Borders */
  --color-border:     rgba(55, 181, 190, 0.2);
  --color-border-dim: rgba(255, 255, 255, 0.06);
}
```

---

## Typografie

### Schriften

| Rolle         | Familie      | Gewicht       | Einsatz                                  |
|---------------|--------------|---------------|------------------------------------------|
| Interface     | Inter        | 600 SemiBold  | UI, Buttons, Labels, Nav, Code           |
| Serif / Brand | Baskerville  | 400 / 700     | Überschriften, Logos, Editorial-Titel    |
| Mono          | JetBrains Mono | 400 / 500   | Code, Log-Ausgaben, Terminal             |

### Typografische Hierarchie

```css
/* Haupt-Headline — Baskerville */
h1 {
  font-family: 'Baskerville', 'Libre Baskerville', serif;
  font-size: clamp(2rem, 5vw, 3.5rem);
  font-weight: 700;
  color: var(--color-text);
  letter-spacing: -0.02em;
}

/* Subheadlines — Baskerville leichter */
h2 {
  font-family: 'Baskerville', serif;
  font-size: clamp(1.4rem, 3vw, 2rem);
  font-weight: 400;
  color: var(--color-primary);
}

/* UI-Text — Inter SemiBold */
body, button, label, nav {
  font-family: 'Inter', sans-serif;
  font-weight: 600;
  font-size: 15px;
  line-height: 1.65;
  color: var(--color-text);
}

/* Fließtext — Inter Regular */
p, li, td {
  font-weight: 400;
  color: var(--color-text);
  font-size: 15px;
}

/* Muted / Sekundär */
.text-muted {
  color: var(--color-text-muted);
  font-size: 13px;
}
```

---

## UI-Komponenten

### Buttons

```css
/* Primary Button */
.btn-primary {
  background: var(--color-primary);
  color: #0A0A0A;            /* Schwarz auf Türkis — sehr guter Kontrast */
  font-family: 'Inter', sans-serif;
  font-weight: 600;
  padding: 10px 22px;
  border-radius: 8px;
  border: none;
  transition: opacity 0.15s;
}
.btn-primary:hover { opacity: 0.85; }

/* Secondary / Ghost Button */
.btn-ghost {
  background: transparent;
  color: var(--color-primary);
  border: 1.5px solid var(--color-primary);
  padding: 10px 22px;
  border-radius: 8px;
  font-weight: 600;
}

/* Accent Button (CTA, Warning) */
.btn-accent {
  background: var(--color-accent);
  color: #0A0A0A;
  font-weight: 600;
  padding: 10px 22px;
  border-radius: 8px;
  border: none;
}
```

### Cards

```css
.card {
  background: var(--color-bg-card);      /* #141414 */
  border: 0.5px solid var(--color-border);
  border-radius: 12px;
  padding: 20px 24px;
}

/* Hervorgehobene Card */
.card-featured {
  border-color: var(--color-primary);
  box-shadow: 0 0 0 1px var(--color-primary);
}
```

### Badges & Tags

```css
/* Primär — Türkis */
.badge-primary {
  background: rgba(55, 181, 190, 0.15);
  color: var(--color-primary);
  border: 0.5px solid rgba(55, 181, 190, 0.35);
  padding: 3px 10px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

/* Sekundär — Lila */
.badge-secondary {
  background: rgba(139, 61, 139, 0.15);
  color: #D97FD9;
  border: 0.5px solid rgba(139, 61, 139, 0.35);
  padding: 3px 10px;
  border-radius: 20px;
  font-size: 12px;
}

/* Warning — Orange */
.badge-warning {
  background: rgba(255, 140, 66, 0.15);
  color: var(--color-accent);
  border: 0.5px solid rgba(255, 140, 66, 0.3);
  padding: 3px 10px;
  border-radius: 20px;
  font-size: 12px;
}
```

### Log-Ausgaben / Terminal

```css
.log {
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 13px;
  background: var(--color-bg-muted);  /* #1C1C1C */
  border-radius: 8px;
  padding: 16px;
  line-height: 1.7;
}

/* Log-Level Farben */
.log-info    { color: var(--color-primary);  }  /* Türkis  — INFO  */
.log-success { color: #4ADE80;               }  /* Grün    — OK    */
.log-warning { color: var(--color-accent);   }  /* Orange  — WARN  */
.log-error   { color: #FF5555;               }  /* Rot     — ERROR */
.log-muted   { color: var(--color-text-muted); } /* Grau   — DEBUG */
```

---

## Logos & Icons

### Farbregeln für Logos

- **Hauptvariante:** Türkis `#37B5BE` auf transparentem oder schwarzem Hintergrund
- **Invertiert:** Schwarz `#0A0A0A` auf Türkis-Fläche
- **Monochrom:** Nur `#F0FAFA` (für helle Hintergründe anderer Plattformen)
- **Niemals:** Lila oder Orange als Logo-Primärfarbe verwenden — diese sind Akzentfarben

### Icon-System

| Kontext           | Farbe          | Hinweis                                      |
|-------------------|----------------|----------------------------------------------|
| Navigation aktiv  | `#37B5BE`      | Türkis zeigt aktiven State                   |
| Navigation inaktiv| `#A8C8CA`      | Gedämpfter Ton für inaktive Icons            |
| Aktion / Button   | `#37B5BE`      | Konsistent mit Primary                       |
| Warnung           | `#FF8C42`      | Orange für Alerts, Hinweise                  |
| Fehler            | `#FF5555`      | Klares Rot — außerhalb der Kernpalette       |
| Erfolg            | `#4ADE80`      | Grün — ebenfalls außerhalb, sparsam nutzen   |

### SVG-Export Empfehlungen

```xml
<!-- Monochrom für Logos — currentColor nutzen -->
<svg fill="currentColor" xmlns="http://www.w3.org/2000/svg">
  <!-- Pfade hier -->
</svg>
```

CSS dann:
```css
.icon { color: var(--color-primary); }
.icon-muted { color: var(--color-text-muted); }
```

---

## Barrierefreiheit (WCAG)

| Kombination                         | Kontrast  | WCAG-Level    |
|-------------------------------------|-----------|---------------|
| Eisweiß `#F0FAFA` auf `#0A0A0A`    | ~19.5:1   | ✅ AAA        |
| Türkis `#37B5BE` auf `#0A0A0A`     | ~7.2:1    | ✅ AAA        |
| Orange `#FF8C42` auf `#0A0A0A`     | ~6.8:1    | ✅ AA (groß)  |
| Lila `#8B3D8B` auf `#0A0A0A`       | ~4.1:1    | ⚠️ AA (nur große Schrift / Icons) |
| Schwarz `#0A0A0A` auf Türkis       | ~7.2:1    | ✅ AAA        |

> **Hinweis:** Lila `#8B3D8B` nur für Badges, Tags, große Elemente verwenden — nicht für kleinen Fließtext auf Schwarz.

---

## Dark-Mode-Tipps

1. **Niemals reines Weiß `#FFFFFF`** auf reinem Schwarz — augenschmerzhaft. Eisweiß `#F0FAFA` ist die richtige Wahl.
2. **Borders sehr subtil halten** — `rgba(55, 181, 190, 0.2)` statt harter Linien.
3. **Tiefe durch Graustufen** — `#0A0A0A` → `#141414` → `#1C1C1C` für layerweise Tiefe (Background → Card → Input).
4. **Glows sparsam** — `box-shadow: 0 0 20px rgba(55, 181, 190, 0.15)` auf Primary-Elementen wirkt elegant, nicht kitschig.
5. **Bilder** immer mit `filter: brightness(0.9)` leicht abdunkeln damit sie nicht zu sehr aus dem Dark Theme herausstechen.

---

## Schnellreferenz

```
Background  #0A0A0A   rgb(10, 10, 10)
Text        #F0FAFA   rgb(240, 250, 250)
Primary     #37B5BE   rgb(55, 181, 190)
Secondary   #8B3D8B   rgb(139, 61, 139)
Accent      #FF8C42   rgb(255, 140, 66)

Fonts: Inter 600 (UI) · Baskerville 400/700 (Titel) · JetBrains Mono (Code)
```

---

*Erstellt mit Claude · Stand 2026*
