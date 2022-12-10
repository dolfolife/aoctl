// from https://adventofcode.com/2021/day/3

//What is the life support rating of the submarine?
export const lifeSupportRating = async (): Promise<Number> => {

    const input = await Deno.readTextFile("./diagnosticreport.txt");
    let diagnosticReport: Array<string> = input.split("\n");
    let diagnosticReportCO2: Array<string> = input.split("\n");

    const oxygen = filterReportBy(diagnosticReport, getTheMostSignificantBit);
    const co2 = filterReportBy(diagnosticReportCO2, getTheLeastSignificantBit);

   return parseInt(oxygen, 2) * parseInt(co2, 2)
}

lifeSupportRating().then((position) => {
    console.log(position);
});

function filterReportBy(report: Array<string>, filterBy: Function): string {
    let position = 0;
    while(report.length > 1) {
        let condition = filterBy(report, position);
        report = report.filter(report => report.charAt(position) === condition);
        position++;
    }

    return report[0];
}

function getTheMostSignificantBit(diagnosticReport: string[], position: number) {
    let mostSignificantBits = mapBitInPosition(diagnosticReport, position);
    return mostSignificantBits.get("1")! >= mostSignificantBits.get("0")!? "1" : "0";
}

function getTheLeastSignificantBit(diagnosticReport: string[], position: number) {
    let leastSignificantBits = mapBitInPosition(diagnosticReport, position);
    return leastSignificantBits.get("1")! >= leastSignificantBits.get("0")!? "0" : "1";
}

function mapBitInPosition(diagnosticReport: string[], position: number) {
    let leastSignificantBits = new Map<string, number>([["1", 0], ["0", 0]]);
    for (let i = 0; i < diagnosticReport.length; i++) {
        leastSignificantBits.set(diagnosticReport[i].charAt(position), (leastSignificantBits.get(diagnosticReport[i].charAt(position)) || 0) + 1);
    }
    return leastSignificantBits;
}

