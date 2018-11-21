pragma solidity ^0.4.24;

/**
    @title OrderedIntervalList
    @dev List of ordered by intervals with non intersections checks. 
 */
library OrderedIntervalList {
    struct Interval {
        uint64 begin; // inclusive
        uint64 end; // exclusive
    }
    struct Data {
        Interval[] intervals;
    }

    function getSize(Data storage self) internal view returns(uint256 size) {
        return self.intervals.length;
    }

    /**
    @notice Get interval by list index
    @param _index interval index in list
    @return interval tuple
     */
    function get(Data storage self, uint _index) internal view returns(uint64 begin, uint64 end) {
        require(self.intervals.length > 0 && _index < self.intervals.length, "check list is non empty and index have proper bounds");
        Interval memory interval = self.intervals[_index];        
        return (interval.begin, interval.end);    
    }

    /**
    @notice Append interval to the end of list
    @param _begin left bound of interval (inclusive)
    @param _end right bound of interval (exclusive)    
     */
    function append(Data storage self, uint64 _begin, uint64 _end) internal {       
        insert(self, self.intervals.length, _begin, _end);
    }

    /**
    @notice Insert interval by specific index in list
    @dev Method also check that new interval doesn't intersect with existed intervals in list
    @param _index target index in list
    @param _begin left bound of interval (inclusive)
    @param _end right bound of interval (exclusive)        
     */
    function insert(Data storage self, uint256 _index, uint64 _begin, uint64 _end) internal {
        require(_begin < _end, "right bound greater than left bound");      
        require(_index <= self.intervals.length, "valid index bounds");
        uint prev = 0;
        uint prevIndex = 0;
        uint next = 0;
        uint nextRawIndex = _index + 1;
        
        
        if (_index > 0) {
            prevIndex = _index - 1;
            prev = self.intervals[prevIndex].end;
        }        
        
        nextRawIndex = _index + 1;        
        while (nextRawIndex < self.intervals.length && next == 0) {
            Interval memory nextEl = self.intervals[nextRawIndex];
            if (nextEl.end != 0) {
                next = nextEl.end;
            }
            nextRawIndex = nextRawIndex + 1;                        
        }       

        require((prev <= _begin || prev == 0) && (next == 0 || next >= _end), "valid interval bounds");

        
        if (_index == self.intervals.length) {
            self.intervals.push(Interval({begin: _begin, end: _end}));
        } else {
            require(self.intervals[_index].end == 0, "free slot exists");       
            self.intervals[_index] = Interval({begin: _begin, end: _end});
        }
        
    } 

    /**
        @notice Remove interval by index
        @param _index interval index in list        
     */
    function remove(Data storage self, uint64 _index) internal {
        require(_index < self.intervals.length, "valid index bounds");
        delete self.intervals[_index];        
    }    
    
}