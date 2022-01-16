<template>
  <div>
    <div class="container">
      <div class="h1">Close a jackpot</div>
      <div class="cagnotte">
        <input
          type="text"
          v-model="cagnotteName"
          class="cagnotte__input"
          placeholder="Enter the name of your jackpot"
        />

        <div
          :class="['button', `button__enabled__${!!address}`]"
          @click="onSubmit"
        >
          close
          <div class="loader" v-if="loader"></div>
        </div>
      </div>
      <div class="popup" v-if="success">
        <div class="overlay"></div>
        <div class="content">
          <button class="close-btn" @click="closePop">&times;</button>
          <p>your cagnotte has been closed successfully</p>
        </div>
      </div>
      <div class="popup" v-if="fail">
        <div class="overlay"></div>
        <div class="content">
          <button class="close-btn" @click="closePop">&times;</button>

          <p v-if="code == 7">
            The cagnotte {{ this.cagnotteName }} is already closed
          </p>
          <p v-if="code == 18">
            The cagnotte {{ this.cagnotteName }} cannot be closed
          </p>
          <p v-if="code == 3">
            The cagnotte "{{ this.cagnotteName }}" is not registered
          </p>
          <p v-if="code == 4">
            Unauthorized action
          </p>
          <B v-else>
           An error occured
          </B>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { TxCloseCagnotte } from "../static/blockchainMethods";
export default {
  name: "CloseTx",
  data() {
    return {
      cagnotteName: "",
      loader: true,
      success: false,
      fail: false,
      code: Number,
    };
  },
  mounted() {
    this.loader = false;
  },
  computed: {
    isDisabled() {
      return false;
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
      this.loader = true;
      TxCloseCagnotte(this.$store.state.client, this.cagnotteName).then(
        (response) => {
          if (response.code) {
            this.code = response.code;
          }

          this.success = !response.code;
          this.fail = !this.success;
          this.loader = false;
        }
      );
    },
    closePop() {
      if (this.success) {
        this.success = !this.success;
      } else if (this.fail) {
        this.fail = !this.fail;
      }
      this.cagnotteName = "";
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
  top: 5px;
  width: 15px;
  height: 15px;
  background: rgb(255, 255, 255);
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
  color: rgba(0, 125, 255, 0.65);
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
