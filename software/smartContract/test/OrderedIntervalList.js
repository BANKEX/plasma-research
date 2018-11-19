

const { keccak256, bufferToHex } = require('ethereumjs-util');
const { assertRevert } = require('./helpers/assertRevert.js');

const OrderedIntervalList = artifacts.require('OrderedIntervalListWrapper');
const BigNumber = web3.BigNumber;

require('chai')
  .use(require('chai-bignumber')(BigNumber))
  .should();

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
            await this.orderedList.append(0, 100);
            const interval = await this.orderedList.get(0);

            interval[0].should.be.bignumber.equal(0);
            interval[1].should.be.bignumber.equal(100);
            (await this.orderedList.getSize()).should.be.bignumber.equal(1)
        })

        it('insert twice', async function() {
            await this.orderedList.append(0, 100);
            await this.orderedList.append(100, 200);

            const intervalFirst = await this.orderedList.get(0);
            const intervalSecond = await this.orderedList.get(1);

            intervalFirst[0].should.be.bignumber.equal(0);
            intervalFirst[1].should.be.bignumber.equal(100);
            intervalSecond[0].should.be.bignumber.equal(100);
            intervalSecond[1].should.be.bignumber.equal(200);
            (await this.orderedList.getSize()).should.be.bignumber.equal(2)
        })

        it('insert error', async function(){
            await this.orderedList.append(0, 100);
            await this.orderedList.append(100, 200);
            
            // already inserted position
            assertRevert(this.orderedList.set(1, 100, 200)); 

            // range collision
            assertRevert(this.orderedList.append(150, 200));


            await this.orderedList.append(200, 300);

            const interval = await this.orderedList.get(2);
            interval[0].should.be.bignumber.equal(200);
            interval[1].should.be.bignumber.equal(300);     
            
            
        })
    })

    describe('remove', function () {
        it('check init state' , async function () {
            
            
            this.orderedList.remove(0)
            this.orderedList.remove(1);
            this.orderedList.remove(-1);   


        })

        it('remove one', async function() {
            await this.orderedList.append(0, 100);
            await this.orderedList.remove(0);      
            
            const interval = await this.orderedList.get(0);

            interval[0].should.be.bignumber.equal(0);
            interval[1].should.be.bignumber.equal(0);

            (await this.orderedList.getSize()).should.be.bignumber.equal(1);
        })

        it('make and fill hole', async function() {
            await this.orderedList.append(0, 100);
            await this.orderedList.append(100, 200);      
            await this.orderedList.append(200, 300);

            // already inserted position
            assertRevert(this.orderedList.set(1, 100, 150)); 


            await this.orderedList.remove(1);
            var interval = await this.orderedList.get(1);
            interval[0].should.be.bignumber.equal(0);
            interval[1].should.be.bignumber.equal(0);


            await this.orderedList.set(1, 100, 150);
            
            interval = await this.orderedList.get(1);

            interval[0].should.be.bignumber.equal(100);
            interval[1].should.be.bignumber.equal(150);

            (await this.orderedList.getSize()).should.be.bignumber.equal(3);

        })

    })

  })
