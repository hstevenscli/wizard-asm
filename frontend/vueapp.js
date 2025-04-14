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
            showTutorialPage: false,
            showLandingPage: true,
            currentReplay: {},
            frameToDisplay: {
                ArenaFrame: [
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                ]
            },
            usernameInput: "",
            passwordInput: "",
            passwordConfirmationInput: "",
            disabledButton: true,
            errors: {},
            hamburgerEnabled: false,
            whoami: "",
            score: null,
            duelUserInput: "",
            ploc: {},
            tempProgram: [""],
            premiumSpells: {},
            showReportModal: false,
            bugReportMessage: "",
            bugEmail: "",
            showBugNotification: false,
            currentReplayInfoDisplay: {},
            notifications: {
                registration: false,
                saveProgram: false,
                playNow: false,
                getDuelInvalidUsername: false,
                getDuelInvalidUsernameWeird: false,
                getDuelBp1Empty: false,
                getDuelBp1Empty: false,
                invalidUserPassword: false,
                loginBadRequest: false,
                loginDatabaseError: false,
                loginSessionError: false,
            },
            opp: "",
            setTimeoutId: null,
            replayRunning: false,
            bugreports: [],
        };
    },
    methods: {
        testSession: async function () {
            var url = "/testsession";
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
            var url = "/session";
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
                this.getScore();
            } else {
                console.log("Error:", response.status)
            }
        },
        playNotification: function () {
            this.notifications.playNow = true;
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
        getPlayerColor(name) {
            if (name === this.whoami) {
                return 'hsl(171, 100%, 41%)';
            }
            if (name === this.opp) {
                return 'hsl(348, 100%, 61%)'
            }
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
            let button = document.getElementById("register-button");
            button.classList.add("is-loading");
            let username = this.usernameInput;
            let password = this.passwordInput;
            var url = "/users";
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
                // this.showLoginModal = false;
                this.loginLayout();
                this.notifications['registration'] = true;
                console.log("Response:", json)
                button.classList.remove("is-loading");
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
            let button = document.getElementById("login-button");
            button.classList.add("is-loading");
            var url = "/login";
            let jsonbody = JSON.stringify({ username: this.usernameInput, password: this.passwordInput });
            let response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: jsonbody
            });
            let json = await response.json();
            if (response.ok) {
                console.log("Response:", json)
                this.showLoginModal = false;
                this.usernameInput = "";
                this.passwordInput = "";
                this.getSessionInfo();
                this.showTutorial();
            } else if (response.status == 401 && json.status === "invalid username or password") {
                console.log(this.notifications["invalidUserPassword"])
                console.log("INVALID USERNAME OR PASSWORD");

                this.notifications["invalidUserPassword"] = true;

                console.log(this.notifications["invalidUserPassword"])
                setTimeout(() => {
                    this.notifications["invalidUserPassword"] = false;
                }, 10000);

            } else if (response.status == 400 && json.status === "Bad Request") {
                console.log("Login bad Request");

                this.notifications["loginBadRequest"] = true;
                setTimeout(() => {
                    this.notifications["loginBadRequest"] = false;
                }, 10000);


            } else if (response.status == 500) {
                if (json.status === "server error") {
                    console.log("DATABASE ERROR");

                    this.notifications["loginDatabaseError"] = true;
                    setTimeout(() => {
                        this.notifications["loginDatabaseError"] = false;
                    }, 10000);

                } else if (json.status === "Internal server error, please try again later") {
                    console.log("ERROR GENERATING SESSION");

                    this.notifications["loginSessionError"] = true;
                    setTimeout(() => {
                        this.notifications["loginSessionError"] = false;
                    }, 10000);

                }
            } else {
                alert("Unkown error encountered please try again later" + response.status);
            }
            button.classList.remove("is-loading");
        },
        logout: async function () {
            console.log("Logging out")
            var url = "/logout";
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
        getScore: async function () {
            var url = "/users/" + this.whoami;
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let user = await response.json();
                console.log("USER OBJ", user);
                this.score = user.Score;
            }
        },
        getCellClass(value) {
            if (value === 1) {
                return 'has-text-primary';
            }
            if (value === 2) {
                return 'has-text-danger';
            }
            if (value === 3) {
                return 'has-text-primary';
            }
            if (value === 4) {
                return 'has-text-danger';
            }
            if (value === 0) {
                return 'has-text-white';
            }
            if (value === 7) {
                return 'has-text-link';
            }
        },
        getBugReports: async function () {
            var url = "/bugreports";
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json();
                this.bugreports = json;
                console.log("Bugs:", this.bugreports);
            } else {
                alert("Error getting reports");
            }
        },
        makeMid: function () {
            const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
            let result = '';
            for (let i = 0; i < 16; i++) {
                result += chars.charAt(Math.floor(Math.random() * chars.length));
            }
            return result;
        },
        postBugReport: async function () {
            if (this.bugReportMessage === "") {
                return
            }
            let button = document.getElementById("bugSubmitButton")
            button.classList.add("is-loading");
            var url = "/bugreports";
            var mid = this.makeMid();
            console.log("MID", mid)
            var report = { message: this.bugReportMessage, email: this.bugEmail, mid: mid}
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
        deleteBugReport: async function (id) {
            var url = "/bugreports/" + id;
            let response = await fetch(url, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            if (response.ok) {
                let json = await response.json()
                console.log(json);
                this.getBugReports();
            } else {
                alert("HTTP-Error: ", response.status)
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
            var url = "/battleprogram"
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
                console.log("Response:", json);
                this.notifications['saveProgram'] = true;
                setTimeout(() => {
                    this.notifications['saveProgram'] = false;
                }, 2000);
            } else if (response.status === 401){
                alert("HTTP-Error: " + "Please log in first");
                button.classList.remove("is-loading");
            }
            button.classList.remove("is-loading");
        },
        getBattleProgram: async function () {
            let url = "/battleprogram/" + this.whoami;
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
        checkProgram: function () {
            let validInstructions = [
                "MAGMA", 
                "LIGHTNING", 
                "ACID", 
                "MOVE", 
                "SHIELD", 
                "TELEPORT", 
                "WAIT", 
                "RECHARGE", 
                "DIVINATION", 
                "SLOOP",
                "ELOOP",
                "JUMP",
                "CJUMP",
            ]
            let program = this.getTextareaLines();
            for (let i = 0; i < program.length; i++ ) {
                let args = [];
                let line = program[i].split(" ");
                let instruction = line[0].trim();
                // Check instruction
                if (!validInstructions.includes(instruction)) {
                    console.log("Instruction not recognized:", instruction);
                    console.log("On line ", i+1);
                }
                console.log("Instruction check passed!");
                // make sure that the length of line is equal to what it should be
                // this is to ensure that theres not extra arguments on something
                // that shouldnt have arguments
                //
                //
                // if instruction passes now check args
                // let count = this.returnCountOfArgs(instruction);
                // if (count == 2) {
                //     let arg1 = line[1]
                //     let arg2 = line[2]
                //     args.push(arg1, arg2)
                // } else if (count == 1) {
                //     let arg1 = line[1]
                //     args.push(arg1)
                // }
                // console.log("ARGS:", args);

            }
        },
        returnCountOfArgs: function (instruction) {
            let count = -1
            let twocount = [
                "magma",
                "acid",
                "teleport",
            ]
            let onecount = [
                "move",
                "shield",
                "recharge",
                "lightning",
                "divination",
                "sloop",
                "jump",
                "cjump",
            ]
            if (twocount.includes(instruction)) {
                count = 2;
            }
            if (onecount.includes(instruction)) {
                count = 1;
            }
            return count;
        },
        extractInstructionsFromBP: function (bp) {
            console.log("BP:", bp);
            let text = document.querySelector(".bptextarea");
            text.value = "";
            let instructions = [];
            for (let i = 0; i < bp.instructions.length; i++) {
                let instruction = bp.instructions[i].instruction;
                let args = bp.instructions[i].args;
                console.log("Args:", args);
                if (args) {
                    console.log("GOT HERE")
                    instructions.push(`${instruction} ${args.join(" ")}`);
                } else {
                    console.log("IN ELSE")
                    instructions.push(`${instruction}`);
                }
            }
            text.value = instructions.join("\n");
            console.log(instructions);
        },
        showLanding: function () {
            this.showLandingPage = true;
            this.showBattlePage = false;
            this.showReplaysPage = false;
            this.showTutorialPage = false;
        },
        showTutorial: function () {
            this.showBattlePage = false;
            this.showReplaysPage = false;
            this.showTutorialPage = true;
            this.showLandingPage = false;

        },
        showBattle: async function () {
            this.showBattlePage = true;
            this.showReplaysPage = false;
            this.showTutorialPage = false;
            this.showLandingPage = false;
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
            this.showLandingPage = false;
            // this.drawBlankCanvas();
        },
        runGame: async function () {
            let button = document.getElementById("gamebutton");
            button.classList.add("is-loading");
            var url = "/game"
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
            let url = "/battlereplay" 
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
        stopReplay: function () {
            clearTimeout(this.setTimeoutId);
            this.replayRunning = false;
        },
        startReplay: function () {
            this.replayRunning = true;
            console.log(this.replayRunning);
            var index = 0;
            var waitTime = 1000;


            const step = () => {
                if (index >= this.currentReplay.Frames.length) {
                    this.replayRunning = false;
                    this.frameToDisplay = this.currentReplay.Frames[index - 1];
                    return; 
                }

                this.frameToDisplay = this.currentReplay.Frames[index];
                this.getPlayerLocationsInCurrentFrame();

                index++;
                waitTime = Math.max(30, waitTime -2*index)

                this.setTimeoutId = setTimeout(step, waitTime);
            }

            step();
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
            let url = "/duels/" + this.duelUserInput; 
            console.log("hi");
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            let json = await response.json();
            if (response.ok) {
                this.currentReplay = json;
                this.currentReplayInfoDisplay = json.GameoverInfo;
                this.currentReplayInfoDisplay.RealCount = json.Frames.length;
                this.getScore();
                this.opp = json.Opp;
            } else {
                this.handleErrorResponses(response, json);
            }
        },
        getDuelRandom: async function () {
            let url = "/duels/random"; 
            console.log("hi");
            let response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            });
            let json = await response.json();
            if (response.ok) {
                this.currentReplay = json;
                this.currentReplayInfoDisplay = json.GameoverInfo;
                this.currentReplayInfoDisplay.RealCount = json.Frames.length;
                this.getScore();
                this.opp = json.Opp;
            } else {
                this.handleErrorResponses(response, json);
            }
        },
        handleErrorResponses: function (response, json) {
            if (response.status == 409 && json.status === "User2 not found") {
                console.log("USER 2 NOT FOUND");

                this.notifications["getDuelInvalidUsername"] = true;
                setTimeout(() => {
                    this.notifications["getDuelInvalidUsername"] = false;
                }, 3000);

            } else if (response.status == 409 && json.status === "Battle program 2 is empty") {
                console.log("BATTLE PROGRAM 2 IS EMPTY");

                this.notifications["getDuelBp2Empty"] = true;
                setTimeout(() => {
                    this.notifications["getDuelBp2Empty"] = false;
                }, 10000);


            } else if (response.status == 409 && json.status === "User1 not found") {
                console.log("USER 1 NOT FOUND");

                this.notifications["getDuelInvalidUsernameWeird"] = true;
                setTimeout(() => {
                    this.notifications["getDuelInvalidUsernameWeird"] = false;
                }, 10000);

            } else if (response.status == 409 && json.status === "Battle program 1 is empty") {
                console.log("BATTLE PROGRAM 1 IS EMPTY");

                this.notifications["getDuelBp1Empty"] = true;
                setTimeout(() => {
                    this.notifications["getDuelBp1Empty"] = false;
                }, 10000);
            } else {
                alert("Unknown/Internal Server Error. Please try again later");
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
