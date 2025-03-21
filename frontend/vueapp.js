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
            usernameInput: "",
            passwordInput: "",
            passwordConfirmationInput: "",
            disabledButton: true,
            errors: {},
            hamburgerEnabled: false,
            whoami: "",
            duelUserInput: "",
            ploc: {},
            tempProgram: [""],
            premiumSpells: {},
            showReportModal: false,
            bugReportMessage: "",
            bugEmail: "",
            showBugNotification: false,
        };
    },
    methods: {
        testSession: async function () {
            var url = "http://localhost:8081/testsession";
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                console.log("Response:", json)
            } else {
                console.log("Error:", response.status)
            }
        },
        getSessionInfo: async function () {
            var url = "http://localhost:8081/session";
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                console.log("Response:", json);
                this.whoami = json.session.Username;
            } else {
                console.log("Error:", response.status)
            }
        },
        buyPremiumSpell: function (event, whichspell) {
            console.log(whichspell);
            let button = event.target;
            button.classList.add("is-loading");
            setTimeout(() => {
                button.classList.remove("is-loading");
                button.classList.remove("is-warning");
                button.classList.add("is-primary");
                button.innerHTML = "";

                let purchasedspan = document.createElement("span");
                purchasedspan.textContent = "Purchased";

                let iconspan = document.createElement("span");
                iconspan.classList.add("icon");
                iconspan.classList.add("is-small");
                let iele = document.createElement("i");
                iele.classList.add("fas");
                iele.classList.add("fa-check");
                iconspan.appendChild(iele);
                button.appendChild(iconspan);


                button.appendChild(purchasedspan);
                button.disabled = true;

                this.premiumSpells[whichspell] = true;
                let price = document.querySelector(`.${whichspell}`);
                let inner = price.innerHTML;
                price.innerHTML = "";

                let strikethrough = document.createElement("s");
                strikethrough.innerHTML = inner;
                price.appendChild(strikethrough);
            }, 2000);
            console.log("BOUGHT SPELL");
        },
        helloThere: function () {
            // console.log("hello");
        },
        toggleHamburger: function () {
            this.hamburgerEnabled = !this.hamburgerEnabled;
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
            this.usernameInput = "";
            this.passwordInput = "";
        },
        loginLayout: function () {
            this.showVerifyPassword = false;
            this.showLoginButton = true;
            this.showRegisterButton = false;
            this.showLoginText = false;
            this.showSignUpText = true;
            this.usernameInput = "";
            this.passwordInput = "";
            this.passwordConfirmationInput = "";
            delete this.errors.username;
        },
        register: async function () {
            let username = this.usernameInput;
            let password = this.passwordInput;
            var url = "http://localhost:8081/users";
            let jstring = JSON.stringify({ username: username, password: password });
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: jstring
            });
            if (response.ok) {
                let json = await response.json();
                this.usernameInput = "";
                this.passwordInput = "";
                this.passwordConfirmationInput = "";
                this.showLoginModal = false;
                this.loginLayout();
                console.log("Response:", json)
            } else {
                alert("HTTP-Error: ", + response.status);
            }
        },
        verify: function () {
            let username = this.usernameInput;
            let password = this.passwordInput;
            let passwordConfirmation = this.passwordConfirmationInput;
            let pwcheck = false;
            let usercheck = false;


            if (this.showVerifyPassword){

                if (passwordConfirmation !== "") {
                    if (password !== passwordConfirmation) {
                        this.disabledButton = true;
                        this.errors.passwordmatch = "Passwords do not match"
                    } else {
                        pwcheck = true;
                        delete this.errors.passwordmatch;
                    }
                }
                if (username.length <= 4) {
                    this.disabledButton = true;
                    this.errors.username = "Username must be at least 5 characters"
                } else {
                    usercheck = true;
                    delete this.errors.username;
                }
            }
            if (pwcheck && usercheck) {
                this.disabledButton = false;
            }
            // if (username.)
        },
        login: async function () {
            console.log("Logging in");
            var url = "http://localhost:8081/login";
            let jsonbody = JSON.stringify({ username: this.usernameInput, password: this.passwordInput });
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: jsonbody
            });
            if (response.ok) {
                let json = await response.json();
                console.log("Response:", json)
                this.showLoginModal = false;
                this.usernameInput = "";
                this.passwordInput = "";
                this.getSessionInfo();
            } else {
                // TODO change this to a more user friendly response
                alert("HTTP-Error: ", + response.status);
            }
        },
        logout: async function () {
            console.log("Logging out")
            var url = "http://localhost:8081/logout";
            let response = await fetch(url, {
                method: 'POST',
            });
            if (response.ok) {
                let json = await response.json();
                console.log("Response:", json);
                this.whoami = "";
            } else {
                // TODO change this to a more user friendly response
                alert("HTTP-Error: ", + response.status);
            }
        },
        activateReportModal: function () {
            this.showReportModal = true;
        },
        postBugReport: async function () {
            if (this.bugReportMessage === "") {
                return
            }
            let button = document.getElementById("bugSubmitButton")
            button.classList.add("is-loading");
            var url = "http://localhost:8081/bugreport";
            var report = { message: this.bugReportMessage, email: this.bugEmail}
            console.log(report)
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(report)
            });
            if (response.ok) {
                let json = await response.json()
                button.classList.remove("is-loading");
                console.log(json)
                this.bugReportMessage = "";
                this.bugEmail = "";
                this.showBugNotification = true;
            } else {
                alert("HTTP-Error: ", response.status)
                button.classList.remove("is-loading");
            }
        },
        getTextareaLines: function () {
            let text = document.querySelector(".bptextarea").value;
            let arr = text.split("\n");
            return arr
        },
        linesToObject: function () {
            // @TODO change to get actual user
            var user = this.whoami;

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
            console.log(jsonProgram)
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
            } else if (response.status === 401){
                alert("HTTP-Error: " + "Please log in first");
                button.classList.remove("is-loading");
            }
            button.classList.remove("is-loading");
        },
        getBattleProgram: async function () {
            let url = "http://localhost:8081/battleprogram/" + this.whoami;
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                }
            });
            if (response.ok) {
                let json = await response.json();
                console.log("BattleProgram Found", json);
                this.extractInstructionsFromBP(json)
                this.saveTempProgram();
            } else if (response.status === 404) {
                console.log("Need to log in first");
            } else {
                alert("Something Happened");
            }
        },
        saveTempProgram: function () {
            let program = this.getTextareaLines();
            this.tempProgram = program;
            console.log(this.tempProgram);
        },
        extractInstructionsFromBP: function (bp) {
            let text = document.querySelector(".bptextarea");
            text.value = "";
            let instructions = [];
            for (let i = 0; i < bp.instructions.length; i++) {
                let instruction = bp.instructions[i].instruction;
                let args = bp.instructions[i].args;
                instructions.push(`${instruction} ${args.join(" ")}`);
            }
            text.value = instructions.join("\n");
            console.log(instructions);
        },
        showTutorial: function () {
            this.showBattlePage = false;
            this.showReplaysPage = false;
            this.showTutorialPage = true;

        },
        showBattle: async function () {
            this.showBattlePage = true;
            this.showReplaysPage = false;
            this.showTutorialPage = false;
            this.$nextTick(() =>{
                let text = document.querySelector(".bptextarea");
                console.log("Temp Program", this.tempProgram);
                text.value = this.tempProgram.join("\n");
            });
        },
        showReplays: function () {
            this.showBattlePage = false;
            this.showReplaysPage = true;
            this.showTutorialPage = false;
            this.drawBlankCanvas();
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

            const step = () => {
                if (index >= this.currentReplay.Frames.length) {
                    this.frameToDisplay = this.currentReplay.Frames[index - 1];
                    return; 
                }

                this.frameToDisplay = this.currentReplay.Frames[index];
                this.getPlayerLocationsInCurrentFrame();

                index++;
                waitTime = Math.max(50, waitTime -2*index)

                setTimeout(step, waitTime);
            }

            step();

            // let interval = setInterval(() => {
            //     this.frameToDisplay = this.currentReplay.Frames[index];
            //     this.getPlayerLocationsInCurrentFrame();
            //     this.drawPlayers();
            //     // console.log(this.frameToDisplay.ArenaFrame)
            //     index++;
                
            //     // clearInterval(interval);
            //     if (index >= this.currentReplay.Frames.length) {
            //         // -1 is the last frame to be displayed
            //         this.frameToDisplay = this.currentReplay.Frames[index - 1];
            //         clearInterval(interval);
            //         return;
            //     }
            //     waitTime -= 10;
            //     console.log("Waittime:", waitTime);
            // }, waitTime);
        },
        getPlayerLocationsInCurrentFrame: function () {
            let ploc = {
                p1row: -1,
                p1col: -1,
                p2row: -1,
                p2col: -1
            }
            console.log("IN getplayerlocations, this.frameToDisplay:", this.frameToDisplay);
            for (let i = 0; i < this.frameToDisplay.ArenaFrame.length; i++ ) {
                let p1rowindex = this.frameToDisplay.ArenaFrame[i].indexOf(1);
                let p2rowindex = this.frameToDisplay.ArenaFrame[i].indexOf(2);
                if (p1rowindex !== -1) {
                    ploc.p1row = i;
                    ploc.p1col = p1rowindex;
                }
                if (p2rowindex !== -1) {
                    ploc.p2row = i;
                    ploc.p2col = p2rowindex;
                }
            }
            this.ploc = ploc;
            return ploc
        },
        getDuel: async function () {
            let url = "http://localhost:8081/duels/" + this.duelUserInput; 
            console.log("hi");
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                this.currentReplay = json;
                console.log(json)
            } else {
                alert("HTTP-Error: ", + response.status);
            }

        },
        drawPlayers: function () {
            const canvas = document.getElementById("myCanvas");
            const ctx = canvas.getContext("2d");
            let x = this.ploc.p1row;
            let y = this.ploc.p1col;
            xt = x*40-20 + (2*x - 2);
            yt = y*40-20 + (2*x - 2);
            if (xt < 0) {
                xt = 0;
            }
            if (yt < 0) {
                yt = 0;
            }
            console.log("X:", xt, "Y:", yt);
            ctx.beginPath();
            ctx.arc(yt, xt, 18, 0, 2 * Math.PI);
            // ctx.fill();
            ctx.stroke();
        },
        drawBlankCanvas: function () {
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
                    var iterableNum = offset;
                    for (let i= 0; i < 16; i++) {
                        ctx.fillStyle = color1;
                        ctx.fillRect(iterableNum, row, size, size);
                        iterableNum = offset + iterableNum + size;
                    }
                    row = offset + row + size;
                }
                // ctx.beginPath();
                // ctx.arc(20, 20, 18, 0, 2 * Math.PI);
                // // ctx.fill();
                // ctx.stroke();

            }, 1);
        }
    },
    created: async function () {
    },
    mounted: function () {
        this.helloThere();
        this.getSessionInfo();
    },

}).mount("#app")
