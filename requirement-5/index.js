/**
Extend String to support startsWith functionality that
returns true if the string starts with a given string
*/
String.prototype.startsWith = function(str) {
  if (typeof str != 'string') {
    return false;
  }
  if(this.indexOf(str)===0) {
    return true
  }
  return false
}

/**
Extend String to support endsWith functionality that
returns true if the string ends with a given string
this time, substring is used because string lengths may vary
*/
String.prototype.endsWith = function(str) {
  if (typeof str != 'string') {
    return false;
  }

  if(this.length < str.length) {
    return false;
  }

  if(str === '') { // edge case
    return true;
  }

  if(this.substring(this.length - str.length) === str) {
    return true;
  }
  return false
}
