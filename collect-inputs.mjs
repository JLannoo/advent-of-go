import fs from "node:fs";

// Read year directories
const years = fs.readdirSync("./").filter((file) => /^\d{4}$/.test(file));

// For each year, read day directories
const yearDayMap = {};

for (const year of years) {
    const days = fs.readdirSync(year).filter((file) => /^day\d+$/.test(file));
    yearDayMap[year] = days;
}

// For each day, fetch input files from AoC if they don't exist
for(const [year, days] of Object.entries(yearDayMap)) {
    for(const day of days) {
        if(!shortInputExists(year, day)) {
            console.log(`Downloading short input for ${year} ${day}`)
            const shortInput = await downloadShortInputForDay(year, day);
            fs.writeFileSync(getShortInputPath(year, day), shortInput);
        }

        if(!inputExists(year, day)) {
            console.log(`Downloading input for ${year} ${day}`)
            const input = await downloadInputForDay(year, day);
            fs.writeFileSync(getInputPath(year, day), input);
        }
    }
}

function shortInputExists(year, day) {
    return fs.existsSync(getShortInputPath(year, day));
}

function inputExists(year, day) {
    return fs.existsSync(getInputPath(year, day));
}

function getShortInputPath(year, day) {
    const FILE_NAME = "short-input.txt";
    return `./${year}/${day}/${FILE_NAME}`;
}

function getInputPath(year, day) {
    const FILE_NAME = "input.txt";
    return `./${year}/${day}/${FILE_NAME}`;
}

async function downloadInputForDay(year, day) {
    const URL = `https://adventofcode.com/${year}/day/${day.slice(3)}/input`;
    const response = await fetch(URL, {
        headers: {
            cookie: `session=${process.env.AOC_SESSION_COOKIE}`
        }
    })
    const text = await response.text();
    
    return text.trim();
}

async function downloadShortInputForDay(year,day) {
    const URL = `https://adventofcode.com/${year}/day/${day.slice(3)}`;
    const response = await fetch(URL, {
        headers: {
            cookie: `session=${process.env.AOC_SESSION_COOKIE}`
        }
    })

    const text = await response.text();

    const input = text.match(/<pre><code>([\s\S]*?)<\/code><\/pre>/)[1];
    
    
    return input.trim();
}