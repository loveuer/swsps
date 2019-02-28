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
        
        <li class="login">
            <div class="login-menu" v-show="logined">
                <li>我的信息</li>
                <li>修改密码</li>
                <li>退出登录</li>
            </div>
            <div class="login-true" v-show="logined" @click="showUserMenu">
                <profile-icon ref="ProfileIcon"/>
            </div>
            <div class="login-false" v-show="!logined" @click="meetUrl('/login');"></div>
        </li>
    </div>
</template>

<script>
import router from "../router"
import ProfileIcon from "./ProfileIcon.vue"

export default {
    data() {
        return {
            logined: false,
        }
    },
    methods: {
        meetUrl: function(url) {
            router.push(url);
        },
        showUserMenu: function() {

        },
    },
    mounted() {
        // check session
        this.$http
            .get("/api/session/" + this.$getcookie("usession"))
            .then(resp => {
                if(resp.data["StatusCode"] === 200 && resp.data["Msg"] !== "wrong session"){
                    this.logined = !this.logined;
                    if(this.$user.id !== resp.data["Val"]){
                        this.$user.id = resp.data["Val"];
                        this.$refs.ProfileIcon.getProfile();
                    }
                }
            })
            .catch(error => {console.log(error)})
    },
    components: {
        ProfileIcon,
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
    width:250px;
    height: 50px;
    margin-right: 30px;
    display: flex;
    align-items: center;
}
.login-true{
    height: 50px;
    width: 50px;
    margin-left: auto;
}
.login-false{
    margin-left: auto;
    width: 46px;
    height: 46px;
    border: 2px solid #0082E4;
    border-radius: 50%;
    background-image: url("~@/assets/img/登录.svg");
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    background-size: 85%;
}
.login-false:hover{
    box-shadow: 0 0 3 #0082E4;
}
.login-menu{
    width: 200px;
    height: 50px;
}
.login-menu > li {
    display: flex;
    height: 46px;
    float: left;
    border: 1px solid #0082E4;
    align-items: center;
}
.login-menu > li:hover{
    background-color: #0082E4;
    color: #fff;
}
</style>
