<template>
import LoginVue from './Login.vue';
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
                        @focus="focusInputFlex"
                        @blur="focusInputFlex"
                        v-model="username"
                    />
                </div>
                <div class="input-warning"></div>
            </div>
            <div class="input-zone">
                <label for="password">密码:</label><br>
                <div class="input-flex">
                    <span class="input-logo logo-password"></span>
                    <input
                        type="password" 
                        placeholder="password" 
                        class="input-ipt" 
                        @focus="focusInputFlex"
                        @blur="focusInputFlex"
                        v-model="password"
                    />
                </div>
                <div class="input-warning"></div>
            </div>

            <div class=check-zone>
                <input type="checkbox" id="remember"><label for="remember">Remember me</label>
                <a href="/forget">Forgot password ?</a>
            </div>
            <div class="submit-zone">
                <button type="button" @click="submitForm">Log in</button>
            </div>
            <div class="register-zone">
                or <a href="/register">register</a> now!
            </div>
        </form>
    </div>
</template>

<script>
export default {
    props: [
        "username", "password"
    ],
    data: () => {
        return {
            username: "",
            password: "",
        }
    },
    methods: {
        focusInputFlex: event => {
            event.target.parentNode.classList.toggle("inputFlexFocus");
        },
        submitForm: () => {
            console.log("entry");

            // if(this.$data.username === "") {
            //     this.$methods.shake(1, "please input username");
            //     return
            // }
            // if(this.$data.password === "") {
            //     this.$methods.shake(2, "please input password");
            //     return 
            // }
        },
        shake: (pos, msg) => {
            console.log(pos, msg);
        }
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
    margin-bottom: 10px;
    display: inline-flex;
    border-radius: 5px;
    border: 1px solid #afafaf;
    background-color: rgba(255,255,255,.7);
}
.input-flex:hover{
    border: 1px solid rgba(24,144,255,1);
}
.inputFlexFocus{
    border: 1px solid rgba(24,144,255,1);
    box-shadow: 0 0 5px rgba(24,144,255,1);
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
}
.input-ipt:focus + .input-flex{
    border: 1px solid rgba(24,144,255,1);
    box-shadow: 0 0 5px rgba(24,144,255,1);
}
.input-warning{
    height: 15px;
    width: 100%;
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
    border: 1px solid rgba(24,144,255,1);
    border-radius: 5px; 
    cursor: pointer;
    background-color: rgba(24,144,255,1);
    color: white;
    outline: none;
}
.submit-zone > button:hover {
    box-shadow: 0 0 5px rgba(24,144,255,1);
}
.register-zone{
    width: 100%;
    height: 20px;
    margin-top: 10px;
}
.register-zone > a {
    color: rgba(24,144,255,1);
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
