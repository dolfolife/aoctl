import { Board } from "./Board.ts";

interface BingoGame {
    plays: number[]
    boards: Board[]
}

//What is the power consumption of the submarine?
export const worstBingoWinner = async (): Promise<number> => {
    const input = await Deno.readTextFile("./bingogame.txt");

    let game: BingoGame = buildGame(input.split("\n"));

    let play: number = 0;
    while(game.boards.length > 1) {
        game.boards = game.boards.filter(board => {
            board.markNumber(game.plays[play]);
            return board.wonWith === undefined;
        });
        play++;
    }

    while (game.boards[0].wonWith === undefined) {
        game.boards[0].markNumber(game.plays[play]);
        play++;
    }

    return game.boards[0].result();
}

worstBingoWinner().then((board) => {
    console.log(`${board}`);
});

function buildGame(input: string[]): BingoGame {

    let boards = new Array<Board>();
    for (let i = 2; i < input.length; i+=6) {
        let lines: number[][] = new Array<number[]>();
        for (let j = 0; j < 5; j++) {
            lines.push(input[j+i].split(" ").filter(s => s.length > 0).map(s => parseInt(s.trim(),10)));
        }

        boards.push(new Board(lines));
    }

    return {
        plays: input[0].split(",").filter(s => s.length > 0).map( s => parseInt(s.trim(), 10)),
        boards
    };
}

