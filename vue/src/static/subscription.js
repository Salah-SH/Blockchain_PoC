import {
    Secp256k1HdWallet,

} from "@cosmjs/launchpad";


export async function registerUser(url, firstname, lastname, mnemonic) {

    let wallet = await Secp256k1HdWallet.fromMnemonic(mnemonic);
    const [{ address: address }] = await wallet.getAccounts();

    const result = fetch(url, {
        method: 'POST',
        headers: {
            "Content-Type": 'application/json',
            "Access-Control-Allow-Origin": '*',
        },
        body: JSON.stringify({
            address: address,
            identity: {
                firstname: firstname,
                lastname: lastname,
            },

        })

    }).then((res) => res.json())
    return result
}