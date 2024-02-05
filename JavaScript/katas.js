

// <6 kyu> If you can read this...
// https://www.codewars.com/kata/586538146b56991861000293/train/javascript
var NATO = {
	A: "Alfa",
	B: "Bravo",
	C: "Charlie",
	D: "Delta",
	E: "Echo",
	F: "Foxtrot",
	G: "Golf",
	H: "Hotel",
	I: "India",
	J: "Juliet",
	K: "Kilo",
	L: "Lima",
	M: "Mike",
	N: "November",
	O: "Oscar",
	P: "Papa",
	Q: "Quebec",
	R: "Romeo",
	S: "Sierra",
	T: "Tango",
	U: "Uniform",
	V: "Victor",
	W: "Whiskey",
	X: "Xray",
	Y: "Yankee",
	Z: "Zulu",
}

export function toNato(words) {
    return words
        .toUpperCase()
        .split('')
        .filter(v => v !== ' ')
        .map(k => k in NATO ? NATO[k] : k)
        .join(' ')
}

export function toNato2(words) {
    return words
        .toUpperCase()
        .split('')
        .filter(c => c !== ' ')
        .map(c => NATO[c] || c)
        .join(' ')
}

// <4 kyu> Twice Linear
// https://www.codewars.com/kata/5672682212c8ecf83e000050/train/javascript
export function dblLinear(n) {
	const u = [1]
	var i = 0
	var j = 0
	while ( u.length <= n ) {
		var n1 = u[i] * 2 + 1
		var n2 = u[j] * 3 + 1

		if ( n1 < n2 ) {
			i++
		} else if( n1 > n2) {
			j++
		} else {
			i++
			j++
		}
		u.push( n1 <= n2 ? n1 : n2)
	}
	return u[n]
}