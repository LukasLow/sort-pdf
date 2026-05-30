function pad(n) {
    return String(n).padStart(2, '0');
}

function sanitize(value) {
    return value.replace(/[/\\:\0]/g, '').replace(/\.\./g, '');
}

export const form = $state({
    year: new Date().getFullYear(),
    month: new Date().getMonth() + 1,
    day: new Date().getDate(),
    correspondent: '',
    info: '',
    extras: ''
});

export function getFilename() {
    const date = `${form.year}-${pad(form.month)}-${pad(form.day)}`;
    const parts = [date, sanitize(form.correspondent), sanitize(form.info), sanitize(form.extras)].filter(Boolean);
    return parts.join('_') + '.pdf';
}
