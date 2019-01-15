const { BN, shouldFail } = require('openzeppelin-test-helpers');

const OrderedIntervalList = artifacts.require('OrderedIntervalListWrapper');

async function validateList (listContract, arr, ids) {
  arr = arr.map(sub => sub.map(a => new BN(a)));
  ids = ids.map(a => new BN(a));

  // Validate zero element
  const zeroInterval = await listContract.get.call(0);
  zeroInterval[0].should.be.bignumber.equal('0');
  zeroInterval[1].should.be.bignumber.equal('0');
  (await listContract.getPrev.call(0)).should.be.bignumber.equal('0');
  (await listContract.getNext.call(0)).should.be.bignumber.equal('0');

  // Validate linked list with prev refs
  let i = 0;
  let prevId = new BN(0);
  let id = await listContract.firstIndex.call();
  while (i < arr.length && id !== 0) {
    const interval = await listContract.get.call(id);
    interval[0].should.be.bignumber.equal(arr[i][0]);
    interval[1].should.be.bignumber.equal(arr[i][1]);

    if (ids.length > 0) {
      id.should.be.bignumber.equal(ids[i]);
    }

    i++;
    (await listContract.getPrev.call(id)).should.be.bignumber.equal(prevId);
    prevId = id;
    id = await listContract.getNext.call(id);
  }

  i.should.be.equal(arr.length);
  id.should.be.bignumber.equal('0');
  (await listContract.lastIndex.call()).should.be.bignumber.equal(prevId);
}

// async function printList (listContract) {
//   // Validate zero element
//   const zeroInterval = await listContract.get.call(0);
//   const zeroPrev = await listContract.getPrev.call(0);
//   const zeroNext = await listContract.getNext.call(0);
//   let s = 'zeroInterval=(' + zeroInterval[0] + ',' + zeroInterval[1] + ',' + zeroPrev + ',' + zeroNext + ') ';

//   // Validate linked list with prev refs
//   let id = await listContract.firstIndex.call();
//   s += 'last=' + (await listContract.lastIndex.call()) + ' ';
//   s += 'list={ ';
//   while (id !== 0) {
//     const interval = await listContract.get.call(id);
//     const prevId = await listContract.getPrev.call(id);
//     s += id + ':[' + interval[0] + ',' + interval[1] + '):' + prevId + ' ';

//     id = await listContract.getNext.call(id);
//   }
//   s += '}';

//   return s;
// }

contract('OrderedIntervalList', function () {
  beforeEach(async function () {
    this.orderedList = await OrderedIntervalList.new();
  });

  describe('insert', function () {
    it('check init state', async function () {
      await this.orderedList.get.call(0); // fulfilled
      await shouldFail.reverting(
        this.orderedList.get.call(1)
      );
      await shouldFail.reverting(
        this.orderedList.get.call(-1)
      );
      await validateList(this.orderedList, [], []);
    });

    describe('single', function () {
      it('insert single', async function () {
        //
        //  + [0,100)
        //  = [0,100)
        //

        await this.orderedList.insert(0, 0, 0, 100);
        await validateList(this.orderedList, [[0, 100]], [1]);
      });

      it('insert single with wrong range', async function () {
        //
        // !+ [200,100)
        //  = [0,0)
        //

        await shouldFail.reverting(
          this.orderedList.insert(0, 0, 200, 100)
        );
        await validateList(this.orderedList, [], []);
      });

      it('insert single with wrong prev', async function () {
        //
        // !+ [0,100)
        //  = [0,0)
        //

        await shouldFail.reverting(
          this.orderedList.insert(1, 0, 0, 100)
        );
        await validateList(this.orderedList, [], []);
      });

      it('insert single with wrong next', async function () {
        //
        // !+ [0,100)
        //  = [0,0)
        //

        await shouldFail.reverting(
          this.orderedList.insert(0, 1, 0, 100)
        );
        await validateList(this.orderedList, [], []);
      });
    });

    it('insert second', async function () {
      //
      //  + |-----|
      //  +       |-----|
      //  = |-----------|
      //

      await this.orderedList.insert(0, 0, 0, 100);
      await validateList(this.orderedList, [[0, 100]], [1]);
      await this.orderedList.insert(1, 0, 100, 200);
      await validateList(this.orderedList, [[0, 200]], [1]);
    });

    it('fails to insert after rightest with gap', async function () {
      //
      //  + |-----|
      // !+        |-----|
      // !+         |----|
      // !+              |-----|
      //  = |-----|
      //

      await this.orderedList.insert(0, 0, 0, 100);
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 101, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 105, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 200, 300)
      );
      await validateList(this.orderedList, [[0, 100]], [1]);
    });

    it('fails to insert with intersection', async function () {
      //
      //  +        |-----|
      // !+   |-----|
      // !+             |-----|
      // !+        |-----|
      // !+       |------|
      // !+        |------|
      // !+       |-------|
      //  =        |-----|
      //

      await this.orderedList.insert(0, 0, 100, 200);
      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 50, 101)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 50, 101)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 50, 101)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 50, 101)
      );

      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 199, 250)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 199, 250)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 199, 250)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 199, 250)
      );

      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 100, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 100, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 100, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 100, 200)
      );

      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 99, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 99, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 99, 200)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 99, 200)
      );

      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 100, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 100, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 100, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 100, 201)
      );

      await shouldFail.reverting(
        this.orderedList.insert(0, 0, 99, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(0, 1, 99, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 0, 99, 201)
      );
      await shouldFail.reverting(
        this.orderedList.insert(1, 1, 99, 201)
      );

      await validateList(this.orderedList, [[100, 200]], [1]);
    });

    it('insert second on left', async function () {
      //
      //  +       |-----|
      //  + |-----|
      //  = |-----------|
      //

      await this.orderedList.insert(0, 0, 100, 200);
      await validateList(this.orderedList, [[100, 200]], [1]);
      await this.orderedList.insert(0, 1, 0, 100);
      await validateList(this.orderedList, [[0, 200]], [1]);
    });

    it('insert second on left with gap', async function () {
      //
      //  +       |-----|
      //  + |--|
      //  = |--|  |-----|
      //

      await this.orderedList.insert(0, 0, 100, 200);
      await validateList(this.orderedList, [[100, 200]], [1]);
      await this.orderedList.insert(0, 1, 0, 50);
      await validateList(this.orderedList, [[0, 50], [100, 200]], [2, 1]);
    });
  });

  describe('remove', function () {
    it('check init state', async function () {
      await shouldFail.reverting(
        this.orderedList.remove(0, 0, 100)
      );
      await shouldFail.reverting(
        this.orderedList.remove(1, 0, 100)
      );
      await shouldFail.reverting(
        this.orderedList.remove(-1, 0, 100)
      );
    });

    it('full remove one element', async function () {
      //
      //  + |-----|
      //  - |-----|
      //  =
      //

      await this.orderedList.insert(0, 0, 0, 100);
      await this.orderedList.remove(1, 0, 100);
      await validateList(this.orderedList, [], []);

      const firstInterval = await this.orderedList.get.call(1);
      firstInterval[0].should.be.bignumber.equal('0');
      firstInterval[1].should.be.bignumber.equal('0');

      (await this.orderedList.firstIndex.call()).should.be.bignumber.equal('0');
      (await this.orderedList.lastIndex.call()).should.be.bignumber.equal('0');
    });

    it('make hole inside interval', async function () {
      //
      //    |0  |50  |70  |100
      //  + |-------------|
      //  -     |----|
      //  = |---|    |----|
      //

      await this.orderedList.insert(0, 0, 0, 100);
      await this.orderedList.remove(1, 50, 70);
      await validateList(this.orderedList, [[0, 50], [70, 100]], [1, 2]);
    });

    it('make and fill hole', async function () {
      //
      //    |0        |100        |200      |300
      //
      // 1:
      //  + |-------------------------------|
      //  -           |-----------|
      //  = |----1----|           |----2----|
      //
      // 2:
      //  +               |---|
      //  = |----1----|   |-3-|   |----2----|
      //  +           |---|
      //  = |------1----------|   |----2----|
      //  +                   |---|
      //  = |--------------1----------------|
      //

      // 1
      await this.orderedList.insert(0, 0, 0, 300);
      await validateList(this.orderedList, [[0, 300]], [1]);
      await this.orderedList.remove(1, 100, 200);
      await validateList(this.orderedList, [[0, 100], [200, 300]], [1, 2]);

      // 2
      await this.orderedList.insert(1, 2, 130, 170);
      await validateList(this.orderedList, [[0, 100], [130, 170], [200, 300]], [1, 3, 2]);
      await this.orderedList.insert(1, 3, 100, 130);
      await validateList(this.orderedList, [[0, 170], [200, 300]], [1, 2]);
      await this.orderedList.insert(1, 2, 170, 200);
      await validateList(this.orderedList, [[0, 300]], [1]);
    });

    it('make and fill hole from multiple intervals', async function () {
      //
      //    |0        |100      |200      |300      |400      |500
      //
      // 1:
      //  +                                         |----1----|
      //  +                               |----1----|
      //  + |----2----|
      //  +           |----2----|
      //  = |---------2---------|         |---------1---------|
      //
      // 2:
      //  -       |-------|
      //  = |--2--|       |--3--|         |---------1---------|
      //  -                                     |-------|
      //  = |--2--|       |--3--|         |--1--|       |--4--|
      //
      // 3:
      //  +                     |---------|
      //  = |--2--|       |----------3----------|       |--4--|
      //

      // 1
      await this.orderedList.insert(0, 0, 300, 400);
      await validateList(this.orderedList, [[300, 400]], [1]);
      await this.orderedList.insert(1, 0, 400, 500);
      await validateList(this.orderedList, [[300, 500]], [1]);
      await this.orderedList.insert(0, 1, 0, 100);
      await validateList(this.orderedList, [[0, 100], [300, 500]], [2, 1]);
      await this.orderedList.insert(2, 1, 100, 200);
      await validateList(this.orderedList, [[0, 200], [300, 500]], [2, 1]);

      // 2
      await this.orderedList.remove(2, 50, 150);
      await validateList(this.orderedList, [[0, 50], [150, 200], [300, 500]], [2, 3, 1]);
      await this.orderedList.remove(1, 350, 450);
      await validateList(this.orderedList, [[0, 50], [150, 200], [300, 350], [450, 500]], [2, 3, 1, 4]);

      // 3
      await this.orderedList.insert(3, 1, 200, 300);
      await validateList(this.orderedList, [[0, 50], [150, 350], [450, 500]], [2, 3, 4]);
    });

    it('remove first and last element', async function () {
      //
      //    |0        |100      |200      |300      |400      |500
      //
      // 1:
      //  +                                          |---1---|
      //  +                      |---2---|
      //  +            |---3---|
      //  +  |---4---|
      //  =  |---4---| |---3---| |---2---|           |---1---|
      //
      // 2:
      //  -  |---4---|
      //  -                                          |---1---|
      //  =            |---3---| |---2---|
      //

      // 1
      await this.orderedList.insert(0, 0, 401, 500);
      let idLast = await this.orderedList.lastInserted.call();
      await this.orderedList.insert(0, 1, 201, 300);
      let idNewLast = await this.orderedList.lastInserted.call();
      await this.orderedList.insert(0, 2, 101, 200);
      let idNewFirst = await this.orderedList.lastInserted.call();
      await this.orderedList.insert(0, 3, 0, 100);
      let idFirst = await this.orderedList.lastInserted.call();
      await validateList(this.orderedList, [[0, 100], [101, 200], [201, 300], [401, 500]], [4, 3, 2, 1]);

      // 2
      await this.orderedList.remove(idFirst, 0, 100);
      await this.orderedList.remove(idLast, 401, 500);
      await validateList(this.orderedList, [[101, 200], [201, 300]], [3, 2]);

      (await this.orderedList.getNext.call(idNewLast)).should.be.bignumber.equal('0');
      (await this.orderedList.getPrev.call(idNewFirst)).should.be.bignumber.equal('0');
      (await this.orderedList.firstIndex.call()).should.be.bignumber.equal('3');
      (await this.orderedList.lastIndex.call()).should.be.bignumber.equal('2');
    });

    it('invalid range removing', async function () {
      await this.orderedList.insert(0, 0, 101, 300);
      await this.orderedList.insert(0, 1, 0, 100);

      // empty range remove
      await shouldFail.reverting(
        this.orderedList.remove(2, 99, 99)
      );
      // intersect only prefix
      await shouldFail.reverting(
        this.orderedList.remove(2, 50, 150)
      );
      // intersect only suffix
      await shouldFail.reverting(
        this.orderedList.remove(1, 50, 150)
      );
      // range greater than interval
      await shouldFail.reverting(
        this.orderedList.remove(1, 0, 500)
      );
    });
  });

  describe('merge intervals', function () {
    it('full merge', async function () {
      //
      //    |0        |100      |200      |300      |400
      //
      // 1:
      //  +                      |---1---|
      //  +            |---2---|
      //  +  |---3---|
      //  =  |---3---| |---2---| |---1---|
      //
      // 2:
      //  +          |-|
      //  =  |--------3--------| |---1---|
      //  +                    |-|
      //  =  |-------------3-------------|
      //
      // 3:
      //  +                              |----------|
      //  =  |------------------3-------------------|
      //

      // 1
      await this.orderedList.insert(0, 0, 201, 300);
      await this.orderedList.insert(0, 1, 101, 200);
      await this.orderedList.insert(0, 2, 0, 100);
      await validateList(this.orderedList, [[0, 100], [101, 200], [201, 300]], [3, 2, 1]);

      // 2
      await this.orderedList.insert(3, 2, 100, 101);
      await validateList(this.orderedList, [[0, 200], [201, 300]], [3, 1]);
      await this.orderedList.insert(3, 1, 200, 201);
      await validateList(this.orderedList, [[0, 300]], [3]);

      // 3
      await this.orderedList.insert(3, 0, 300, 400);
      await validateList(this.orderedList, [[0, 400]], [3]);
    });

    it('merge with previous interval', async function () {
      //
      //    |0        |100      |200      |300      |400      |500      |600
      //
      // 1:
      //  +                                         |----1----|
      //  +                     |----2----|
      //  +  |----3----|
      //  =  |----3----|        |----2----|         |----1----|
      //
      // 2:
      //  +                                                   |---------|
      //  =  |----3----|        |----2----|         |---------1---------|
      //  +                               |----|
      //  =  |----3----|        |-------2------|    |---------1---------|
      //
      //

      // 1
      await this.orderedList.insert(0, 0, 400, 500);
      await this.orderedList.insert(0, 1, 200, 300);
      await this.orderedList.insert(0, 2, 0, 100);
      await validateList(this.orderedList, [[0, 100], [200, 300], [400, 500]], [3, 2, 1]);

      // 2
      await this.orderedList.insert(1, 0, 500, 600);
      await validateList(this.orderedList, [[0, 100], [200, 300], [400, 600]], [3, 2, 1]);
      await this.orderedList.insert(2, 1, 300, 350);
      await validateList(this.orderedList, [[0, 100], [200, 350], [400, 600]], [3, 2, 1]);
    });

    it('merge with next interval', async function () {
      //
      //    |0        |100      |200      |300      |400      |500      |600
      //
      // 1:
      //  +                                         |----1----|
      //  +                     |----2----|
      //  +     |--3--|
      //  =     |--3--|         |----2----|         |----1----|
      //
      // 2:
      //  +  |--|
      //  =  |----3---|         |----2----|         |----1----|
      //  +                |----|
      //  =  |----3---|    |-------2------|         |----1----|
      //

      // 1
      await this.orderedList.insert(0, 0, 400, 500);
      await this.orderedList.insert(0, 1, 200, 300);
      await this.orderedList.insert(0, 2, 50, 100);
      await validateList(this.orderedList, [[50, 100], [200, 300], [400, 500]], [3, 2, 1]);

      // 2
      await this.orderedList.insert(0, 3, 0, 50);
      await validateList(this.orderedList, [[0, 100], [200, 300], [400, 500]], [3, 2, 1]);
      await this.orderedList.insert(3, 2, 150, 200);
      await validateList(this.orderedList, [[0, 100], [150, 300], [400, 500]], [3, 2, 1]);
    });
  });
});
