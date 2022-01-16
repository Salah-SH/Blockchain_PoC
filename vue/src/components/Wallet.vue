<template>
  <div>
    <div class="container">
      <div class="h1">{{ address ? "Your Account" : "Sign in" }}</div>
      <div v-if="!address" class="password">
        <input
          type="text"
          v-model="password"
          class="password__input"
          placeholder="Password (mnemonic)"
        />
        <div
          :class="[
            'button',
            `button__error__${!!error}`,
            `button__enabled__${!!mnemonicValid}`,
          ]"
          @click="signIn"
        >
          Sign in
          <div class="loader" v-if="loader"></div>
        </div>
      </div>
      <div v-else class="card">
        <div class="card">
          <div class="card__row">
            <div class="card__icon">
              <icon-user />
            </div>
            <div class="card__desc">
              {{ address }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <p>{{ message }}</p>
  </div>
</template>

<script>
import IconUser from "@/components/IconUser.vue";
import * as bip39 from "bip39";

export default {
  name:'Wallet',
  props: ["message", "password"],
  components: {
    IconUser,
  },
  data() {
    return {
      // password: "",
      error: false,
      loader: false,
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
      console.log("this is validate mnemonic", this.passwordClean)
      return bip39.validateMnemonic(this.passwordClean);
    },
    passwordClean() {
      return this.password.trim();
    },
  },
  methods: {
    async signIn() {
      this.loader = true;

      if (this.mnemonicValid && !this.error) {
        const mnemonic = this.passwordClean;
        await this.$store
          .dispatch("accountRegister", { mnemonic })
          .then(() => {
            this.loader = false;
          })
          .catch(() => {
            this.loader = false;
          });

        
      }
    },
  },
};
</script>

<style scoped>
.loader {
  transform: translate(-50%, -50%);
  border: 10px solid #f3f3f3; /* Light grey */
  border-top: 8px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 45px;
  height: 45px;
  animation: spin 2s linear infinite;
  margin-left: 1rem;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
.container {
  padding: 1rem 1rem 1rem;
  max-width: 900px;
  width: 100%;
  height: 20%;
  margin-left: auto;
  margin-right: auto;
}
.narrow {
  padding-left: 40px;
  padding-right: 40px;
  box-sizing: border-box;
}

.card {
  background: rgba(0, 0, 0, 0.03);
  border-radius: 0.25rem;
  color: rgba(0, 0, 0, 0.5);
  padding: 0.25rem 0.75rem;
}
.card__row {
  display: flex;
  align-items: center;
  margin: 0.5rem 0;
  color: rgba(0, 0, 0, 0.25);
  font-size: 0.875rem;
  font-weight: 500;
  line-height: 1.5;
}
.card__icon {
  width: 1.75rem;
  height: 1.75rem;
  fill: rgba(0, 0, 0, 0.15);
  flex-shrink: 0;
}
.card__desc {
  letter-spacing: 0.02em;
  padding: 0 0.5rem;
  word-break: break-all;
}
.h1 {
  margin-right: 5rem;
  widows: 1rem;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
  margin-top: 0rem;
}
.password {
  margin-top: 0.5rem;
  display: flex;
}
.password__input {
  border: none;
  width: 100%;
  padding: 0.75rem;
  box-sizing: border-box;
  font-family: inherit;
  background: rgba(0, 0, 0, 0.03);
  font-size: 0.85rem;
  border-radius: 0.25rem;
  color: rgba(0, 0, 0, 0.5);
}
.button {
  margin-left: 0.2rem;
  background: rgba(0, 0, 0, 0.03);
  padding: 0 1.5rem;
  white-space: nowrap;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.85rem;
  color: rgba(0, 0, 0, 0.25);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  border-radius: 0.25rem;
  transition: all 0.1s;
  user-select: none;
}
.button.button__error__true {
  animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  background: rgba(255, 228, 228, 0.5);
  color: rgb(255, 0, 0);
}
.button__enabled__false {
  cursor: not-allowed;
    pointer-events:none;

}
.button__enabled__true {
  color: rgba(0, 125, 255);
  font-weight: 700;
  cursor: pointer;
}
.button__enabled__true:active {
  color: rgba(255, 0, 43, 0.65);
}
.password__input:focus {
  outline: none;
  background: rgba(0, 0, 0, 0.06);
  color: rgba(0, 0, 0, 0.5);
}
.password__input::placeholder {
  color: rgba(0, 0, 0, 0.35);
  font-weight: 500;
}
.coin__amount {
  text-transform: uppercase;
  font-size: 0.75rem;
  letter-spacing: 0.02em;
  font-weight: 600;
}
.coin__amount:after {
  content: ",";
  margin-right: 0.25em;
}
.coin__amount:last-child:after {
  content: "";
  margin-right: initial;
}
@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }
  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }
  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }
  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}
</style>
