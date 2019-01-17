function updateData() {
    window.amount = [];
    window.data = [];
    window.event_type = [];
    for (let i = 0; i < localStorage.getItem(ITEMS_LEN); i++) {
        const item = JSON.parse(localStorage.getItem(i.toString()));
        window.data.push(item.date);
        window.event_type.push(item.event_type);
        window.amount.push(item.val);
    }
}

function setData() {
    for (let j = 0; j < localStorage.getItem(ITEMS_LEN); j++) {
        document.getElementById("history").innerHTML +=
            `<tbody>\n` +
            `<tr>\n` +
            `<td class=\"text-center\">${data[j]}</td>\n` +
            `<td class=\"text-center\"> ${event_type[j]}</td>\n` +
            `<td class=\"text-center\"> ${amount[j]}</td>\n` +
            `</tr>\n` +
            `</tbody>`
    }
}

function clearData() {
    document.getElementById("history_table").innerHTML =
        `<thead>\n` +
        `<tr>\n` +
        `<th class=\"text-center\">Date</th>\n` +
        `<th class=\"text-center\">Type</th>\n` +
        `<th class=\"text-center\">Sum</th>\n` +
        `</tr>\n` +
        `</thead>\n` +
        `<tbody id=\"history\">\n` +
        `</tbody>`
}
