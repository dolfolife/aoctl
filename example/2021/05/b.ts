

//What is the power consumption of the submarine?
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

        getPoints(x1,y1,x2,y2).forEach(point => {
            let currentLines: number = linesCovered.get(point) || 0;
            if (linesCovered.get(point) === 1) dangerousPointCount++;
            linesCovered.set(point, currentLines +1);
        });
    })

    return dangerousPointCount;
}

liensOverlap().then((count) => {
    console.log(`${count}`);
});

function getPoints(x1: number,y1: number,x2: number,y2: number): Array<string> {
    var points = new Array<string>();
    const yDiff = y2 - y1;
    const xDiff = x2 - x1;
    const slope: number = 1.0 * ((yDiff) / (xDiff));

    const pointCount = getDistance(Math.abs(xDiff), Math.abs(yDiff));

    for (let i = 0; i < pointCount; i++) {
        let y = slope == 0 ? 0 : yDiff * (i / pointCount);
        let x = slope == 0 ? xDiff * (i / pointCount) : y / slope;
        points.push(`${Math.round(x) + x1},${Math.round(y) + y1}`);
    }

    points.push(`${x2},${y2}`);
    return points;
}

function getDistance(start: number , end: number): number {
    if (end == 0)
        return start;
    return getDistance(end, start % end);
}