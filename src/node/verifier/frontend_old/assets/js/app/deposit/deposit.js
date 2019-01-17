async function deposit() {
    const value = getInputValue("deposit_amount");
    const result = await getRequest(value, "", "deposit");
    let i = localStorage.getItem(ITEMS_LEN) || 0;
    let data = JSON.stringify({date: new Date(), val: value, event_type: "Deposit", txHash: result.txHash});
    await getRequestBal2("test/change" + "/" + current_address + "/" + value);
    localStorage.setItem(i, data);
    i++;
    localStorage.setItem(ITEMS_LEN, i);
}