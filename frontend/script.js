

main();

var url = "http://localhost:8081/json"

function getTextareaLines() {
    let text = document.querySelector(".textarea").value;
    let arr = text.split("\n");
    return arr
}

async function main() {

    // Need to get a cookie or somethign for the appropriate user?
    let user = "Hunter"

    let submitbutton = document.querySelector("#submitbutton");
    submitbutton.addEventListener("click", submitProgram);
    async function submitProgram(event) {
        let program = {user: user, instructions: []}
        let lines = getTextareaLines();

        for (let i = 0; i < lines.length; i++) {
            let standbyobj = {instruction: "", args: []}
            let line = lines[i].split(" ");
            standbyobj.instruction = line[0];
            for (let j = 1; j < line.length; j++) {
                let copy = line[j]
                if (!isNaN(Number(copy))) {
                    standbyobj.args.push(Number(line[j]))
                } else {
                    standbyobj.args.push(line[j])
                }
            }
            program.instructions.push(standbyobj);
        }
        let jsonProgram = JSON.stringify(program);
        let response = await fetch(url);
        if (response.ok) {
            let json = await response.json();
            console.log("Response:", json)
        } else {
            alert("HTTP-Error: " + response.status);
        }
    }
}
