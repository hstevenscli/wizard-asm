
main();

function getTextareaLines() {
    let text = document.querySelector(".textarea").value;
    let arr = text.split("\n");
    return arr
}

async function main() {

    let submitbutton = document.querySelector("#submitbutton");
    submitbutton.addEventListener("click", submit);
    function submit(event) {
        console.log("submit");

        let lines = getTextareaLines();
        console.log(lines);
    }
}
