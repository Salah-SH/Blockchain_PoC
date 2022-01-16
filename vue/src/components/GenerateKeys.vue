<template>
  <div>
    <div class="container">
      <div class="box">
        <div :class="['button', 'button__enabled__true']" @click="generate">
          Generate your keys
        </div>
      </div>

      <div class="card" v-if="mnemonic">
        <div class="card__icon"></div>
        {{ mnemonic }}
        <div class="card__desc"></div>
      </div>
    </div>
  </div>
</template>


<script>
import * as bip39 from "bip39";
import { Secp256k1HdWallet } from "@cosmjs/launchpad";
export default {
  components: {},
  data() {
    return {
      password: "",
      mnemonic: "",
      error: false,
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
.container {
  padding: 4rem 1rem 1rem;
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 0%;
}

.card {
  background: rgb(255, 255, 255);
  border: 2px solid rgb(146, 146, 148);
  border-radius: 0.25rem;
  color: rgba(0, 0, 0, 1);
  padding: 0.25rem 0.75rem;
  margin: 0.5rem 0;
  width: 75%;
  border-color: mediumseagreen;
}

.card__icon {
  width: 70%;
  height: 0.25rem;
  fill: rgba(167, 20, 20, 1);
  flex-shrink: 0;
}
.card__desc {
  letter-spacing: 0.02em;
  padding: 0 0.5rem;
  word-break: break-all;
}
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
}
.box {
  margin-top: 0.5rem;
  display: flex;
}

.button {
  margin-right: 40rem;
  background: rgba(0, 10, 0, 1);
  padding: 0rem 1.5rem;
  white-space: nowrap;
  align-items: center;
  justify-content: center;
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

.button__enabled__false {
  cursor: not-allowed;
}
.button__enabled__true {
  color: rgb(255, 255, 255);
  font-weight: 700;
  cursor: pointer;
  background: rgb(31, 146, 150);
}
.button__enabled__true:active {
  color: rgba(72, 255, 0, 0.65);
}
</style>
