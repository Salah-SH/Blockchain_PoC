<template>
  <div>
   <Paiements />
    <Signup />
    <div class="container">
      
      <TypeItem />
       
      <div class="h1">
        <CreateCx @create-cagnotte="CreateCagnotte" />
      </div>
 
      <div class="h1">
        <AddCx @add-cagnottev="AddCagnotte" />
      </div>
      <div class="h1">
        <CloseTx @close-cagnotte="CloseCagnotte" />
      </div>
    </div>
  </div>
</template>

<script>
import Paiements from "./components/ListPaiements";
import CreateCx from "./components/CreateCx";
import AddCx from "./components/AddCx";
import CloseTx from "./components/CloseTx";
import TypeItem from "./components/TypeItem";
import Signup from "./components/Signup";
import { registerUser } from "./static/subscription";
export default {
  name: 'App',
  components: {
    CreateCx,
    AddCx,
    CloseTx,
    TypeItem,
    Paiements,
    Signup,
  },
  computed: {
    address() {
      const { client } = this.$store.state;
      const address = client?.signerAddress;
      return address;
    },
  },
  methods: {
    async Subscribe(userAccount) {
      registerUser(
        "backend/subscribe",
        userAccount.firstName,
        userAccount.lastName,
        userAccount.mnemonic
      );
    },

  },
};
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap");
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
body {
  font-family: "Poppins", sans-serif;
}
.h1 {
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 2rem;
}
.container {
  padding: 4rem 1rem 12rem;
  max-width: 900px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}
.btn {
  display: inline-block;
  background: #365;
  color: #fff;
  border: none;
  padding: 10px 20px;
  margin: 5px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}
.btn:focus {
  outline: none;
}
.btn:active {
  transform: scale(0.98);
}
.btn-block {
  display: block;
  width: 100%;
}
</style>



