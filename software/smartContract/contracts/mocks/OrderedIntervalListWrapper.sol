pragma solidity ^0.4.24;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";


contract OrderedIntervalListWrapper {
  using OrderedIntervalList for OrderedIntervalList.Data;

  OrderedIntervalList.Data public _data;
  uint public lastInserted;

  function lastIndex() public view returns(uint) {
    return _data.index;
  }

  function first() public view returns(uint) {
    return _data.getFirst();
  }

  function get(uint index) public view returns(uint begin, uint end) {
    OrderedIntervalList.Interval storage interval = _data.get(index);
    return (interval.begin, interval.end);
  }

  function getNext(uint index) public view returns(uint) {
    return _data.intervals[index].next;
  }

  function getPrevious(uint index) public view returns(uint) {
    return _data.intervals[index].previous;
  }

  function set( uint prev, uint next, uint _begin, uint _end) public returns (uint id) {
    id = _data.insert(prev, next, _begin, _end);
    lastInserted = id;
  }

  function remove(uint _index, uint begin, uint end) public {
    _data.remove(_index, begin, end);
  }
}
