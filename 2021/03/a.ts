// from https://adventofcode.com/2021/day/3

//What is the power consumption of the submarine?
export const powerConsumption = async (): Promise<Number> => {

    const input = await Deno.readTextFile("./diagnosticreport.txt");
    let mostSignificantBits: Array<Map<string, number>> | undefined
    input
    .split("\n")
    .forEach((line: string) => {
        let bits = line.split('');
        if (mostSignificantBits === undefined) mostSignificantBits = new Array(bits.length);
        for (let i = 0; i < bits.length; i++) {
            if (mostSignificantBits[i] == undefined) mostSignificantBits[i] = new Map<string, number>([["1", 0], ["0", 0]]);
            mostSignificantBits[i].set(bits[i], (mostSignificantBits[i].get(bits[i])|| 0)+1);
        }
    });
    return parseInt(mostSignificantBits!.map((value: Map<string, number>) => value.get('1')! > value.get('0')! ? '1': '0').join(''), 2)
    * parseInt(mostSignificantBits!.map((value: Map<string, number>) => value.get('1')! < value.get('0')! ? '1': '0').join(''), 2);
}

powerConsumption().then((position) => {
    console.log(position);
});
