pragma solidity ^0.5.2;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";


contract OrderedIntervalListWrapper {
  using OrderedIntervalList for OrderedIntervalList.Data;

  OrderedIntervalList.Data private _data;
  uint64 private _lastInserted;

  constructor() public {
    _data.initialize();
  }

  function lastInserted() public view returns(uint64) {
    return _lastInserted;
  }

  function maxIndex() public view returns(uint64) {
    return uint64(_data.intervals.length - 1);
  }

  function firstIndex() public view returns(uint64) {
    return _data.getFirstIndex();
  }

  function lastIndex() public view returns(uint64) {
    return _data.getLastIndex();
  }

  function get(uint64 index) public view returns(uint64 begin, uint64 end) {
    OrderedIntervalList.Interval storage interval = _data.get(index);
    return (interval.begin, interval.end);
  }

  function getNext(uint64 index) public view returns(uint64) {
    return _data.intervals[index].next;
  }

  function getPrev(uint64 index) public view returns(uint64) {
    return _data.intervals[index].prev;
  }

  function insert(uint64 prev, uint64 next, uint64 begin, uint64 end) public returns (uint64 id) {
    id = _data.insert(prev, next, begin, end);
    _lastInserted = id;
  }

  function remove(uint64 index, uint64 begin, uint64 end) public {
    _data.remove(index, begin, end);
  }
}
