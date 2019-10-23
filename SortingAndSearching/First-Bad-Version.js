
/* 
  这道题是二分查找思想的运用

  比较蛋疼的是，居然不支持Go语言的判定
*/

/**
 * Definition for isBadVersion()
 * 
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 */

// var isBadVersion = function(version) {
// }


/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function(isBadVersion) {

  /**
   * @param {integer} n Total versions
   * @return {integer} The first bad version
   */
  return function(n) {
    let l = 1, r = n;
    let m = parseInt((l + r) / 2);
    let res, resR;
    while (m >= l) {
      res = isBadVersion(m);
      resR = isBadVersion(m + 1);
      if ((!res && resR) || m == l) {
        break
      }
      if (res) {
        r = m
      } else {
        l = m
      }
      m = parseInt((l + r) / 2);
    }
    if (res) {
      return m;
    }
    return m + 1;
  };
};


console.log(solution(function(version) {
  if (version >= 4) {
    return true;
  }
  return false;
})(5));

console.log(solution(function(version) {
  if (version >= 3) {
    return true;
  }
  return false;
})(3));

console.log(solution(function(version) {
  if (version >= 2) {
    return true;
  }
  return false;
})(2));

console.log(solution(function(version) {
  if (version >= 1702766719) {
    return true;
  }
  return false;
})(2126753390));