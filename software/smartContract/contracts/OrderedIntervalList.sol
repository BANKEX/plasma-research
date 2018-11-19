pragma solidity ^0.4.24;

library OrderedIntervalList {
    struct Interval {
        uint64 begin;
        uint64 end; // exclusion
    }
    struct Data {
        mapping (bytes32 => Interval) intervals;
        uint64 size;

    }

    function getSize(Data storage self) internal view returns(uint64 size) {
        return self.size;
    }

    function get(Data storage self, uint index) internal returns(uint64 begin, uint64 end) {
        require(self.size > 0 && index < self.size && index >= 0);
        Interval memory interval = self.intervals[getIndex(index)];
        return (interval.begin, interval.end);    
    }
    function append(Data storage self, uint64 _begin, uint64 _end) internal returns(bool) {
        uint64 newIndex = self.size;
        self.size = newIndex + 1;
        insert(self, newIndex, _begin, _end);
    }

    function insert(Data storage self, uint64 index, uint64 _begin, uint64 _end) internal returns(bool) {
        require(_begin <= _end);
        require(index < self.size);
        uint prev = 0;
        bytes32 prevIndex = 0;
        uint next = 0;
        uint64 nextRawIndex = index + 1;
        
        
        if (index > 0) {
            prevIndex = getIndex(index - 1);
            prev = self.intervals[prevIndex].end;
        }        

        nextRawIndex = index + 1;
        bytes32 nextIndex = 0;
        while (nextRawIndex < self.size && next == 0) {
            nextIndex = getIndex(nextRawIndex);
            Interval memory nextEl = self.intervals[nextIndex];
            if (nextEl.end != 0) {
                next = nextEl.end;
            }
            nextRawIndex = nextRawIndex + 1;
                        
        }
        if (nextRawIndex >= self.size) {
            self.size = nextRawIndex;
        }

        require((prev <= _begin || prev == 0) && (next == 0 || next >= _end));
        bytes32 curIndexHash = getIndex(index);
        require(self.intervals[curIndexHash].end == 0, "check free slot");       

        self.intervals[curIndexHash] = Interval({begin: _begin, end: _end});
        return true;
        
    }
    function getIndex(uint index) private returns(bytes32) {
        return sha3(index);
    }

    function remove(Data storage self, uint64 _index) internal returns(bool) {
        delete self.intervals[getIndex(_index)];
        return true;
    }

    
    
}