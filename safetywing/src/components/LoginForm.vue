<template>
    <div id="form">
        <form @keypress.enter="submitForm">
            <div class="input-zone">
                <label for="username">用户名:</label><br>
                <div class="input-flex">
                    <span class="input-logo logo-username">
                    </span>
                    <input
                        type="text" 
                        placeholder="username" 
                        class="input-ipt" 
                        @focus="focusInputFlex($event)"
                        @blur="unFocusInputFlex($event)"
                        @mouseenter="hoverInputFlex($event)"
                        @mouseleave="unHoverInputFlex($event)"
                        v-model="username"
                        id="username"
                    />
                </div>
                <div class="input-warning">
                    <font></font>
                </div>
            </div>
            <div class="input-zone">
                <label for="password">密码:</label><br>
                <div class="input-flex">
                    <span class="input-logo logo-password"></span>
                    <input
                        type="password" 
                        placeholder="password" 
                        class="input-ipt" 
                        @focus="focusInputFlex($event)"
                        @blur="unFocusInputFlex($event)"
                        @mouseenter="hoverInputFlex($event)"
                        @mouseleave="unHoverInputFlex($event)"
                        v-model="password"
                        id="password"
                    />
                </div>
                <div class="input-warning">
                    <font></font>
                </div>
            </div>

            <div class=check-zone>
                <input type="checkbox" id="remember"><label for="remember">Remember me</label>
                <a href="/forget">Forgot password ?</a>
            </div>
            <div class="submit-zone">
                <button type="button" @click="submitForm" :style="loginBtnStyle" @mouseenter="submitBorder($event)" @mouseleave="submitUnBorder($event)">Log in</button>
            </div>
            <div class="register-zone">
                or <a href="/register">register</a> now!
            </div>
        </form>
    </div>
</template>

<script>
import router from "../router"

export default {
    data() {
        return {
            username: "",
            password: "",
            loginBtnStyle: {
                backgroundColor: this.$color,
                border: "1px solid " + this.$color,
            },
        }
    },  
    methods: {
        focusInputFlex: function(event) {
            event.target.parentNode.style.border = "1px solid " + this.$color;
            event.target.parentNode.style.boxShadow = "0 0 5px " + this.$color;
            event.target.parentNode.classList.remove("shakeDom");
        },
        unFocusInputFlex: function(event) {
            event.target.parentNode.style.border = "1px solid rgba(222,222,222,1)";
            event.target.parentNode.style.boxShadow = "none";
        },
        hoverInputFlex: function(event) {
            event.target.parentNode.style.border = "1px solid " + this.$color;
        },
        unHoverInputFlex: function(event) {
            if(!(event.target === document.activeElement)){
                event.target.parentNode.style.border = "1px solid rgba(222,222,222,1)";
            };
        },
        submitForm: function() {
            if(this.username === "") {
                this.shake(1, "please input username");
                return
            }
            if(this.password === "") {
                this.shake(2, "please input password");
                return 
            }
            this.$http.post("/api/login", {
                username: this.username,
                password: this.password
            }).then(resp => {
                if(resp.data["Msg"] === "200") {
                    //login sucess
                    this.$setcookie("usession", resp.data["Val"], 60);
                    this.$user.username = this.username;
                    router.push("/");
                } else {
                    this.shake(1, "username or password not right");
                    this.shake(2, "username or password not right");
                    this.password = "";
                    console.log(resp.data["Msg"]);
                };
            }).catch(function(error){
                console.log(error)
            }).fanally;
            return
        },
        shake: (pos, msg) => {
            let whichDict = {1: "#username", 2: "#password"};
            let shakeDom = document.querySelector(whichDict[pos]).parentNode;
            shakeDom.classList.toggle("shakeDom");
            document.querySelectorAll("font")[pos-1].innerHTML = msg;
            setTimeout(function() {
                shakeDom.classList.remove("shakeDom");
                document.querySelectorAll("font")[pos-1].innerHTML = "";
            }, 3000);     
        },
        submitBorder:function(e) {
            e.target.style.boxShadow = "0 0 5px " + this.$color;
        },
        submitUnBorder:function(e){
            e.target.style.boxShadow = "none";
        }
    },
    mounted() {
        document.querySelector("input[type=text]").focus();
    }
}
</script>

<style scoped>
#form{
    width: 100%;
}
.input-zone{
    width: 100%;
}
label{
    font-size: 14px;
    font-weight: bold;
}
.input-flex{
    width: 100%;
    margin-top: 5px;
    display: inline-flex;
    border-radius: 5px;
    background-color: rgba(255,255,255,.7);
}
.input-logo {
    height: 25px;
    width: 25px;
    padding: 3px;
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    background-size: 65%;
}
.logo-username{
    background-image: url("~@/assets/img/账号.svg");
}
.logo-password{
    background-image: url("~@/assets/img/密码.svg");
}
.input-ipt{
    width: 100%;
    height: 25px;
    border: 0;
    outline: none;
    padding: 3px;
    text-indent: 3px;
    background-color: rgba(0,0,0,0);
    border-radius: 5px;
}
.check-zone{
    height: 25px;
    width: 100%;
    display: flex;
    align-items: center;
}
.check-zone > input[type=checkbox]{
    height: 15px;
    width: 15px;
}
.check-zone > a {
    margin-left: auto;
    color: rgba(24,144,255,1);
}
.submit-zone{
    height: 30px;
    width: 100%;
    margin-top: 10px;
}
.submit-zone > button {
    width: 100%;
    height: 100%;
    border-radius: 5px; 
    cursor: pointer;
    color: white;
    outline: none;
    font-size: 14px;
}
.register-zone{
    width: 100%;
    height: 20px;
    margin-top: 10px;
}
.register-zone > a {
    color: rgba(24,144,255,1);
}

.input-warning{
    height: 15px;
    width: 100%;
    margin-bottom: 5px;
}
.input-warning>font {
    font-size: 14px;
    font-weight: bold;
    color: #ed4337;
}

.shakeDom{
    animation: shakeInput 0.7s forwards;
    border: 1px solid#ed4337 !important;
    box-shadow: 0 0 5px #ed4337 !important;
}
@keyframes shakeInput {
    10%, 90% {
        transform: translate3d(-1px, 0, 0);
    }
    20%, 80% {
        transform: translate3d(2px, 0, 0);
    }

    30%, 50%, 70% {
        transform: translate3d(-4px, 0, 0);
    }

    40%, 60% {
        transform: translate3d(4px, 0, 0);
    }
}
</style>
