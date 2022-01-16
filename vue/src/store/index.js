import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import { Secp256k1HdWallet, SigningCosmosClient } from "@cosmjs/launchpad";

Vue.use(Vuex);

const API = "api";
import {
  TxCreateCagnotte,
  TxAddCagnotte,
  TxCloseCagnotte,
} from "../static/blockchainMethods";
export default new Vuex.Store({
  state: {
    account: {},
    client: null,
  },
  mutations: {
    accountUpdate(state, { account }) {
      state.account = account;
    },
    clientUpdate(state, { client }) {
      state.client = client;
    },
  },
  actions: {
    async accountSignIn({ commit }, { mnemonic }) {
      // eslint-disable-next-line no-async-promise-executor
      return new Promise(async (resolve, reject) => {
        const wallet = await Secp256k1HdWallet.fromMnemonic(mnemonic);

        const [{ address }] = await wallet.getAccounts();
        const url = `${API}/auth/accounts/${address}`;
        const acc = (await axios.get(url)).data;
        if (acc.result.value.address === address) {
          const account = acc.result.value;
          const client = new SigningCosmosClient(API, address, wallet);
          commit("accountUpdate", { account });
          commit("clientUpdate", { client });
          resolve(account);
        } else {
          reject("Account doesn't exist.");
        }
      });
    },
    async accountRegister({ dispatch }, { mnemonic }) {
      // eslint-disable-next-line no-async-promise-executor
      return new Promise(async () => {
        const wallet = await Secp256k1HdWallet.fromMnemonic(mnemonic);
        const [{ address }] = await wallet.getAccounts();
        const url = "backend/subscribe";
        await axios({
          method: "post",
          url: url,
          data: {
            address: address,
          },
        })
          .then(() => {
            dispatch("accountSignIn", { mnemonic }).catch(() => {});
          })
          .catch((error) => {
            if (error.response.status === 400) {
              dispatch("accountSignIn", { mnemonic });
            }
          });
      });
    },

    async accountUpdate({ state, commit }) {
      const url = `${API}/auth/accounts/${state.client.senderAddress}`;
      const acc = (await axios.get(url)).data;
      const account = acc.result.value;
      commit("accountUpdate", { account });
    },

    async createCagnotte({ state }, cagnotte) {
      TxCreateCagnotte(state.client, cagnotte.cagnotteName);
    },
    async addCagnotte({ state }, cagnotte) {
      TxAddCagnotte(state.client, cagnotte.cagnotteName, cagnotte.Amount);
    },
    async closeCagnotte({ state }, cagnotteName) {
      TxCloseCagnotte(state.client, cagnotteName);
    },
  },
});
