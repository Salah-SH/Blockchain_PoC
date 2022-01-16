/* eslint-disable vue/valid-v-for */
<template>
  <div class="container" v-if="hasAddress">
    <h2 type="h2">Your Jackpots</h2>
    <div
      class="item"
      v-for="cagnotte in data"
      :key="cagnotte.name"
      @click="plot"
    >
      <div class="item__field__key">
        <B> Jackpot name : </B>
        <p>{{ cagnotte.name }}</p>
      </div>

      <div class="item__field__key">
        <B> Amount : </B>
        <p>{{ cagnotte.amount}}</p>
      </div>
       
      <div class="item__field__key" >
       <B> Pending Amount : </B>
        <p>{{ cagnotte.pendingamount? cagnotte.pendingamount.reduce((accumulator, currentValue) => accumulator + parseFloat(currentValue.amount),0):0 }}</p>
      </div>
    
      <div class="item__field__value" v-if="cagnotte.Status"></div>
    </div>

    <div v-if="hasAddress && data.length < 1" class="card__empty">
      There are no items yet. Create one using the form below.
      <div v-for="task in data" :key="task.name">
        <h3>{{ task.name }}</h3>
      </div>
    </div>

  </div>
</template>

<style scoped>
h2 {
  display: inline-block;
}
p {
  color: brown;
  display: inline-block;
  font-size: 20px;
  margin-left: 2rem;
}
.container {
  padding: 0;
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  display: inline;
}
.item {
  box-shadow: inset 0 0 0 1px rgb(199, 166, 184);
  background: rgba(255, 255, 255, 0.301);
  margin-bottom: 1rem;
  padding: 1.5rem;
  border-radius: 0.5rem;
  overflow: hidden;
}

.item__field {
  display: grid;
  line-height: 1.5;
  grid-template-columns: 15% 1fr;
  grid-template-rows: 1fr;
  word-break: break-all;
}
.item__field__key {
  color: rgba(0, 0, 0);
  word-break: keep-all;
  overflow: hidden;
}

.card__empty {
  margin-bottom: 1rem;
  border: 1px dashed rgba(0, 0, 0, 0.1);
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  border-radius: 8px;
  color: rgba(0, 0, 0, 0.25);
  text-align: center;
  min-height: 8rem;
}
@keyframes rotate {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(-360deg);
  }
}
@media screen and (max-width: 980px) {
  .narrow {
    padding: 0;
  }
}

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
  top: 100%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: rgb(255, 255, 255);
  color: rgb(33, 36, 33);
  width: 350px;
  height: 150px;
  z-index: 2;
  text-align: left;
  padding: 20px;
  box-sizing: border-box;
}

.popup .close-btn {
  position: absolute;
  right: 10px;
  top: 5px;
  width: 30px;
  height: 30px;
  background: rgb(255, 255, 255);
  color: rgb(19, 40, 129);
  font-size: 25px;
  font-weight: 600;
  line-height: 0px;
  text-align: center;
  border-radius: 50%;
}
</style>

<script>
import { GetCagnotteList } from "../static/blockchainMethods";
export default {
  data: function () {
    return {
      data: [],
      onScreen: false,
      size:Number,
    };
  },

  computed: {
    hasAddress() {
      const { client } = this.$store.state;
      const address = client && client.signerAddress;
      return address;
    },
  },
  mounted() {
    this.onScreen = false;
  },
  watch: {
    async hasAddress() {
      const x = await GetCagnotteList(this.$store.state.account.address);

      if (x.length > 0) {
        this.data = x;
      }
      setInterval(() => {
        this.find(this.$store.state.account.address);
      }, 2000);
    },
  },
  methods: {
    plot() {
      this.onScreen = !this.onScreen;
    },
    async find(addr) {
      const cagnottes = await GetCagnotteList(addr);
      if (cagnottes.length > 0) {
        this.data = cagnottes;
      }
    },
    closePop() {
      if (this.onScreen) {
        this.onScreen = false;
      }
    },
  },
};
</script>
