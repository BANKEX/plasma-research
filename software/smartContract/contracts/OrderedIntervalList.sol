pragma solidity ^0.4.24;


/**
 * @title OrderedIntervalList
 * @dev List of ordered by intervals with non intersections checks.
*/
library OrderedIntervalList {
  struct Interval {
    uint256 begin; // inclusive
    uint256 end; // exclusive

    uint256 next;
    uint256 prev;
  }

  struct Data {
    Interval[] intervals; // sparsed array
    uint firstIndex;
    uint lastIndex;
  }

  function getFirstIndex(Data storage self) internal view returns(uint) {
    return self.firstIndex;
  }

  function getLastIndex(Data storage self) internal view returns(uint) {
    return self.lastIndex;
  }

  /**
   * @notice Get interval by the index
   * @param id interval index in the list
   * @return interval tuple
   */
  function get(Data storage self, uint id) internal view returns(Interval storage interval) {
    require(id <= self.intervals.length, "interval id doesn't exists in interval set");
    interval = self.intervals[id];
    require(interval.end != 0, "interval id doesn't exsits in interval set");
  }

  /**
   * @notice Insert interval in the specific place in a list
   * @dev Method also check that new interval doesn't intersect with existed intervals in list
   * @param prev id of the prev interval in the list. Zero if it's a first interval.
   * @param next id of the next interval in the list. Zero if it's a last interval.
   * @param begin left bound of the new interval (inclusive)
   * @param end right bound of the new interval (exclusive)
   * @return id of the interval that contain new interval. Could be a new interval or an existed with
   * extended bounds in case of adjacent bounds of the inserted interval with his neighbors.
   */
  function insert(Data storage self, uint prev, uint next, uint begin, uint end) internal returns(uint id) {
    require(begin < end, "right bound less or equal to left bound");

    bool concatLeft = false;
    bool concatRight = false;
    id = self.intervals.length;
    if (id == 0) {
      self.intervals.length += 1;
      id += 1;
    }

    if (self.firstIndex == 0) {
      //  require(prev == 0 && next == 0, "insert in empty list with non zero link to neighbor elements");
      self.firstIndex = id;
      self.lastIndex = id;
    } else {
      require(prev > 0 || next > 0, "previous and next element doesn't exists");
      Interval storage prevInterval = self.intervals[prev];
      Interval storage nextInterval = self.intervals[next];

      if (prev > 0 && next > 0) { // insert between to existed intervals
        require(prevInterval.end != 0, "previous interval doesn't exists");
        require(nextInterval.end != 0, "next interval doesn't exists");
        require(prevInterval.end <= begin && nextInterval.begin >= end, "new interval out of bounds of neighbors intervals");
        if (prevInterval.end == begin) {
          concatLeft = true;
          prevInterval.end = end;
        }
        if (nextInterval.begin == end) {
          concatRight = true;
          nextInterval.begin = begin;
        }
        if (! concatLeft && !concatRight ) {

          nextInterval.prev = id;
          prevInterval.next = id;
        }
        if (concatLeft && concatRight) {
          prevInterval.end = nextInterval.end;
          prevInterval.next = nextInterval.next;
          delete self.intervals[next];
        }
      } else if (prev > 0 && next == 0) { // insert as last elemnt
        require(prevInterval.end != 0, "prev interval doesn't exists");
        require(prevInterval.next == 0, "prev element is not last element");
        require(prevInterval.end <= begin, "new interval out of bounds of prev interval");

        if (prevInterval.end == begin) {
          concatLeft = true;
          prevInterval.end = end;
        } else {
          self.lastIndex = id;
          prevInterval.next = id;
        }
      } else if (prev == 0 && next >= 0) { // insert as first element
        require(nextInterval.end != 0, "next interval doesn't exists");
        require(nextInterval.prev == 0, "next element is not first element");
        require(nextInterval.begin >= end, "new interval out of bounds of next interval");

        if (nextInterval.begin == end) {
          concatRight = true;
          nextInterval.begin = begin;
        } else {
          self.firstIndex = id;
          nextInterval.prev = id;
        }
      }
    }

    if (!concatRight && !concatLeft) {
      self.intervals.push(Interval({
        begin: begin,
        end: end,
        prev: prev,
        next: next
      }));
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
   * @notice Remove range in interval by index
   * @param index interval index in list
   * @param begin left range bound
   * @param end right range bound
   * @return index of the new interval if new one was created (was made a hole insided existed interval) or zero.
   */
  function remove(Data storage self, uint index, uint begin, uint end) internal returns (uint newInterval) {
    require(begin < end, "right bound less than left bound");
    require(index <= self.intervals.length, "valid index bounds");


    Interval storage modifiedInterval = self.intervals[index];
    require(modifiedInterval.end != 0, "removed interval doesn't exists");
    require(modifiedInterval.begin <= begin && modifiedInterval.end >= end, "incorrect removed range bounds");

    if (begin > modifiedInterval.begin ) {

      uint oldEnd = modifiedInterval.end;
      modifiedInterval.end = begin;
      if (end < oldEnd) {
        newInterval = insert(self, index, modifiedInterval.next, end, oldEnd);
        modifiedInterval.next = newInterval;
      }
    } else {
      modifiedInterval.begin = end;
    }

    if (modifiedInterval.begin == modifiedInterval.end) {
      if (modifiedInterval.prev > 0) {
        Interval storage prevInterval = self.intervals[modifiedInterval.prev];
        prevInterval.next = modifiedInterval.next;
      }
      if (modifiedInterval.next > 0) {
        Interval storage nextInterval = self.intervals[modifiedInterval.next];
        nextInterval.prev = modifiedInterval.prev;
      }
      if (modifiedInterval.prev == 0) {
        self.firstIndex = modifiedInterval.next;
      }
      if (modifiedInterval.next == 0) {
        self.lastIndex = modifiedInterval.next;
      }
      delete self.intervals[index];
    }
  }
}
