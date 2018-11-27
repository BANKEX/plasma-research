pragma solidity ^0.4.24;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";


contract OrderedIntervalListWrapper {
  using OrderedIntervalList for OrderedIntervalList.Data;

  OrderedIntervalList.Data private _data;
  uint private _lastInserted;

  function lastInserted() public view returns(uint) {
      return _lastInserted;
  }

  function maxIndex() public view returns(uint) {
    return _data.intervals.length - 1;
  }

  function firstIndex() public view returns(uint) {
    return _data.getFirstIndex();
  }

  function lastIndex() public view returns(uint) {
    return _data.getLastIndex();
  }

  function get(uint index) public view returns(uint begin, uint end) {
    OrderedIntervalList.Interval storage interval = _data.get(index);
    return (interval.begin, interval.end);
  }

  function getNext(uint index) public view returns(uint) {
    return _data.intervals[index].next;
  }

  function getPrev(uint index) public view returns(uint) {
    return _data.intervals[index].prev;
  }

  function set(uint prev, uint next, uint begin, uint end) public returns (uint id) {
    id = _data.insert(prev, next, begin, end);
    _lastInserted = id;
  }

  function remove(uint index, uint begin, uint end) public {
    _data.remove(index, begin, end);
  }
}
