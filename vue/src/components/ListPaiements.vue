/* eslint-disable vue/valid-v-for */
<template>
  <div class="container" v-if="hasAddress">
    <a id="close-image" class="notification" @click="notification">
      <span class="badge" v-if="hasAddress">{{ this.size }}</span>
      <img
        src="https://pics.freeicons.io/uploads/icons/png/7365665041556281661-512.png"
      />
    </a>
    <h2 type="h2" v-if="notif">
      <div class="dropdown">
        <div class="dropdown-content">
          <a v-if="size == 0">You have no notifications </a>
          <div v-else v-for="(paiement1, index) in paiements" :key="index">
            <Paiement :paiement="paiement1" :index="index" />
          </div>
          
        </div>
      </div>
    </h2>
  </div>
</template>


<style scoped>
.dropdown {
  position: inherit;

  right: 1rem;
  display: inline-block;
}

.dropdown-content {
  display: relative;
  position: inherit;
  margin-top: 3rem;
  margin-left: 25rem;
  background-color: rgb(41, 44, 39);
  min-width: 160px;
  overflow: auto;
  box-shadow: 0px 20px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 2;
  padding: 0.5rem;
  border-radius: 5px;
}
.dropdown-content a {
  color: rgba(255, 255, 255);
  padding: 12px 16px;
  text-decoration: transparent;
  display: block;
  font-size: small;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.container {
  position: absolute;
  padding: 0;
  max-width: 900px;
  width: 100%;
  margin-top: 4rem;
  margin-left: 0rem;
  margin-right: auto;
  display: flex;
}
.popup .overlay {
  position: absolute;
  background: rgba(255, 255, 255, 0.07);
}

.popup .content {
  border-radius: 20px;
  text-transform: none;
  position: absolute;

  transform: translate(0%, -50%);
  background: rgba(153, 147, 147);
  color: rgb(9, 110, 9);

  z-index: 2;
  text-align: center;
  padding: 5px;
  box-sizing: border-box;
  border: 2px;
  display: inline-block;
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
.btn {
  background-color: #555;
  color: white;
  text-decoration: none;
  padding: 15px 26px;
  position: relative;
  display: inline-block;
  border-radius: 2px;
}
#close-image img {
  margin-left: 35rem;
  margin-top: 1rem;
  height: 30px;
  width: 30px;
}
.notification {
  background-color: inherit;
  color: rgb(255, 255, 255);
  text-decoration: none;
  padding: -3rem 0rem;
  position: relative;
  display: inline-block;
  border-radius: 2px;
  margin-left: 90%;
  height: 2rem;
  width: 3rem;
}

.notification .badge {
  color: crimson;
  position: absolute;
  top: 0px;
  left: 37rem;
  padding: 1px 10px;
  border-radius: 50%;
  background-color: red;
  color: white;
}
.notification:hover {
  background: white;
}
</style>

<script>
import { GetpaiementTokens } from "../static/blockchainMethods";
import Paiement from "./Paiement";
export default {
  data: function () {
    return {
      paiements: [],
      size: Number,
      notif: false,
    };
  },
  components: {
    Paiement,
  },

  computed: {
    hasAddress() {
      return !!this.$store.state.account.address;
    },
  },

  watch: {
    async hasAddress() {
      const paiementList = await GetpaiementTokens(
        this.$store.state.account.address
      );
      this.size = paiementList.length;

      if (paiementList.length > 0) {
        this.paiements = paiementList;
      }
      setInterval(() => {
        this.findPaiements(this.$store.state.account.address);
      }, 2000);
    },
  },
  methods: {
    closePop() {
      this.notif = !this.notif;
    },
    notification() {
      this.notif = !this.notif;
    },
    async findPaiements(addr) {
      const y = await GetpaiementTokens(addr);
      this.size = y.length;
      if (y.length > 0) {
        this.paiements = y;
      }
    },
  },
};
</script>
