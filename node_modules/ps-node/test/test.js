var PS = require('../index');
var CP = require('child_process');
var assert = require('assert');
var Path = require('path');

var serverPath = Path.resolve(__dirname, './node_process_for_test.js');
var UpperCaseArg = '--UPPER_CASE';
var child = null;
var pid = null;

function startProcess() {
  child = CP.fork(serverPath, [UpperCaseArg]);
  pid = child.pid;
}

describe('test', function () {
  before(function (done) {
    PS.lookup({arguments: 'node_process_for_test'}, function (err, list) {
      var processLen = list.length;
      var killedCount = 0;
      if (processLen) {
        list.forEach(function (item) {
          PS.kill(item.pid, function () {
            killedCount++;
            if (killedCount === processLen) {
              startProcess();
              done();
            }
          });
        });
      } else {
        startProcess();
        done();
      }
    });
  });

  describe('#lookup()', function () {

    it('by id', function (done) {
      PS.lookup({pid: String(pid)}, function (err, list) {
        assert.equal(list.length, 1);
        assert.equal(list[0].arguments[0], serverPath);

        done();
      });
    });

    it('by command & arguments', function (done) {
      PS.lookup({command: '.*(node|iojs).*', arguments: 'node_process_for_test'}, function (err, list) {
        assert.equal(list.length, 1);
        assert.equal(list[0].pid, pid);
        assert.equal(list[0].arguments[0], serverPath);
        done();
      });
    });

    it('by arguments, the matching should be case insensitive ', function (done) {
      PS.lookup({arguments: 'UPPER_CASE'}, function (err, list) {
        assert.equal(list.length, 1);
        assert.equal(list[0].pid, pid);
        assert.equal(list[0].arguments[0], serverPath);

        PS.lookup({arguments: 'upper_case'}, function (err, list) {
          assert.equal(list.length, 1);
          assert.equal(list[0].pid, pid);
          assert.equal(list[0].arguments[0], serverPath);
          done();
        });
      });
    });

    it('empty result list should be safe ', function (done) {
      PS.lookup({command: 'NOT_EXIST', psargs: 'l'}, function (err, list) {
        assert.equal(list.length, 0);
        done();
      });
    });

    it('should work correctly with options `aux`', function (done) {
      PS.lookup({command: 'node', psargs: 'aux'}, function (err, list) {
        assert.equal(list.length > 0, true);
        list.forEach(function (row) {
          assert.equal(/^\d+$/.test(row.pid), true);
        });
        done();
      });
    });
  });

  describe('#kill()', function () {

    it('kill', function (done) {

      PS.kill(pid, function (err) {
        assert.equal(err, null);
        PS.lookup({pid: String(pid)}, function (err, list) {
          assert.equal(list.length, 0);
          done();
        });
      });
    });

    it('should force kill when opts.signal is 9', function (done) {
      startProcess();

      PS.kill(pid, {signal: 9}, function (err) {
        assert.equal(err, null);
        PS.lookup({pid: String(pid)}, function (err, list) {
          assert.equal(list.length, 0);
          done();
        });
      });
    });

    it('should throw error when opts.signal is invalid', function (done) {
      startProcess();
      PS.kill(pid, {signal: 'INVALID'}, function (err) {
        assert.notEqual(err, null);
        PS.kill(pid, function(){
            done();
        });
      });
    });

    it('should not throw an exception if the callback is undefined', function (done) {
      assert.doesNotThrow(function () {
        PS.kill(pid);
        setTimeout(done, 400);
      });
    });
  });
});
