pragma solidity ^0.4.24;


/**
 * @title OrderedIntervalList
 * @dev List of ordered by intervals with non intersections checks.
*/
library OrderedIntervalList {
  struct Interval {
    uint64 begin; // inclusive
    uint64 end;   // exclusive

    uint64 next;
    uint64 prev;
  }

  struct Data {
    Interval[] intervals; // sparsed array
    uint64 firstIndex;
    uint64 lastIndex;
  }

  function isInitialized(Data storage self) internal view returns(bool) {
    return self.intervals.length > 0;
  }

  function getFirstIndex(Data storage self) internal view returns(uint64) {
    return self.firstIndex;
  }

  function getLastIndex(Data storage self) internal view returns(uint64) {
    return self.lastIndex;
  }

  /**
   * @notice Check if OrderedIntervalList initialized
   * @return was initialized or not
   */
  function initialize(Data storage self) internal {
    require(self.intervals.length == 0, "OrderedIntervalList was already initialized");
    self.intervals.push(Interval(0,0,0,0));
  }

  /**
   * @notice Get interval by the index
   * @param id interval index in the list
   * @return interval tuple
   */
  function get(Data storage self, uint64 id) internal view returns(Interval storage interval) {
    require(id < self.intervals.length, "interval id doesn't exists in interval set");
    interval = self.intervals[id];
    //require(interval.end != 0, "interval id doesn't exsits in interval set");
  }

  /**
   * @notice Check interval existance by the index
   * @param id interval index in the list
   * @return is existing or not
   */
  function exist(Data storage self, uint64 id) internal view returns (bool) {
    return self.intervals[id].end != 0;
  }

  /**
   * @notice Add interval after the lates interval
   * @param size length of the new interval
   * @return id of the latest interval
   */
  function append(
    Data storage self,
    uint64 size
  )
    internal
    returns(
      uint64 id,
      uint64 begin,
      uint64 end
    )
  {
    Interval storage lastInterval = self.intervals[self.lastIndex];
    begin = lastInterval.end;
    end = lastInterval.end + size - 1;
    id = insert(self, self.lastIndex, 0, begin, end);
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
  function insert(
    Data storage self,
    uint64 prev,
    uint64 next,
    uint64 begin,
    uint64 end
  )
    internal
    returns(uint64 id)
  {
    return _insert(
      self,
      prev,
      next,
      begin,
      end,
      false
    );
  }

  /**
   * @notice Remove range in interval by index
   * @param index interval index in list
   * @param begin left range bound
   * @param end right range bound
   * @return index of the new interval if new one was created (was made a hole insided existed interval) or zero.
   */
  function remove(
    Data storage self,
    uint64 index,
    uint64 begin,
    uint64 end
  )
    internal
    returns(uint64 newInterval)
  {
    require(begin < end, "right bound less than left bound");
    require(index < self.intervals.length, "valid index bounds");

    Interval storage modifiedInterval = self.intervals[index];
    Interval storage prevInterval = self.intervals[modifiedInterval.prev];
    Interval storage nextInterval = self.intervals[modifiedInterval.next];
    require(modifiedInterval.end != 0, "removed interval doesn't exists");
    require(modifiedInterval.begin <= begin && end <= modifiedInterval.end, "incorrect removed range bounds");

    bool shrinkBegin = (begin == modifiedInterval.begin);
    bool shrinkEnd = (end == modifiedInterval.end);

    if (shrinkBegin && shrinkEnd) {
      // Remove whole interval

      if (modifiedInterval.prev > 0) {
        prevInterval.next = modifiedInterval.next;
      } else {
        self.firstIndex = modifiedInterval.next;
      }

      if (modifiedInterval.next > 0) {
        nextInterval.prev = modifiedInterval.prev;
      } else {
        self.lastIndex = modifiedInterval.prev;
      }

      delete self.intervals[index];
    } else
    if (shrinkBegin) {
      // Shrink from left side
      modifiedInterval.begin = end;
    } else
    if (shrinkEnd) {
      // Shrink from right side
      modifiedInterval.end = begin;
    } else {
      // Make a hole
      uint64 oldEnd = modifiedInterval.end;
      modifiedInterval.end = begin;
      modifiedInterval.next = _insert(
        self,
        index,
        modifiedInterval.next,
        end,
        oldEnd,
        true
      );
      newInterval = modifiedInterval.next;
    }
  }

  function _insert(
    Data storage self,
    uint64 prev,
    uint64 next,
    uint64 begin,
    uint64 end,
    bool allowGapAfterLast
  )
    private
    returns(uint64 id)
  {
    require(begin < end, "right bound less or equal to left bound");
    require((prev != 0 || next != 0) == (self.firstIndex > 0), "prev and next could be zero iff no intervals");

    if (!isInitialized(self)) {
      initialize(self);
    }

    Interval storage prevInterval = self.intervals[prev];
    Interval storage nextInterval = self.intervals[next];

    require(prev == 0 || begin >= prevInterval.end, "begin could not intersect prev interval");
    require(next == 0 || end <= nextInterval.begin, "end could not intersect next interval");

    if ((prev > 0) == (next > 0)) {
      // Adding between existing intervals or very first interval
      require(
        prevInterval.next == next && nextInterval.prev == prev,
        "prev and next should refer to the neighboring intervals"
      );
    } else
    if (next > 0) {
      // Adding before first interval
      require(
        self.firstIndex == next && nextInterval.prev == 0,
        "next should refer to the first interval"
      );
    } else
    if (prev > 0) {
      // Adding after last interval
      require(
        self.lastIndex == prev && prevInterval.next == 0,
        "prev should refer to the last interval"
      );
      require(
        allowGapAfterLast || prev != self.lastIndex || prevInterval.end == begin, 
        "should begin from the end of latest interval when adding to the end"
      );
    }

    bool concatPrev = (prev > 0 && begin == prevInterval.end);
    bool concatNext = (next > 0 && end == nextInterval.begin);

    if (!concatPrev && !concatNext) {
      id = uint64(self.intervals.length);
      self.intervals.push(Interval({
        begin: begin,
        end: end,
        prev: prev,
        next: next
      }));

      if (next > 0) {
        nextInterval.prev = id;
      } else {
        self.lastIndex = id;
      }

      if (prev > 0) {
        prevInterval.next = id;
      } else {
        self.firstIndex = id;
      }
    } else
    if (concatPrev && concatNext) {
      prevInterval.end = nextInterval.end;
      prevInterval.next = nextInterval.next;
      id = prev;

      // When attaching pre last to last
      if (next == self.lastIndex) {
        self.lastIndex = id;
      } else {
        self.intervals[nextInterval.next].prev = id;
      }

      delete self.intervals[next];
    } else
    if (concatPrev) {
      prevInterval.end = end;
      id = prev;
    } else
    if (concatNext) {
      nextInterval.begin = begin;
      id = next;
    }
  }
}
