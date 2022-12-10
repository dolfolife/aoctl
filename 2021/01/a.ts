// from https://adventofcode.com/2021/day/1

//How many measurements are larger than the previous measurement?
export const depthMeasurementIncreasesCount = async (): Promise<Number> => {

    const input = await Deno.readTextFile("./depthMeasurement.txt");
    const measurements = input.split("\n").map(line => parseInt(line, 10));

    let lastMeasurement: number | undefined = undefined;

    let count = 0;
    measurements.forEach( (measurement: number) => {
        if (lastMeasurement !== undefined && measurement > lastMeasurement) {
            count++;
        }
        lastMeasurement = measurement;
    })

    return count;
}

depthMeasurementIncreasesCount().then((count) => {
    console.log(count);
});
