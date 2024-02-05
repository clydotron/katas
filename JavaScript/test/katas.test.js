
import { assert, config } from 'chai'
import { toNato, dblLinear } from './../katas.js'

function testing(actual, expected) {
	Test.assertEquals(actual, expected)
}
describe('Tests', function () {
    const { strictEqual } = assert;

    config.truncateThreshold = 0;

    function doTest(input, expected) {
        const actual = toNato(input);
        strictEqual(actual, expected, `for string = "${input}"\n`);
    }

    it('Sample Tests', function () {
        doTest('If you can read',
            'India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta'
        );
        doTest('Did not see that coming',
            'Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf'
        );
        doTest('go for it!', 'Golf Oscar Foxtrot Oscar Romeo India Tango !');
    });
});

describe("dblLinear",function() {
it("Basic tests",function() { 
	assert.equal(dblLinear(10), 22)
    assert.equal(dblLinear(20), 57)
    assert.equal(dblLinear(30), 91)
    assert.equal(dblLinear(50), 175)
    assert.equal(dblLinear(100), 447)
})})