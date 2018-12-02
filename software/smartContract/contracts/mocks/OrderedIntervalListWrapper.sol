pragma solidity ^0.4.24;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";


contract OrderedIntervalListWrapper {
  using OrderedIntervalList for OrderedIntervalList.Data;

  OrderedIntervalList.Data private _data;
  uint256 private _lastInserted;

  constructor() public {
    _data.initialize();
  }

  function lastInserted() public view returns(uint256) {
    return _lastInserted;
  }

  function maxIndex() public view returns(uint256) {
    return _data.intervals.length - 1;
  }

  function firstIndex() public view returns(uint256) {
    return _data.getFirstIndex();
  }

  function lastIndex() public view returns(uint256) {
    return _data.getLastIndex();
  }

  function get(uint256 index) public view returns(uint256 begin, uint256 end) {
    OrderedIntervalList.Interval storage interval = _data.get(index);
    return (interval.begin, interval.end);
  }

  function getNext(uint256 index) public view returns(uint256) {
    return _data.intervals[index].next;
  }

  function getPrev(uint256 index) public view returns(uint256) {
    return _data.intervals[index].prev;
  }

  function insert(uint256 prev, uint256 next, uint256 begin, uint256 end) public returns (uint256 id) {
    id = _data.insert(prev, next, begin, end);
    _lastInserted = id;
  }

  function remove(uint256 index, uint256 begin, uint256 end) public {
    _data.remove(index, begin, end);
  }
}
