function pad(n) {
    return String(n).padStart(2, '0');
}

export const form = $state({
    year: new Date().getFullYear(),
    month: new Date().getMonth() + 1,
    day: new Date().getDate(),
    correspondent: '',
    info: '',
    extras: '',
    tags: []
});

export function getFilename() {
    const date = `${form.year}-${pad(form.month)}-${pad(form.day)}`;
    const parts = [date, form.correspondent, form.info, form.extras].filter(Boolean);
    return parts.join('_') + '.pdf';
}
