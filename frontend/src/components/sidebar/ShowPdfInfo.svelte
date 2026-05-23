<script>
    import { GetPdfInfo } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let { currentPdf } = $props();

    let pdfInfo = $state(null);

    function formatMB(bytes) {
        return (bytes / 1_000_000).toFixed(2) + " MB";
    }

    function formatKB(bytes) {
        return (bytes / 1000).toFixed(2) + " KB";
    }

    $effect(() => {
        if (currentPdf) {
            (async () => {
                try {
                    pdfInfo = await GetPdfInfo(currentPdf);
                } catch (e) {
                    console.error("Failed to get PDF info", e);
                    pdfInfo = null;
                }
            })();
        }
    });
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl text-zinc-100 font-sans"
>
    <div class="text-sm font-semibold text-blue-400 mb-2">Aktuelle Datei</div>

    {#if currentPdf}
        <div class="mt-4">
            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Dateiname:</div>

                <div
                    class="font-mono font-semibold text-blue-300 text-right pl-4 break-all"
                >
                    {currentPdf}
                </div>
            </div>

            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Dateigröße:</div>

                <div class="font-mono font-semibold text-zinc-400 text-right">
                    {pdfInfo ? formatMB(pdfInfo.sizeBytes) : "N/A"}
                </div>
            </div>

            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Pro Seite:</div>

                <div class="font-mono font-semibold text-zinc-400 text-right">
                    {pdfInfo && pdfInfo.pageCount > 0
                        ? formatKB(pdfInfo.sizeBytes / pdfInfo.pageCount)
                        : "N/A"}
                </div>
            </div>

            <div class="flex justify-between items-center py-1.5 text-sm">
                <div class="text-zinc-100">Seiten:</div>

                <div class="font-mono font-semibold text-zinc-400 text-right">
                    {pdfInfo ? pdfInfo.pageCount + " s" : "N/A"}
                </div>
            </div>
        </div>
    {:else}
        <div class="text-sm font-semibold text-zinc-400">
            Keine PDF ausgewählt
        </div>
    {/if}
</div>
