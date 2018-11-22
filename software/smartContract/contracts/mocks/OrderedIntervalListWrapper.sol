pragma solidity ^0.4.24;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";

contract OrderedIntervalListWrapper {
    using OrderedIntervalList for OrderedIntervalList.Data;

    OrderedIntervalList.Data public _data;
    function lastIndex() public view returns(uint64) {
        return _data.index;
    }

    function first() public view returns(uint64) {
        return _data.first;
    }

    
    function get(uint64 index) public view returns(uint64 begin, uint64 end) {

        OrderedIntervalList.Interval storage interval = _data.get(index);
        return (interval.begin, interval.end);        
    }

    function getNext(uint64 index) public view returns(uint64) {
        return _data.intervals[index].next;
    }
    function getPrevious(uint64 index) public view returns(uint64) {
        return _data.intervals[index].previous;
    }

    function set( uint64 prev, uint64 next, uint64 _begin, uint64 _end) public returns (uint64 id) {
        id = _data.insert(prev, next, _begin, _end);
        
    }
  
    function remove(uint64 _index, uint64 begin, uint64 end) public {
        _data.remove(_index, begin ,end);
    }


}
