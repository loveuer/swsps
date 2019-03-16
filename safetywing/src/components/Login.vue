<template>
    <div id="login-container">
        <div class="login-form">
            <div>
                <el-row>
                    <label for="username">账号:</label>
                </el-row>
                <el-row>
                    <div class="login-input" v-bind:class="{ ushake:isushake }">
                        <div class="login-input-logo">
                            <svg style="width:19px;height:19px;margin-left:8px;margin-top:8px;" viewBox="0 0 1024 1024" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink"><path d="M512.03 580.92c-142.52 0-258.46-115.94-258.46-258.46S369.51 64 512.03 64s258.46 115.94 258.46 258.46-115.95 258.46-258.46 258.46z m0-482.46c-123.52 0-224 100.49-224 224s100.48 224 224 224 224-100.49 224-224-100.49-224-224-224z" p-id="877" fill="#bfbfbf"></path><path d="M150.18 960c-9.52 0-17.23-7.71-17.23-17.23 0-218.53 170.05-396.31 379.08-396.31 93.54 0 183.38 35.91 252.96 101.13 6.95 6.51 7.3 17.42 0.79 24.35-6.52 6.95-17.42 7.29-24.35 0.79-63.18-59.2-144.64-91.81-229.4-91.81-190.02 0-344.62 162.33-344.62 361.85 0 9.52-7.72 17.23-17.23 17.23z" p-id="878" fill="#bfbfbf"></path><path d="M873.82 960c-9.52 0-17.23-7.71-17.23-17.23 0-199.52-154.6-361.85-344.62-361.85-84.76 0-166.22 32.61-229.4 91.81-6.93 6.51-17.83 6.18-24.35-0.79-6.51-6.93-6.16-17.84 0.79-24.35 69.58-65.22 159.42-101.13 252.96-101.13 209.02 0 379.08 177.78 379.08 396.31 0 9.52-7.71 17.23-17.23 17.23z" p-id="879" fill="#bfbfbf"></path></svg>
                        </div>
                        <input type="text" id="username" class="login-input-username" v-model="loginform.username">
                        <div class="login-input-tool"></div>
                    </div>
                </el-row>
            </div>
            <div>
                <el-row>
                    <label for="password">密码:</label>
                </el-row>
                <el-row>
                    <div class="login-input" v-bind:class="{ pshake:ispshake }">
                        <div class="login-input-logo">
                            <svg style="width:19px;height:19px;margin-left:8px;margin-top:8px;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" ><path d="M718.77 960H305.23c-66.51 0-120.62-54.12-120.62-120.62V502.91c0-66.5 54.11-120.62 120.62-120.62h383.52V275.21c0-97.46-79.29-176.75-176.75-176.75-97.45 0-176.74 79.29-176.74 176.75v38.89c0 9.52-7.72 17.23-17.23 17.23s-17.23-7.71-17.23-17.23v-38.89C300.8 158.75 395.54 64 512 64s211.21 94.75 211.21 211.21v107.17c64.46 2.34 116.17 55.51 116.17 120.53v336.47c0 66.5-54.1 120.62-120.61 120.62zM305.23 416.76c-47.5 0-86.15 38.65-86.15 86.15v336.47c0 47.5 38.65 86.15 86.15 86.15h413.54c47.5 0 86.15-38.65 86.15-86.15V502.91c0-47.5-38.65-86.15-86.15-86.15H305.23zM512 853.15c-9.52 0-17.23-7.71-17.23-17.23V722.66c-56.59-8.36-100.15-57.25-100.15-116.11 0-64.73 52.66-117.4 117.38-117.4 64.73 0 117.39 52.67 117.39 117.4 0 58.86-43.56 107.74-100.16 116.11v113.26c0 9.52-7.71 17.23-17.23 17.23z m0-329.54c-45.73 0-82.92 37.2-82.92 82.94 0 45.72 37.2 82.91 82.92 82.91s82.93-37.19 82.93-82.91c0-45.74-37.2-82.94-82.93-82.94z" p-id="1032" fill="#bfbfbf"></path></svg>
                        </div>
                        <input :type="passwordType" id="password" class="login-input-password" v-model="loginform.password">
                        <div class="login-input-tool"></div>
                    </div>
                </el-row>
            </div>
            <div style="margin-top:20px;">
                <el-button type="primary" @click="doLogin">登录</el-button>
                <el-button @click="backHome">取消</el-button>
            </div>
        </div>
    </div>    
</template>

<script>
import router from '../router'
import qs from 'qs';
import { endianness } from 'os';

export default {
    data() {
        return {
            loginform: {
                username: '',
                password: '',
            },
            isushake: false,
            ispshake: false,
            passwordType: 'password',
        };
    },
    methods: {
        backHome: () => {
            router.push('/');
        },
        doLogin: function() {
            this.shakeInput(3, '账号或密码错误!');
            // this.$http.post('/api/login', qs.stringify(this.loginform))
            //     .then(res => {
            //         if (res.status === 200) {
            //             console.log(res.data);
            //         } else {
            //             this.shakeInput('账号或密码不正确');
            //         };
            //     });
        },
        shakeInput: function(which, msg) {
            let username = this.loginform.username;
            if (which === 1) {
                this.isushake = true;
                this.loginform.username = msg;
            } else if (which === 2) {
                this.ispshake = true;
                this.passwordType = 'text';
                this.loginform.password = msg;
            } else if (which === 3) {
                this.isushake = true;
                this.ispshake = true;
                this.passwordType = 'text';
                this.loginform.username = msg;
                this.loginform.password = msg;
            };
            setTimeout(() => {
                this.endShakeInput(username);
            }, 1000);

        },
        endShakeInput(u) {
            this.ispshake = false;
            this.isushake = false;
            this.loginform.username = u;
            this.loginform.password = "";
            this.passwordType = 'password';
        },
    },
};
</script>

<style scoped>
#login-container{
    min-height: 100%;
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background-image: url("~@/assets/img/runway2.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
    background-color: rgba(0,0,0,0.4);
    /* background-color: #409EFF; */
}
.login-form{
    min-height: 220px;
    min-width: 400px;
    width: 400px;
    margin-top: -30px;
    padding: 10px;
    border-radius: 15px;
    border: 2px solid #ddd;
    background-color: rgba(255,255,255,0.4);
}
.login-form > div {
    margin-top: 10px;
    width: 100%;
}
.login-input{
    border: 1px solid #ddd;
    border-radius: 5px;
    height: 35px;
    line-height: 35px;
    margin-top: 5px;
    width: 398px;
    display: inline-flex;
    background-color: #fff;
}
.login-input:hover {
    border: 1px solid #409EFF;
}
.login-input:focus-within{
    border: 1px solid #409EFF;
    box-shadow: 0 0 5px #409EFF;
}
.ushake,.pshake{
    animation-name: shake;
    animation-duration: 0.8s;
    animation-fill-mode: none;
    border: 1px solid #F56C6C;
}
.ushake > input , .pshake > input {
    color: #F56C6C;
}
.login-input > input {
    border: 0;
    outline: none;
    height: 33px;
    line-height: 33px;
    font-size: 16px;
    background-color: rgba(255,255,255,0);
    width: 100%;
    text-indent: 5px;
}
.login-input > .login-input-logo {
    height: 35px;
    width: 35px;
    border-radius: 5px;    
}
.login-input > .login-input-tool {
    height: 35px;
    width: 35px;
    border-radius: 5px;    
}
@keyframes shake {
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
input:-webkit-autofill,
input:-webkit-autofill:hover, 
input:-webkit-autofill:focus, 
input:-webkit-autofill:active  {
    -webkit-box-shadow: 0 0 0 30px white inset !important;
}
</style>

