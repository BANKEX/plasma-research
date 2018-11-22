pragma solidity ^0.4.24;

/**
    @title OrderedIntervalList
    @dev List of ordered by intervals with non intersections checks. 
 */
library OrderedIntervalList {
    struct Interval {
        uint64 begin; // inclusive
        uint64 end; // exclusive

        uint64 next;  
        uint64 previous;
    }
    struct Data {
        //Interval[] intervals;
        mapping (uint64 => Interval) intervals;
        uint64 first;
        uint64 index;
    }

  

    /**
    @notice Get interval by list index
    @param id interval index in list
    @return interval tuple
     */
    function get(Data storage self, uint64 id) internal view returns(Interval storage interval) {
        require(id <= self.index, "interval id doesn't exists in interval set");        
        interval = self.intervals[id];    
        require(interval.end != 0, "interval id doesn't exsits in interval set");    
        return interval;
    }  

    

    /**
    @notice Insert interval by specific index in list
    @dev Method also check that new interval doesn't intersect with existed intervals in list
    @param prev previous index
    @param next next index
    @param begin left bound of interval (inclusive)
    @param end right bound of interval (exclusive)        
     */
    function insert(Data storage self, uint64 prev, uint64 next, uint64 begin, uint64 end) internal returns(uint64 id) {
        require(begin < end, "right bound less than left bound");              
       
        bool concatLeft = false;
        bool concatRight = false;
        id = self.index + 1;
        self.index = self.index + 1;
        
        if (self.first == 0) { 
            prev = 0;
            next = 0;
            self.first = id;
        } else {
            require(prev > 0 || next > 0);
            Interval storage prevInterval = self.intervals[prev];   
            Interval storage nextInterval = self.intervals[next];        
           
            if (prev > 0 && next > 0) {
                require(prevInterval.end != 0, "previous interval doesn't exists");     
                require(nextInterval.end != 0, "next interval doesn't exists");       
                require(
                    prevInterval.end <= begin && nextInterval.begin >= end
                );
                if (prevInterval.end == begin) {
                    concatLeft = true;               
                    prevInterval.end = end;
                }
                if (nextInterval.begin == end) {
                    concatRight = true;
                    nextInterval.begin = begin;
                }
                if (! concatLeft && !concatRight ) {                

                    nextInterval.previous = id;
                    prevInterval.next = id;                  
                }
                if (concatLeft && concatRight) {
                    prevInterval.end = nextInterval.end;
                    prevInterval.next = nextInterval.next;                  
                    delete self.intervals[next];
                }
            } else if (prev > 0 && next == 0) {
                require(prevInterval.end != 0, "previous interval doesn't exists");     
                require(prevInterval.next == 0, "previous element is not last element");
                require(
                    prevInterval.end <= begin
                );

                if (prevInterval.end == begin) {
                    concatLeft = true;               
                    prevInterval.end = end;
                } else {
                    prevInterval.next = id;
                }
            } else if (prev == 0 && next >= 0) {
                require(nextInterval.end != 0, "next interval doesn't exists");     
                require(nextInterval.previous == 0, "next element is not first element");
                require(
                    nextInterval.begin >= end
                );
                if (nextInterval.begin == end) {
                    concatRight = true;
                    nextInterval.begin = begin;
                } else {
                    self.first = id;
                    nextInterval.previous = id;
                }
                
            }
        }
 
        if (! concatRight && ! concatLeft) {
            self.index = id;
            self.intervals[id] = Interval({begin: begin, end: end, previous: prev, next: next});   
        } else {
            
            if (concatLeft) {                
                id = prev;
            } else{          
                if (concatRight) {      
                    id = next;
                }
            }            
        }        
      
    } 

    /**
        @notice Remove interval by index
        @param index interval index in list        
        @param begin left bound
        @param end right bound
     */
    function remove(Data storage self, uint64 index, uint64 begin, uint64 end) internal returns (uint64 newInterval) {
        require(begin < end, "right bound less than left bound");        
        require(index <= self.index, "valid index bounds");
        

        Interval storage modifiedInterval = self.intervals[index];        
        require(modifiedInterval.end != 0, "removed interval doesn't exists");
        require(modifiedInterval.begin <= begin && modifiedInterval.end >= end, "incorrect removed range bounds");

        if (begin > modifiedInterval.begin ) {

            uint64 oldEnd = modifiedInterval.end;
            modifiedInterval.end = begin;
            if (end < oldEnd) {
                newInterval = insert(self, index, modifiedInterval.next, end, oldEnd);
                modifiedInterval.next = newInterval;
            } 

        } else {
            modifiedInterval.begin = end;
        }

        if (modifiedInterval.begin == modifiedInterval.end) {         
            if (modifiedInterval.previous > 0) {
                Interval storage prevInterval = self.intervals[modifiedInterval.previous];   
                prevInterval.next = modifiedInterval.next;
            }
            if (modifiedInterval.next > 0) {
                Interval storage nextInterval = self.intervals[modifiedInterval.next];                
                nextInterval.previous = modifiedInterval.previous;
            }
            if (modifiedInterval.previous == 0) {
                self.first = modifiedInterval.next;
            }  
            
            delete self.intervals[index];
        }


       

    }    
    
}