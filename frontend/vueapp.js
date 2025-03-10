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
                console.log("IAM:", this.whoami)
                this.whoami = json.session.Username;
                console.log("IAM:", this.whoami)
            } else {
                console.log("Error:", response.status)
            }
        },
        helloThere: function () {
            console.log("hello");
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
        getTextareaLines: function () {
            let text = document.querySelector(".textarea").value;
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
                    var iterableNum = offset;
                    for (let i= 0; i < 16; i++) {
                        ctx.fillStyle = color1;
                        ctx.fillRect(iterableNum, row, size, size);
                        iterableNum = offset + iterableNum + size;
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
        this.getSessionInfo();
    },
    mounted: function () {
    }

}).mount("#app")
