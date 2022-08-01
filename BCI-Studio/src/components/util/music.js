export function rand(max) {
    return Math.floor(Math.random() * max)
}

export function chance(pct) {
    return rand(100) <= pct
}

export function numToNote(num) {
    let note = ""
    if(num % 12 == 0) {
        note = "C"
    } else if(num % 12 == 1) {
        note = 'C#'
    } else if(num % 12 == 2) {
        note = 'D'
    } else if(num % 12 == 3) {
        note = 'D#'
    } else if(num % 12 == 4) {
        note = 'E'
    } else if(num % 12 == 5) {
        note = 'F'
    } else if(num % 12 == 6) {
        note = 'F#'
    } else if(num % 12 == 7) {
        note = 'G'
    } else if(num % 12 == 8) {
        note = 'G#'
    } else if(num % 12 == 9) {
        note = 'A'
    } else if(num % 12 == 10) {
        note = 'A#'
    } else if(num % 12 == 11) {
        note = 'B'
    }

    let octave = Math.floor(num / 12) - 1

    return note + octave
}

export function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}