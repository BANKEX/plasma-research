pragma solidity ^0.4.24;

import { OrderedIntervalList } from "../OrderedIntervalList.sol";

contract OrderedIntervalListWrapper {
    using OrderedIntervalList for OrderedIntervalList.Data;

    OrderedIntervalList.Data _data;

    function getSize() public view returns(uint _size) {
        return _data.size;
    }

    
    function get(uint index) public view returns(uint64 begin, uint64 end) {

        (begin, end) = _data.get(index);
        //return (begin, end);
        
    }

    function append(uint64 _begin, uint64 _end) public returns(bool) {
        return _data.append(_begin, _end);
    }

    function set(uint64 index, uint64 _begin, uint64 _end) public returns(bool) {
        return _data.insert(index, _begin, _end);
        
    }
  
    function remove( uint64 _index) public returns(bool) {
        return _data.remove(_index);
    }


}
