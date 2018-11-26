<template>
  <div>
    <div v-bind:class="{ config: confReady }">
      Operator Address: {{operator}}

      Node address: {{node}}

      Smart Contract address: {{smart}}
    </div>
    <button @click="getConfigInfo()">Get Info</button>

    <div v-bind:class="{ pb: pbReady }" >
      Plasma Balance: {{pb}}
    </div>
    <button @click="getPlasmaBalance()">Plasma Balance</button>

    <div v-bind:class="{ scb: scbReady }" >
      Smart Contract Balance: {{scb}}
    </div>
    <button @click="getSCBalance()">Get smart contract balance</button>
    <br>
    <br>
    <label for="amount-metamask">Amount</label>
    <input id="amount-metamask" v-on:input="amountM = $event.target.value">
    <br>
    <button @click="depositViaMetaMask(amountM)">Deposit via MetaMask</button>
    <br>
    <br>
    <label for="private-key">Private Key</label>
    <input id="private-key" v-on:input="privateKey = $event.target.value">
    <br>
    <label for="amount">Amount</label>
    <input id="amount" v-on:input="amount = $event.target.value">
    <br>
    <button @click="deposit(privateKey, amount)">Deposit without MetaMask</button>
    <br>
    <br>
    <div v-bind:class="{ dth: depositReady }" >
      Tx-hash: {{dth}}
    </div>
    <br>
    <br>
    <div v-bind:class="{ error: errorReady }">
      {{error}}
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Web3 from 'web3'
import ethereumjs from 'ethereumjs-tx'
import keythereum from 'keythereum'

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
      dth:"",
      depositTxHash:"",
      pbReady:true,
      scbReady:true,
      depositReady:true,
      errorReady:true,
      ABI: [{"constant":false,"inputs":[{"name":"operator","type":"address"},{"name":"","type":"address"},{"name":"tokenId","type":"uint256"},{"name":"","type":"bytes"}],"name":"onERC721Received","outputs":[{"name":"","type":"bytes4"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"fromIndex","type":"uint256"},{"name":"newBlocks","type":"bytes"},{"name":"protectedBlockNumber","type":"uint256"},{"name":"protectedBlockHash","type":"address"}],"name":"submitBlocks","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"fromIndex","type":"uint256"},{"name":"newBlocks","type":"bytes"},{"name":"protectedBlockNumber","type":"uint256"},{"name":"protectedBlockHash","type":"address"},{"name":"rsv","type":"bytes"}],"name":"submitBlocksSigned","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"renounceOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"blocksLength","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"isOwner","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"token","type":"address"},{"name":"amount","type":"uint256"}],"name":"depositERC20","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"deposit","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"token","type":"address"},{"name":"tokenId","type":"uint256"}],"name":"depositERC721","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"token","type":"address"},{"name":"tokenId","type":"uint256"}],"name":"calculateAssetId","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"name":"i","type":"uint256"}],"name":"blocks","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"MAIN_COIN_ADDRESS","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"token","type":"address"},{"indexed":true,"name":"who","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"AssetDeposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"who","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"CoinDeposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"token","type":"address"},{"indexed":true,"name":"who","type":"address"},{"indexed":false,"name":"amount","type":"uint256"}],"name":"ERC20Deposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"token","type":"address"},{"indexed":true,"name":"who","type":"address"},{"indexed":true,"name":"tokenId","type":"uint256"}],"name":"ERC721Deposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"length","type":"uint256"},{"indexed":false,"name":"time","type":"uint256"}],"name":"BlocksSubmitted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"previousOwner","type":"address"},{"indexed":true,"name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"}],
      scAddress: document.getElementById("contract").innerText.substring(1, document.getElementById("contract").innerText.length-1),
      error:"",
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
    },
    async deposit(privateKey, amount) {

      const instance = this.blockchain().getInstance(this.ABI, this.scAddress);
      const transactionData = this.blockchain().getCallData(instance, "deposit", []);
      const rawTransaction = await this.blockchain().signTransaction(privateKey, this.scAddress, amount, transactionData);
      const tx = await this.blockchain().sendSignedTransaction(rawTransaction).catch(err=>{
        this.error = err.message;
        this.errorReady = false;
      });
      this.dth = tx.transactionHash;
      this.depositReady = false;
    },
    async depositViaMetaMask(amount) {
      const instance = this.blockchain().getInstance(this.ABI, this.scAddress);
      const transactionData = this.blockchain().getCallData(instance, "deposit", []);
      const tx = await this.blockchain().sendTransactionViaMetaMask(this.scAddress, amount, transactionData).catch(err=>{
        this.error = err.message;
        this.errorReady = false;
      });
      this.dth = tx.transactionHash;
      this.depositReady = false;
    },
    blockchain() {
      return {
        async get(instance, method, parameters) {
          return await instance.methods[method](...parameters).call();
        },
        async sendSignedTransaction(rawTransaction) {
          return await web3.eth.sendSignedTransaction(rawTransaction);
        },
        async sendTransactionViaMetaMask(receiver, amount, transactionData) {
          const txParam = {
            to: receiver,
            from: web3.eth.accounts.givenProvider.selectedAddress,
            value: amount,
            data: transactionData !== undefined ? transactionData : '',
            gasPrice: 30000000000,
            gas: 210000
          };
          console.log(txParam)
          return await web3.eth.sendTransaction(txParam)
        },
        async signTransaction(privateKey, receiver, amount, transactionData) {
          window.web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:9545'));
          const userAddress = this.getAddress(privateKey);
          const txParam = {
            nonce: await web3.eth.getTransactionCount(userAddress),
            to: receiver,
            value: Number(amount),
            from: userAddress,
            data: transactionData !== undefined ? transactionData : '',
            gasPrice: "0x3133333639",
            gas: "0x33450"
          };
          console.log(txParam)
          const privateKeyBuffer = keythereum.str2buf(privateKey.substring(2), "hex");
          const tx = new ethereumjs(txParam);
          tx.sign(privateKeyBuffer);
          const serializedTx = tx.serialize();
          return '0x' + serializedTx.toString('hex');
        },
        getCallData(instance, method, parameters) {
          return instance.methods[method](...parameters).encodeABI();
        },
        getInstance(ABI, address) {
          return new web3.eth.Contract(ABI, address);
        },
        getAddress(privateKey) {
          let _privateKey = privateKey.substring(2, privateKey.length);
          return keythereum.privateKeyToAddress(keythereum.str2buf(_privateKey));
        }
      }
    }
  }
}


window.web3 = new Web3(web3.currentProvider || new Web3.providers.HttpProvider('http://127.0.0.1:9545'));


</script>
<style>

.config, .dth, .pb, .scb, .error {
  visibility: hidden;
}

</style>

