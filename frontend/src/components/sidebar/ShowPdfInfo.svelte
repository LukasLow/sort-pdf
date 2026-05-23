<script>
    let { currentPdf } = $props();
</script>

<!-- Die Hauptbox -->
<div class="pdf-info-card">
    <div class="section-title">Aktuelle Datei</div>

    {#if currentPdf}
        <!-- Erster Dateiname oben -->
        <div class="filename-header">
            {currentPdf}
        </div>

        <div class="meta-container">
            <!-- Jede Zeile ist jetzt ein kv-row Flexbox-Element -->
            <div class="kv-row">
                <div class="key">Dateiname:</div>
                <div class="value-highlight">{currentPdf}</div>
            </div>

            <div class="kv-row">
                <div class="key">Dateigröße in MB:</div>
                <div class="value-muted">N/A</div>
            </div>

            <div class="kv-row">
                <div class="key">Seitengröße in KB:</div>
                <div class="value-muted">N/A</div>
            </div>

            <div class="kv-row">
                <div class="key">DPI:</div>
                <div class="value-muted">N/A</div>
            </div>

            <div class="kv-row">
                <div class="key">OCR:</div>
                <div class="value-muted">N/A</div>
            </div>
        </div>
    {:else}
        <div class="no-pdf">Keine PDF ausgewählt</div>
    {/if}

    <div class="footer-status">🚧 später: Preview + OCR + Analyse</div>
</div>

<!-- Test-Bereich außerhalb der Box -->
<div class="test-area">
    <div class="text-yellow">Dies sollte gelb sein.</div>
    <div class="text-red">Dies sollte rot sein.</div>
</div>

<style>
    /* Die Hauptbox */
    .pdf-info-card {
        padding: 1rem;
        border-radius: 1rem;
        background-color: #18181b; /* Entspricht bg-zinc-900 */
        border: 1px solid #3f3f46; /* Entspricht border-zinc-700 */
        box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
        color: #f4f4f5; /* Standard-Textfarbe (Weiß/Hellgrau) */
        font-family: sans-serif;
    }

    .section-title {
        font-size: 0.875rem;
        font-weight: 600;
        color: #60a5fa; /* Entspricht text-blue-400 */
        margin-bottom: 0.5rem;
    }

    .filename-header {
        font-size: 1rem;
        font-family: monospace;
        font-weight: 600;
        color: #e4e4e7;
        word-break: break-all;
    }

    .meta-container {
        margin-top: 1rem;
    }

    /* Flexbox-Zeile für die Tabelle */
    .kv-row {
        display: flex;
        justify-content: space-between; /* Drückt Key nach links, Value nach rechts */
        align-items: center;
        padding: 0.35rem 0; /* Angenehmer Zeilenabstand */
        font-size: 0.875rem;
    }

    /* Linke Seite (Keys) */
    .key {
        color: #f4f4f5; /* Explizit hellgrau auf dunklem Grund */
    }

    /* Rechte Seite (Values) */
    .value-highlight {
        font-family: monospace;
        font-weight: 600;
        color: #93c5fd; /* Entspricht text-blue-300 */
        word-break: break-all;
        text-align: right;
        padding-left: 1rem; /* Verhindert Zusammenstoßen bei langen Namen */
    }

    .value-muted {
        font-family: monospace;
        font-weight: 600;
        color: #a1a1aa; /* Grau für N/A */
        text-align: right;
    }

    .no-pdf {
        font-size: 0.875rem;
        font-weight: 600;
        color: #a1a1aa;
    }

    .footer-status {
        font-size: 10px;
        font-weight: 600;
        color: #60a5fa;
        margin-top: 0.75rem;
    }

    /* Test-Klassen */
    .test-area {
        margin-top: 1rem;
    }
    .text-yellow {
        color: #eab308; /* Echtes Gelb */
    }
    .text-red {
        color: #ef4444; /* Echtes Rot */
    }
</style>
