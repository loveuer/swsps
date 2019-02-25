<template>
    <div class="nav">
        <li class="logo" @click="meetUrl('/');"></li>
        <li class="menu" @click="meetUrl('/affairs');">
            <div class="menu-img menu-affairs"></div>
            <div class="menu-describe">公司事务</div>
        </li>
        <li class="menu" @click="meetUrl('/sps');">
            <div class="menu-img menu-sps"></div>
            <div class="menu-describe">备件管理</div>
        </li>
        <li class="menu" @click="meetUrl('/admin');">
            <div class="menu-img menu-admin"></div>
            <div class="menu-describe">管理后台</div>
        </li>
        <li class="login" @click="meetUrl('/login');">
            <div class="login-true" v-show="logined"></div>
            <div class="login-false" v-show="!logined"></div>
        </li>
    </div>
</template>

<script>
import router from "../router"
import axios from "axios"

export default {
    data() {
        return {
            logined: false,
        }
    },
    methods: {
        meetUrl: (url) => {
            router.push(url);
        },
    },
    mounted() {
        // check session
        axios
            .get("/api/session/" + this.$cookies("usession"))
            .then(resp => {
                if(resp.data["StatusCode"] === 200 && resp.data["Msg"] !== ""){
                    this.logined = !this.logined;
                }
            })
            .catch(error => {console.log(error)})
    }
}
</script>

<style scoped>
.nav{
    display: flex;
    align-items: center;
    width: 100%;
    height: 50px;
    margin-top: 5px;
    padding: 3px;
}
.nav > li {
    list-style: none;
    display: block;
    height: 50px;
    display: flex;
    align-items: center;
    cursor: pointer;
    flex-wrap: wrap;
}
.nav > .logo {
    width: 240px;
    background-image: url("~@/assets/img/logo.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    background-size: 80%;
}
.nav > .menu {
    margin-left: 20px;
    width: 100px;
    border: 2px solid #fff;
    border-radius: 5px;
}
.nav > .menu:hover{
    border: 2px solid #0082E4;
}
.nav > .menu:hover > .menu-describe {
    color: #0082E4;
}
.nav > .menu > .menu-img {
    height: 30px;
    width: 100%;
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
}
.nav > .menu > .menu-affairs {
    background-image: url("~@/assets/img/公司事务-blue.svg");
}
.nav > .menu > .menu-sps {
    background-image: url("~@/assets/img/备件管理-blue.svg");
}.nav > .menu > .menu-admin {
    background-image: url("~@/assets/img/管理员.svg");
}
.menu-describe {
    width: 100%;
    text-align: center;
}
.login{
    margin-left: auto;
    width:50px;
    height: 50px;
    margin-right: 30px;
}
.login > div {
    height: 46px;
    width: 46px;
    border-radius: 50%;
    border: 2px solid #0082E4;
    cursor: pointer;
}
.login > div:hover{

}
.login-false{
    width: 100%;
    height: 100%;
    background-image: url("~@/assets/img/登录.svg");
    background-repeat: no-repeat;
    /* background-position: 80% top 70% left; */
    background-position: center;
    background-size: contain;
    background-size: 85%;
}
</style>
