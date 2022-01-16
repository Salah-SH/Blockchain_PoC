/* eslint-disable vue/valid-v-for */
<template>
  <div class="container">
    <a id="close-image" class="notification" @click="notification"
      >Notifications
      <!-- <button  @click="notification">Notifications
     
    </button> -->
      <span class="badge" v-if="hasAddress">{{ this.size }}</span>
      <img
        src="https://www.pikpng.com/pngl/b/108-1083508_facebook-notification-icon-vector-wwwimgkidcom-the-logo-linkedin.png"
      />
    </a>
    <h2 type="h2" v-if="notif">
      <transition name="fade">
        Paiements
        <div class="my-0">
          <div class="overlay"></div>
          <div class="content">
            <button class="close-btn" @click="closePop">&times;</button>
            <div v-for="(paiement1, index) in paiements" :key="index">
              <Paiement :paiement="paiement1" :index="index" />
            </div>
          </div>
        </div>
      </transition>
    </h2>
  </div>
</template>


<style scoped>
.my-0 {
  background: lawngreen;
  position: absolute;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.overlay {
  top: 0px;
  left: 0px;
  width: 10vw;
  height: 10vh;
  background: rgba(255, 255, 255, 0.07);
  z-index: 1;
  border-radius: 0px;
}

.content {
  border-radius: 0px;
  text-transform: none;
  position: absolute;
  top: 60%;
  left: 100%;
  transform: translate(50%, -50%);
  background: rgba(153, 147, 147, 0.438);
  color: rgb(9, 110, 9);
  width: 10rem;
  height: inherit;
  z-index: 2;
  text-align: center;
  padding: 20px;
  box-sizing: border-box;
  border: 3px;
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
  display: inline-flex;
  height: 65px;
  width: 50px;
}
.notification {
  background-color: inherit;
  color: rgb(255, 255, 255);
  text-decoration: none;
  padding: -3rem 1rem;
  position: relative;
  display: inline-block;
  border-radius: 2px;
  margin-left: 80%;
  size: 2rem;
  height: 3rem;
  width: 4rem;
}

.notification .badge {
  position: absolute;
  top: 0px;
  right: 0px;
  padding: 5px 10px;
  border-radius: 50%;
  background-color: red;
  color: white;
}
.notification:hover {
  background: red;
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
