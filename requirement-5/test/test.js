var assert = require('assert');
describe('String', function() {
  describe('startsWith()', function() {
    it('should return false when the string doesnt start with a given prefix', function() {
      assert.equal(true, 'hang the dj'.startsWith('hang'));
      assert.equal(false, 'hang the dj'.startsWith('Hang'));
      assert.equal(false, 'hang the dj'.startsWith('I’ve got a room for rent'));
      assert.equal(true, 'hang the dj'.startsWith(''));
      assert.equal(true, 'hang the dj'.startsWith('hang the'));
      assert.equal(true, 'hang the dj'.startsWith('han'));
      assert.equal(true, 'hang the dj'.startsWith('hang t'));
      assert.equal(false, 'hang the dj'.startsWith(42));
      assert.equal(false, 'hang the dj'.startsWith({ first: 'Johnny' }));
    });
  });

  describe('endsWith()', function() {
    it('should return false when the string doesnt end with a given postfix', function() {
      assert.equal(true, 'hang the dj'.endsWith('dj'));
      assert.equal(false, 'hang the dj'.endsWith('panic on the streets'));
      assert.equal(true, 'hang the dj'.endsWith(''));
      assert.equal(true, 'hang the dj'.endsWith('the dj'));
      assert.equal(true, 'hang the dj'.endsWith('e dj'));
      assert.equal(true, 'hang the dj'.endsWith('j'));
      assert.equal(false, 'hang the dj'.endsWith(42));
      assert.equal(false, 'hang the dj'.endsWith({ first: 'Johnny' }));
    });
  });
});
