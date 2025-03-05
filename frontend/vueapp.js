Vue.createApp({
    data() {
        return {
            greeting: "hello",
            showLoginModal: false,
            showP: false,
            showVerifyPassword: false,
            showLoginButton: true,
            showRegisterButton: false,
            showSignUpText: true,
            showLoginText: false,
            showBattlePage: false,
            showReplaysPage: false,
            showTutorialPage: true,
            currentReplay: {},
            frameToDisplay: false,
            emailinput: "",
        };
    },
    methods: {
        helloThere: function () {
            console.log("hello");
        },
        showLogin: function () {
            this.showLoginModal = true;
            console.log(this.showLoginModal);
        },
        signUp: function () {
            this.showVerifyPassword = true;
            this.showLoginButton = false;
            this.showRegisterButton = true;
            this.showLoginText = true;
            this.showSignUpText = false;
        },
        loginLayout: function () {
            this.showVerifyPassword = false;
            this.showLoginButton = true;
            this.showRegisterButton = false;
            this.showLoginText = false;
            this.showSignUpText = true;
        },
        register: async function () {
            console.log("Registering");
            var url = "http://localhost:8081/register"
            let jstring = JSON.stringify({msg: "HEY"})
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: jstring
            });
            if (response.ok) {
                let json = await response.json();
                console.log("Response:", json)
            } else {
                alert("HTTP-Error: ", + response.status);
            }

        },
        login: function () {
            console.log("Logging in");
        },
        getTextareaLines: function () {
            let text = document.querySelector(".textarea").value;
            let arr = text.split("\n");
            return arr
        },
        linesToObject: function () {
            // @TODO change to get actual user
            var user = "Hunter";
            let program = {user: user, instructions: []}
            let lines = this.getTextareaLines();

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
            return program;
        },
        submitProgram: async function () {
            let button = document.getElementById("sbutton");
            button.classList.add("is-loading");
            var url = "http://localhost:8081/battleprogram"
            let program = this.linesToObject();
            let jsonProgram = JSON.stringify(program);
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: jsonProgram
            });
            if (response.ok) {
                let json = await response.json();
                button.classList.remove("is-loading");
                console.log("Response:", json)
            } else {
                alert("HTTP-Error: " + response.status);
                button.classList.remove("is-loading");
            }
            button.classList.remove("is-loading");
        },
        showTutorial: function () {
            this.showBattlePage = false;
            this.showReplaysPage = false;
            this.showTutorialPage = true;

        },
        showBattle: function () {
            this.showBattlePage = true;
            this.showReplaysPage = false;
            this.showTutorialPage = false;
        },
        showReplays: function () {
            this.showBattlePage = false;
            this.showReplaysPage = true;
            this.showTutorialPage = false;
            this.drawCanvas();
        },
        runGame: async function () {
            let button = document.getElementById("gamebutton");
            button.classList.add("is-loading");
            var url = "http://localhost:8081/game"
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                button.classList.remove("is-loading");
                console.log("Response:", json)
            } else {
                alert("HTTP-Error: " + response.status);
                button.classList.remove("is-loading");
            }
            button.classList.remove("is-loading");

        },
        getReplay: async function () {
            let url = "http://localhost:8081/battlereplay" 
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                if (json.Frames === null) {
                    console.log("No games to grab");
                } else {
                    console.log(json);
                    this.currentReplay = json;
                }
            } else {
                alert("HTTP-Error: ", + response.status);
            }
        },
        startReplay: function () {
            var index = 0;
            var waitTime = 1000;
            let interval = setInterval(() => {
                console.log("Index is:", index)
                console.log("Every 1 Second a second passses...", index);
                this.frameToDisplay = this.currentReplay.Frames[index];
                index++;
                waitTime--;
                
                if (index > this.currentReplay.Frames.length) {
                    clearInterval(interval);
                    // -2 is the last frame to be displayed
                    this.frameToDisplay = this.currentReplay.Frames[index-2];
                    return;
                }
            }, waitTime);
        },
        drawCanvas: function () {
            setTimeout(() => {
                const canvas = document.getElementById("myCanvas");
                const ctx = canvas.getContext("2d");
                var offset = 2;
                var row = offset;
                size = 36;
                var color1 = "grey";
                var color2 = "grey";

                ctx.fillStyle = "black";
                ctx.fillRect(0, 0, 610, 610)

                // ctx.fillRect(2, 2, 36, 36)

                for (let j=0; j < 16; j++) {
                    var num = offset;
                    for (let i= 0; i < 8; i++) {
                        if (j % 2 == 0) {
                            ctx.fillStyle = color1;
                        } else {
                            ctx.fillStyle = color2;
                        }
                        ctx.fillRect(num, row, size, size);
                        if ( j % 2 == 0 ) {
                            ctx.fillStyle = color2;
                        } else {
                            ctx.fillStyle = color1;
                        }
                        ctx.fillRect(offset + num + size, row, size, size);
                        num = offset + offset + num + size*2;
                        console.log(i);
                    }
                    row = offset + row + size;
                }
                ctx.beginPath();
                ctx.arc(20, 20, 18, 0, 2 * Math.PI);
                // ctx.fill();
                ctx.stroke();

            }, 1);
        }
    },
    created: async function () {
        this.helloThere();
    },
    mounted: function () {
    }

}).mount("#app")
