<template>
  <div>
    <div class="container">
      <div class="box"  v-if="!address">
        <div class="h1" bold:>Not Registered ??</div>

        <div :class="['button', 'button__enabled__true']" @click="generate">
          Generate your keys
        </div>
      </div>

      
    </div>
    
    <Wallet :password="mnemonic"></Wallet>
  </div>

  
      
</template>

<script>
import * as bip39 from "bip39";
import { Secp256k1HdWallet } from "@cosmjs/launchpad";
import Wallet from "./Wallet";
export default {
  name:'Signup',
  components: { Wallet },
  data() {
    return {
      password: "",
      mnemonic: "",
      error: false,
      messsage: "this is for test",
    };
  },
  computed: {
    account() {
      return this.$store.state.account;
    },
    address() {
      const { client } = this.$store.state;
      const address = client?.signerAddress;
      return address;
    },
    mnemonicValid() {
      return bip39.validateMnemonic(this.passwordClean);
    },
    passwordClean() {
      return this.password.trim();
    },
  },
  methods: {
    async generate() {
      const client = await Secp256k1HdWallet.generate(12);
      this.mnemonic = client.secret.data;
    },
    async signIn() {
      if (this.mnemonicValid && !this.error) {
        const mnemonic = this.passwordClean;
        await this.$store.dispatch("accountRegister", { mnemonic });
      }
    },
  },
};
</script>

<style scoped>
P {
  inline-size: inherit;
  padding-left: 10px;
  margin-left: 4rem;
}
.container {
  padding: 4rem 0rem 0rem;
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 0%;
  margin-top: auto;
  position: relative;
}
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-top: 5rem;
  margin-bottom: 1rem;
  font-size: medium;
}

.button {
  margin-right: 30rem;
  padding: 0rem 1.5rem;
  white-space: nowrap;
  align-items: left;
  justify-content: left;
  font-weight: 700;
  font-size: 1.5rem;
  text-transform: uppercase;
  text-align: center;
  letter-spacing: 0.05em;
  border-radius: 0.25rem;
  transition: all 0.1s;
  user-select: none;
  font-weight: inherit;
  width: 100%;
}
.button__enabled__true {
  color: rgb(255, 255, 255);
  font-weight: 800;
  cursor: pointer;
  background: rgb(31, 146, 150);
}
.box {
  margin-top: 0.5rem;
  display: ruby;
}
.card {
  background-color: inherit;
  color: rgb(211, 16, 16);
  position: relative;
  display: inline-block;

  size: 4rem;
  height: 300px;
  width: 300px;
}
img {
  border-radius: 4px;
  width: 150px;
  height: 100px;
  display: flex;
}
.password {
  margin-top: 0.5rem;
  display: flex;
}
</style>
