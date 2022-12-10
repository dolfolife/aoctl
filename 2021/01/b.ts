// from https://adventofcode.com/2021/day/1

//How many measurements are larger than the previous measurement?
export const depthMeasurementWindowIncreasesCount = async (): Promise<Number> => {

    const input = await Deno.readTextFile("./depthMeasurement.txt");
    const measurements = input.split("\n").map(line => parseInt(line, 10));

    let count = 0;
    for(let i = 0; i < measurements.length; i++) {
        if (measurements[i] < measurements[i + 3] ) {
            count++;
        }
    }

    return count;
}

depthMeasurementWindowIncreasesCount().then((count) => {
    console.log(count);
});
