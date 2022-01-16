<template>
  <div>
    <div class="container">
      <div class="h1">Participate in a jackpot</div>
      <div class="cagnotte">
        <input
          type="text"
          v-model="cagnotteName"
          class="cagnotte__input"
          placeholder="Enter the name of your jackpot"
        />
      </div>
      <div class="cagnotte">
        <input
          type="text"
          v-model="amount"
          class="cagnotte__input"
          placeholder="Enter the amount to add in the jackpot"
        />
      </div>

      <button
        :class="['button', `button__enabled__${!!address}`]"
        @click="onSubmit"
        :disabled="!address"
      >
        participate
        <div class="loader" v-if="loader"></div>
      </button>

      <div class="popup" v-if="info">
        <div class="overlay"></div>
        <div class="content">
          <button class="close-btn" @click="closePop">&times;</button>
          <p>your participation has been registered,</p>
          <p>you will receive a notification for the paiement !</p>
        </div>
      </div>
      <div class="popup" v-if="fail">
        <div class="overlay"></div>
        <div class="content">
          <button class="close-btn" @click="closePop">&times;</button>
           <p v-if="code==3">This jackpot is not found</p>
           <p v-if="code==7">This jackpot is closed</p>
          <p v-else>An error occured, try later</p>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { TxAddCagnotte } from "../static/blockchainMethods";

export default {
  name: "AddCx",
  data() {
    return {
      cagnotteName: "",
      amount: "",
      loader: false,
      info: false,
      fail: false,
      code: Number,
    };
  },
  computed: {
    account() {
      return this.$store.state.account;
    },
    address() {
      const { client } = this.$store.state;
      const address = client && client.signerAddress;
      return address;
    },
  },
  methods: {
    onSubmit() {
      if (!this.cagnotteName) {
        alert("add cagnotte name");
        return;
      }
      if (!this.amount) {
        alert("add amount");
        return;
      }
      this.loader = true;
      TxAddCagnotte(
        this.$store.state.client,
        this.cagnotteName,
        this.amount
      ).then((response) => {
        if (response.code) {
          this.code = response.code;
        }
        this.info = !response.code;
        this.fail = !this.info;
        this.loader = false;
      });

      this.cagnotteName = "";
      this.amount = "";
    },
    closePop() {
      if (this.info) {
        this.info = !this.info;
      } else if (this.fail) {
        this.fail = !this.fail;
      }
    },
  },
};
</script>


<style scoped>
.popup .overlay {
  position: fixed;
  top: 0px;
  left: 0px;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.7);
  z-index: 1;
  border-radius: 0px;
}

.popup .content {
  border-radius: 5px;
  text-transform: none;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: rgb(255, 255, 255);
  color: rgb(33, 36, 33);
  width: inherit;
  height: inherit;
  z-index: 2;
  text-align: left;
  padding: 20px;
  box-sizing: border-box;
  font-size:medium;
}

.popup .close-btn {
  position: absolute;
  right: 10px;
  top: 0px;
  width: 15px;
  height: 15px;
  background: rgb(255, 255, 255,0.001);
  color: rgb(19, 40, 129);
  font-size: 20px;
  font-weight: 600;
  line-height: 0px;
  text-align: center;
  border-radius: 50%;
}
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
  margin-left: auto;
  margin-right: auto;
}
.narrow {
  padding-left: 40px;
  padding-right: 40px;
  box-sizing: border-box;
}

.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 1rem;
}
.h2 {
  font-weight: 50;
  text-transform: none;
  letter-spacing: 0.005em;
  margin-bottom: 1rem;
  margin-top: 1rem;
  color: black;
}
.cagnotte {
  margin-top: 0.5rem;
  display: flex;
}
.cagnotte__input {
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
button {
  background: rgba(8, 19, 1, 0.03);
  border: none;
  color: rgba(0, 125, 255);
  padding: 0.75rem;
  font-size: inherit;
  font-weight: 800;
  font-family: inherit;
  text-transform: uppercase;
  margin-top: 0.5rem;
  cursor: pointer;
  transition: opacity 0.1s;
  letter-spacing: 0.03em;
  transition: color 0.25s;
  display: inline-flex;
  align-items: center;
}
button:focus {
  opacity: 0.85;
  outline: none;
}

.button__enabled__false {
  color: rgba(14, 6, 6, 0.082);
  cursor: not-allowed;
  pointer-events: none;
}
.button__enabled__true {
  color: rgba(0, 125, 255);
  font-weight: 1000;
  cursor: pointer;
}
.button__enabled__true:active {
  color: rgba(255, 0, 0, 0.65);
}
.cagnotte__input:focus {
  outline: none;
  background: rgba(0, 0, 0, 0.06);
  color: rgba(0, 0, 0, 0.5);
}
.cagnotte__input::placeholder {
  color: rgba(0, 0, 0, 0.35);
  font-weight: 500;
}
</style>
