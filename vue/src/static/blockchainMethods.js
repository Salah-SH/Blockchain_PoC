
import {
  coins,

} from "@cosmjs/launchpad";
import axios from "axios";

const defaultGasLimits = { send: "1000000" };
const fee = {
  amount: coins(1, "nametoken"),
  gas: defaultGasLimits.send,
};



export async function GetCagnotteList(UserAddress) {
  const url = `api/cagnotte/txs/${UserAddress}`
  const cagnottes = (await axios.get(url)).data.result
  return cagnottes
}
export async function GetpaiementTokens(UserAddress) {
  const url = `backend/paiement/${UserAddress}`
  const res = (await axios.get(url)).data.response;


  return res
}

export async function TxCreateCagnotte(client, cagnotteName) {


  let msgCreate = {
    type: "cagnotte/Createcagnotte",
    value: {
      name: cagnotteName,
      creator: client.signerAddress,
    }
  }


  const result = await client.signAndBroadcast([msgCreate], fee)
  return result
}


export async function TxCloseCagnotte(client, cagnotteName) {
  let msgClose = {
    type: "cagnotte/Closecagnotte",
    value: {
      executor: client.signerAddress,
      name: cagnotteName,

    }
  }
  const result = await client.signAndBroadcast([msgClose], fee)
  return result
}

export async function TxAddCagnotte(client, cagnotteName, amount) {

  let msgAddCagnotte = {
    type: "cagnotte/Addcagnotte",
    value: {
      name: cagnotteName,
      bid: amount,
      participator: client.signerAddress,
    }
  }
  const result = await client.signAndBroadcast([msgAddCagnotte], fee)
  return result

}






