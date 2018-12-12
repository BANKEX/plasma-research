Deposit  ( one input = amount in Wei ) + Action button

GET: /deposit/sum

---

Transfer ( tw inputs = address (string) and amount in Wei) + Action button

GET: /transfer/address/sum

---

Exit (Just  Action button)

GET: /exit

---

Latest block timer:

GET: /latestBlock

---

Amount of verifiers

GET: /verifiersAmount

---

Amount of users

GET: /usersCount

---

Total balance amount

GET: /totalBalance

---

Operations history

GET /history

Response:   JSON:

Date + Time

Type 

Sum

---

Actions : More info ( Tx hash, Block, Number, Block hash)

GET /history/txhash

Response: 
 Tx hash, Block, Number, Block hash in JSON

