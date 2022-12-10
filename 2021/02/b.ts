// from https://adventofcode.com/2021/day/2

//What do you get if you multiply your final horizontal position by your final depth?
export const shipFinalPosition = async (): Promise<Number> => {

    const input = await Deno.readTextFile("./shipmovement.txt");
    let shipCoordinates = {
        x: 0,
        y: 0,
        aim: 0
    }
    const measurements = input
    .split("\n")
    .forEach((line: string) => {
        let [direction, d] = line.split(" ");
        let distance = parseInt(d, 10);
        switch (direction) {
            case "forward":
                shipCoordinates.x += distance
                shipCoordinates.y += distance*shipCoordinates.aim;
                break;
            case "up":
                shipCoordinates.aim -= distance
                break;
            case "down":
                shipCoordinates.aim += distance
                break;
            default:
                console.error("this is not a valid direction: " + direction);
        }
    });

    return shipCoordinates.x * shipCoordinates.y
}

shipFinalPosition().then((position) => {
    console.log(position);
});
