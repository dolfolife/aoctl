
export const lifeAfter256Days = async (): Promise<number> => {
    const input = await Deno.readTextFile("./input.txt");
    let fishLife = input.split(",").map((fishLife: string) => parseInt(fishLife, 10));

    let mapOfPossibleLife = new Map<number, number>([[0, 0], [1, 0], [2, 0], [3, 0], [4, 0], [5, 0], [6, 0], [7, 0], [8, 0]]);

    fishLife.forEach(((fish: number) => mapOfPossibleLife.set(fish, mapOfPossibleLife.get(fish)! + 1)));

    for (let i = 0; i < 256; i++) {
        let newMap = new Map<number, number>();
        for (let j = mapOfPossibleLife.size - 1; j >= 0; j--) {
            if (j > 0) {
                // reduce the count of days for the fish in this day
                newMap.set(j - 1, mapOfPossibleLife.get(j)!)
            } else {
                //new population
                newMap.set(8, mapOfPossibleLife.get(0)!);
                // new cycle of the fish that gave birth
                newMap.set(6, mapOfPossibleLife.get(0)! + newMap.get(6)!)
            }
        }
        mapOfPossibleLife = newMap;
    }

    let result = 0;
    mapOfPossibleLife.forEach((possibleLife) => result += possibleLife);
    return result;
}

lifeAfter256Days().then((count) => {
    console.log(`${count}`);
});