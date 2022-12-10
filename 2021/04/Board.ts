// from https://adventofcode.com/2021/day/4
export class Board {
    board: number[][];
    markedNumbers: Set<number>;
    wonWith: number | undefined;

    constructor(board: number[][]) {
        this.markedNumbers = new Set<number>();
        this.board = board;
    }

    markNumber(number: number) {
        this.markedNumbers.add(number);
        if (this.won())
            this.wonWith = number;
    }

    won(): boolean {
        let result = false;

        for (let row = 0; row < 5; row++) {
            let won = true;
            for (let col = 0; col < 5; col++) {
                if (!this.markedNumbers.has(this.board[row][col])) {
                    won = false;
                }
            }
            if (won) {
                result = true;
                break;
            }
        }

        if (result === false) {
            for (let col = 0; col < 5; col++) {
                let won = true;
                for (let j = 0; j < 5; j++) {
                    if (!this.markedNumbers.has(this.board[j][col])) {
                        won = false;
                    }
                }
                if (won) {
                    result = true;
                    break;
                }
            }
        }

        return result;
    }

    result(): number {
        let sum = 0;
        for (let col = 0; col < 5; col++) {
            for (let j = 0; j < 5; j++) {
                if (!this.markedNumbers.has(this.board[j][col])) {
                    sum += this.board[j][col];
                }
            }
        }

        return sum * this.wonWith!;
    }
}
