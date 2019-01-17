async function request(opts) {
    return (await fetch(opts)).json();
}

function getInputValue(id) {
    return document.getElementById(id).value;
}