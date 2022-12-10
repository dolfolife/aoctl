
export const lifeAfter80Days = async (): Promise<number> => {
    const input = await Deno.readTextFile("./input.txt");
    let fishLife = input.split(",").map((fishLife: string) => parseInt(fishLife, 10));

    for (let i = 0; i < 80; i++) {
        let newFish: Array<number> = new Array<number>();
        fishLife = fishLife.map((life: number) => {
            if (life === 0) {
                newFish.push(8);
                return 6;
            }
            return life - 1;
        });

        fishLife = fishLife.concat(newFish);
    }

    return fishLife.length;
}

lifeAfter80Days().then((count) => {
    console.log(`${count}`);
});