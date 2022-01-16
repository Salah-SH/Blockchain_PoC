<template>
  <div>
    <div @click="onExecute" class="task" :id="'task' + index">
      <h3>
        <B> {{ paiement.cagnottename }} </B>

        {{ paiement.amount }}Dt
      </h3>
      <form
        target="_blank"
        :id="index"
        method="post"
        action="https://sandbox.paymee.tn/gateway/"
      >
        <input
          type="hidden"
          :id="'test' + index"
          name="payment_token"
          :value="paiement.Paiementtoken"
        />
        <input type="hidden" name="url_ok" value="https://example.com/ok.php" />
        <input type="hidden" name="url_ko" value="https://example.com/ko.php" />
      </form>
    </div>
  </div>
</template>

<script >
// import axios from "axios";
export default {
  name: "Paiement",
  props: {
    paiement: Object,
    index: Number,
  },
  methods: {
    async onExecute() {
      const myForm = document.getElementById(this.index);

      document
        .getElementById("task" + this.index)
        .addEventListener("click", function () {
          myForm.submit();
        });
    },
  },
};
</script>

<style scope>
.fas {
  color: red;
}
.task {
  background: rgb(255, 255, 255);
  margin: 5px;
  padding: 1px 20px;
  cursor: pointer;
  width: inherit;
  height: 2rem;
  font-size: small;
  text-align: center;
}
.task.reminder {
  border-left: 1px solid rgb(0, 102, 128);
}
.task h3 {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>