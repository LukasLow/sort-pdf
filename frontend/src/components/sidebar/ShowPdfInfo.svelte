<script>
    // Svelte 5 Runes Props-Struktur bleibt perfekt erhalten
    let { currentPdf } = $props();
</script>

<!-- Die Hauptbox (bg-zinc-900 mit Rahmen und Schatten) -->
<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl text-zinc-100 font-sans"
>
    <!-- Sektionstitel in hellem Blau -->
    <div class="text-sm font-semibold text-blue-400 mb-2">Aktuelle Datei</div>

    {#if currentPdf}
        <!-- Große Anzeige des aktuellen Dateinamens mit automatischem Zeilenumbruch -->
        <div class="text-base font-mono font-semibold text-zinc-300 break-all">
            {currentPdf}
        </div>

        <!-- Container für die Schlüssel-Wert-Paare -->
        <div class="mt-4">
            <!-- Zeile: Dateiname -->
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Dateiname:</div>
                <div
                    class="font-mono font-semibold text-blue-300 text-right pl-4 break-all"
                >
                    {currentPdf}
                </div>
            </div>

            <!-- Zeile: Dateigröße -->
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Dateigröße in MB:</div>
                <div class="font-mono font-semibold text-zinc-400 text-right">
                    N/A
                </div>
            </div>

            <!-- Zeile: Seitengröße -->
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Seitengröße in KB:</div>
                <div class="font-mono font-semibold text-zinc-400 text-right">
                    N/A
                </div>
            </div>

            <!-- Zeile: DPI -->
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">DPI:</div>
                <div class="font-mono font-semibold text-zinc-400 text-right">
                    N/A
                </div>
            </div>

            <!-- Zeile: OCR -->
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">OCR:</div>
                <div class="font-mono font-semibold text-zinc-400 text-right">
                    N/A
                </div>
            </div>
        </div>
    {:else}
        <!-- Anzeige, wenn keine Datei geladen ist -->
        <div class="text-sm font-semibold text-zinc-400">
            Keine PDF ausgewählt
        </div>
    {/if}

    <!-- Status-Leiste unten (kleinerer blauer Text) -->
    <div class="text-[10px] font-semibold text-blue-400 mt-3">
        🚧 später: Preview + OCR + Analyse
    </div>
</div>
