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
        }


    },
    created: async function () {
        this.helloThere();
    }

}).mount("#app")
