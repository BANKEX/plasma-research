<template>
  <div>
    <div v-bind:class="{ config: confReady }">
      Operator Address: {{operator}}
      <p>
      Node address: {{node}}
      <p>
      Smart Contract address: {{smart}}
    </div>
    <button @click="getConfigInfo()">Get Info</button>

    <div v-bind:class="{ pb: pbReady }" >
      Plasma Balance: {{pb}}
    </div>
    <Button @click="getPlasmaBalance()">Plasma Balance</Button>

    <div v-bind:class="{ scb: scbReady }" >
      Smart Contract Balance: {{scb}}
    </div>
    <Button @click="getSCBalance()">Get smart contract balance</Button>
    
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Verifier',
  data () {
    return {
      operator:"",
      node:"",
      smart:"",
      confReady:true,
      pb:"",
      scb:"",
      pbReady:true,
      scbReady:true
    }
  },
  methods:{
    getConfigInfo(){
      axios.get("/conf")
      .then(x=>{
        this.confReady = false
        this.operator = x.data.Operator
        this.node = x.data.Node
        this.smart = x.data.Smart
      })
    },
    getPlasmaBalance(){
       axios.get("/plasmabalance")
      .then(x=>{
        this.pbReady = false
        this.pb = x.data.Balance
      })
    },
    getSCBalance(){
       axios.get("/scgetbalance")
      .then(x=>{
        this.scbReady = false
        this.scb = x.data.Balance
      })
    }

  }
}
</script>
<style>

.config{
  visibility: hidden;
}
.pb{
  visibility: hidden;
}
.scb{
  visibility: hidden;
}

</style>

