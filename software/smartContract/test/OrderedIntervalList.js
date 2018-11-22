

const { keccak256, bufferToHex } = require('ethereumjs-util');
const { assertRevert } = require('./helpers/assertRevert.js');

const EVMRevert = require('./helpers/EVMRevert');
const EVMThrow = require('./helpers/EVMThrow');

const OrderedIntervalList = artifacts.require('OrderedIntervalListWrapper');
const BigNumber = web3.BigNumber;

require('chai')
  .use(require('chai-bignumber')(BigNumber))
  .use(require('chai-as-promised'))
  .should();

  async function validateList(listContract, size) {
    const first = (await listContract.first.call());
   
    if (first == 0) {
        return false;
    }
    var curIndex = first;
    var count = 1;

    var list = [];
    (await listContract.getPrevious.call(first)).should.be.bignumber.equal(0);
    while (! (await listContract.getNext.call(curIndex)).equals(new BigNumber(0)) && count < size) {        
        nextIndex = await listContract.getNext.call(curIndex)
        previousIndex = await listContract.getPrevious.call(curIndex)
        currentInterval = await listContract.get.call(curIndex)
        list.push("[" + currentInterval[0].toString() + "," + currentInterval[1].toString() + ")")    
        console.log(list)
      
        if (nextIndex > 0) {
            next = await listContract.get.call(nextIndex);
            console.log(nextIndex.toString())
            console.log(next[0].toString())
            console.log(next[1].toString())
            assert(next[0] >= currentInterval[1])
        }   

        if (previousIndex > 0) {
            next = await listContract.get.call(previousIndex);
            assert(next[1] <= currentInterval[0])
        }        

        count = count + 1;
        curIndex = nextIndex;

    }
    currentInterval = await listContract.get.call(curIndex)
    list.push("[" + currentInterval[0].toString() + "," + currentInterval[1].toString() + ")")    
    console.log(list)
    assert((await listContract.getNext.call(curIndex)).equals(new BigNumber(0)))
   
    currentInterval = await listContract.get.call(curIndex);

    (await listContract.getNext.call(curIndex)).should.be.bignumber.equal(0);

    size.should.be.equal(count)
    
  }

  contract('OrderedIntervalList', function () {
      beforeEach(async function() {
          this.orderedList = await OrderedIntervalList.new()
      });       
  

      describe('insert', function () {
        it('check init state' , async function () {            
            
            assertRevert(this.orderedList.get(0))      
            assertRevert(this.orderedList.get(1));
            assertRevert(this.orderedList.get(-1));
            
        })

        it('insert one', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            const interval = await this.orderedList.get(1);

            interval[0].should.be.bignumber.equal(0);
            interval[1].should.be.bignumber.equal(100);

            await validateList(this.orderedList, 1)
          //  (await this.orderedList.getSize()).should.be.bignumber.equal(1)
        })

        it('insert twice', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.set(1, 0, 101, 200);

            const intervalFirst = await this.orderedList.get(1);
            const intervalSecond = await this.orderedList.get(2);

            intervalFirst[0].should.be.bignumber.equal(0);
            intervalFirst[1].should.be.bignumber.equal(100);
            intervalSecond[0].should.be.bignumber.equal(101);
            intervalSecond[1].should.be.bignumber.equal(200);

            await validateList(this.orderedList, 2)
           // (await this.orderedList.getSize()).should.be.bignumber.equal(2)
        })

        it('insert error', async function(){
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.set(1, 0, 101, 200);
            
            // already inserted position
            assertRevert(this.orderedList.set(2, 0, 100, 200)); 

            // range collision
            assertRevert(this.orderedList.set(1, 2, 150, 200));


            await this.orderedList.set(2, 0, 201, 300);
            
            const interval = await this.orderedList.get(3);
            interval[0].should.be.bignumber.equal(201);
            interval[1].should.be.bignumber.equal(300);     

            // zero interval size
            assertRevert(this.orderedList.set(3, 0, 300, 300))
            // begin and end swapped
            assertRevert(this.orderedList.set(3, 0, 305, 302))

            await validateList(this.orderedList, 3)
            
        })
    })

   
      
    describe('remove', function () {
        it('check init state' , async function () {
            
            
           assertRevert(this.orderedList.remove(0, 0, 100));
           assertRevert(this.orderedList.remove(1, 0, 100));
           assertRevert(this.orderedList.remove(-1, 0, 100)); 
        })
      
         it('full remove one element', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.remove(1, 0, 100);      
            
            assertRevert(this.orderedList.get(1));
            validateList(this.orderedList, 0);

            await this.orderedList.set(0, 0, 0, 100);
            await validateList(this.orderedList, 1);
        })  
        
        it('make hole inside interval', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.remove(1, 50, 70);                 
            
            var interval = await this.orderedList.get(1);

            interval[0].should.be.bignumber.equal(0);
            interval[1].should.be.bignumber.equal(50);

            id = (await this.orderedList.lastIndex.call())      
            var interval = await this.orderedList.get(id)

            interval[0].should.be.bignumber.equal(70);
            interval[1].should.be.bignumber.equal(100);

            await validateList(this.orderedList, 2);

        })
        


        it('make and fill hole', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.set(1, 0, 101, 200);      
            await this.orderedList.set(2, 0 , 201, 300);

            // already inserted position
            this.orderedList.set(1, 0, 100, 150).should.be.rejectedWith(EVMRevert); 


            await this.orderedList.remove(2, 101, 200);
            await this.orderedList.get(2).should.be.rejectedWith(EVMRevert);


            await this.orderedList.set(1, 3, 101, 150);
            
            interval = await this.orderedList.get(4);

            interval[0].should.be.bignumber.equal(101);
            interval[1].should.be.bignumber.equal(150);
            await validateList(this.orderedList, 3)

        })

        it('make and fill hole from multiple elements', async function () {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.set(1, 0, 101, 200);   
            await this.orderedList.set(2, 0, 201, 300);  
            await this.orderedList.set(3, 0, 401, 500);

            await this.orderedList.remove(2, 101, 200)
            await this.orderedList.remove(3, 201, 300)

            // invalid previous element
            await this.orderedList.set(2, 4, 150, 170).should.be.rejectedWith(EVMRevert);
            await this.orderedList.set(1, 3, 150, 170).should.be.rejectedWith(EVMRevert);
            await this.orderedList.set(1, 4, 102, 110);
            await validateList(this.orderedList, 3);
            await this.orderedList.set(1, 5, 101, 102);
            await validateList(this.orderedList, 3)

            var interval = await this.orderedList.get(5);
            interval[0].should.be.bignumber.equal(101);
            interval[1].should.be.bignumber.equal(110);

            await validateList(this.orderedList, 3)

        })
       

        it('invalid range removing', async function () {
            await this.orderedList.set(0, 0, 0, 100)
            await this.orderedList.set(1, 0, 101, 300)

            await this.orderedList.remove(1, 0, 0).should.be.rejectedWith(EVMRevert)
            await this.orderedList.remove(1, 50, 150).should.be.rejectedWith(EVMRevert)
            await this.orderedList.remove(2, 0, 500).should.be.rejectedWith(EVMRevert)

        })



    })

    describe('remove and insert', function() {
        it('concat intervals', async function() {
            await this.orderedList.set(0, 0, 0, 100);
            await this.orderedList.set(1, 0, 101, 200);      
            await this.orderedList.set(2, 0, 201, 300);

            await validateList(this.orderedList, 3)

            await this.orderedList.set(1, 2, 100, 101)
            await this.orderedList.set(1, 3, 200, 201)
           
            await validateList(this.orderedList, 1)

        })
    })




   

  })
