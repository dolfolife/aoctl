
export const liensOverlap = async (): Promise<number> => {
    const input = await Deno.readTextFile("./input.txt");

    const  linesCovered: Map<string, number> = new Map<string, number>();
    let min: number, max: number;
    let dangerousPointCount = 0;
    input.split("\n").forEach((line: string) => {
        const [x1, y1, x2, y2] = line
        .split(" -> ")
        .flatMap( (coordinate: string) => coordinate.split(","))
        .map(point => parseInt(point, 10));

        switch(true) {
            case x1 === x2:
                [min, max] = getLower(y1,y2)
                while(min <= max) {
                    let key = `${x1},${min}`;
                    let currentLines: number = linesCovered.get(key) || 0;
                    if (linesCovered.get(key) === 1) dangerousPointCount++;
                    linesCovered.set(key, currentLines +1);
                    min++;
                }
                break;
            case y1 === y2:
                [min, max] = getLower(x1,x2)
                while(min <= max) {
                    let key = `${min},${y1}`;
                    let currentLines: number = linesCovered.get(key) || 0;
                    if (linesCovered.get(key) === 1) dangerousPointCount++;
                    linesCovered.set(key, currentLines +1);
                    min++;
                }
                break;
            default:
                break;
        }
    })

    return dangerousPointCount;
}

liensOverlap().then((count) => {
    console.log(`${count}`);
});

function getLower(y1: number, y2: number): [number, number] {
    return y1 < y2 ? [y1,y2] : [y2,y1]
}