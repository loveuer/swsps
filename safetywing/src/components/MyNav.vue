<template>
    <div class="nav">
        <li class="logo" @click="meetUrl('/');"></li>
        <!-- <li 
            class="menu" 
            @click="meetUrl('/affairs');"
            @mouseenter="borderMenu($event)"
            @mouseleave="unBorderMenu($event)"
        >
            <div class="menu-img menu-affairs">
                <svg 
                    t="1550891888373" 
                    viewBox="0 0 1024 1024" 
                    version="1.1" 
                    xmlns="http://www.w3.org/2000/svg" 
                    p-id="1109" 
                    xmlns:xlink="http://www.w3.org/1999/xlink" 
                    width="100%" 
                    height="30"
                >
                    <path v-bind:fill="this.$store.state.customColor" d="M928.3 920.3H98c-17.7 0-32-14.3-32-32V416.8c0-14.5 9.8-27.2 23.8-30.9 14-3.7 28.8 2.5 36 15.2C195 523 370.8 548.6 506.1 548.6c97.6 0 184.4-14.9 250.8-43 69.9-29.6 117.7-73.3 142-129.8 6-14 21-21.8 35.9-18.7 14.9 3.1 25.5 16.2 25.5 31.3v499.9c0 17.7-14.3 32-32 32z m-798.3-64h766.3v-366c-3 3-6.1 5.9-9.2 8.7-28.9 26.1-64.3 48.1-105.2 65.5-75.3 31.9-168.1 48.1-275.8 48.1-101.1 0-188-13.1-258.2-38.9-46.5-17.1-85.9-39.8-117.9-67.9v350.5z">
                    </path>
                    <path v-bind:fill="this.$store.state.customColor" d="M130 416.8H66V260.6c0-17.7 14.3-32 32-32h830.3c17.7 0 32 14.3 32 32v127.8h-64v-95.8H130v124.2zM508.5 674.1c-17.7 0-32-14.3-32-32V509c0-17.7 14.3-32 32-32s32 14.3 32 32v133c0 17.8-14.3 32.1-32 32.1z">
                    </path>
                    <path v-bind:fill="this.$store.state.customColor" d="M692.1 292.6H331c-24.3 0-44-19.7-44-44v-105c0-24.3 19.7-44 44-44h361.1c24.3 0 44 19.7 44 44v105c0 24.3-19.7 44-44 44z m-341.1-64h321.1v-65H351v65z">
                    </path>
                </svg>
            </div>
            <div class="menu-describe">公司事务</div>
        </li> -->
        <li 
            class="menu" 
            @click="meetUrl('/sps');"
            @mouseenter="borderMenu($event)"
            @mouseleave="unBorderMenu($event)"    
        >
            <div class="menu-img menu-sps">
                <svg t="1550892007743" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1508" xmlns:xlink="http://www.w3.org/1999/xlink" width="100%" height="30">
                    <path v-bind:fill="this.$store.state.customColor" d="M568.889 967.1110000000001v-512H967.1100000000001v512H568.89zM910.222 512H625.778v398.222h284.44399999999996V512zM568.89 56.88900000000001H967.1100000000001v284.44399999999996H568.89V56.889999999999986z m56.889 227.555h284.44399999999996V113.77800000000002H625.778v170.666zM56.88799999999998 682.667h398.223V967.1100000000001H56.889999999999986V682.667z m56.89 227.555h284.44399999999996V739.556H113.77800000000002v170.666zM56.88799999999998 56.889999999999986h398.223v512H56.889999999999986v-512zM113.779 512h284.44399999999996V113.77800000000002H113.77800000000002V512z" ></path>
                </svg>
            </div>
            <div class="menu-describe">备件管理</div>
        </li>
        <li 
            class="menu" 
            @click="meetUrl('/admin');"
            @mouseenter="borderMenu($event)"
            @mouseleave="unBorderMenu($event)"    
        >
            <div class="menu-img menu-admin">
                <svg t="1550894311750" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="864" xmlns:xlink="http://www.w3.org/1999/xlink" width="100%" height="30">
                    <path v-bind:fill="this.$store.state.customColor" d="M512 568.888889c-159.288889 0-284.444444-125.155556-284.444444-284.444445s125.155556-284.444444 284.444444-284.444444 284.444444 125.155556 284.444444 284.444444-125.155556 284.444444-284.444444 284.444445z m0-56.888889c125.155556 0 227.555556-102.4 227.555556-227.555556s-102.4-227.555556-227.555556-227.555555-227.555556 102.4-227.555556 227.555555 102.4 227.555556 227.555556 227.555556zM824.888889 1024c-79.644444 0-142.222222-62.577778-142.222222-142.222222s62.577778-142.222222 142.222222-142.222222 142.222222 62.577778 142.222222 142.222222-62.577778 142.222222-142.222222 142.222222z m0-56.888889c45.511111 0 85.333333-39.822222 85.333333-85.333333s-39.822222-85.333333-85.333333-85.333334-85.333333 39.822222-85.333333 85.333334 39.822222 85.333333 85.333333 85.333333z"></path>
                    <path v-bind:fill="this.$store.state.customColor" d="M796.444444 455.111111h56.888889v284.444445h-56.888889zM853.333333 455.111111h113.777778v56.888889h-113.777778zM853.333333 568.888889h113.777778v56.888889h-113.777778z"></path>
                    <path v-bind:fill="this.$store.state.customColor" d="M512 512l-51.2 62.577778C261.688889 597.333333 113.777778 762.311111 113.777778 967.111111H56.888889v-56.888889h5.688889c28.444444-221.866667 216.177778-398.222222 449.422222-398.222222z m0 0l-51.2 62.577778C261.688889 597.333333 113.777778 762.311111 113.777778 967.111111H56.888889v-56.888889h5.688889c28.444444-221.866667 216.177778-398.222222 449.422222-398.222222z"></path>
                </svg>
            </div>
            <div class="menu-describe">管理后台</div>
        </li>
        <li
            class="menu"
            id="li-colorSelector"
            v-if="this.ishome"
            @mouseenter="borderMenu($event)"
            @mouseleave="unBorderMenu($event)"
        >
            <div class="menu-img">
                <svg t="1551790244181" style="width:100%;height:30;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1871" xmlns:xlink="http://www.w3.org/1999/xlink">
                    <path v-bind:fill="this.$store.state.customColor" d="M214.634667 390.784C213.781333 381.525333 213.333333 372.16 213.333333 362.666667c0-164.949333 133.717333-298.666667 298.666667-298.666667s298.666667 133.717333 298.666667 298.666667c0 9.493333-0.426667 18.858667-1.301334 28.117333C910.976 438.464 981.333333 541.674667 981.333333 661.333333c0 164.949333-133.717333 298.666667-298.666666 298.666667-63.445333 0-122.282667-19.797333-170.666667-53.546667A297.301333 297.301333 0 0 1 341.333333 960C176.384 960 42.666667 826.282667 42.666667 661.333333c0-119.658667 70.357333-222.869333 171.968-270.549333z m20.949333 85.226667A213.248 213.248 0 0 0 128 661.333333c0 117.824 95.509333 213.333333 213.333333 213.333334 38.869333 0 75.306667-10.389333 106.666667-28.544A297.386667 297.386667 0 0 1 384 661.333333c0-9.493333 0.426667-18.858667 1.301333-28.117333a299.754667 299.754667 0 0 1-149.717333-157.226667z m233.770667 182.293333L469.333333 661.333333c0 117.824 95.509333 213.333333 213.333334 213.333334s213.333333-95.509333 213.333333-213.333334c0-79.36-43.306667-148.565333-107.584-185.322666C743.765333 584.746667 636.8 661.333333 512 661.333333a299.093333 299.093333 0 0 1-42.645333-3.029333zM512 576c117.824 0 213.333333-95.509333 213.333333-213.333333S629.824 149.333333 512 149.333333s-213.333333 95.509333-213.333333 213.333334 95.509333 213.333333 213.333333 213.333333z"></path>
                </svg>
            </div>
            <div class="menu-describe">主题颜色</div>
            <div class="color-selector" :style="{ border:'1px solid '+this.$store.state.customColor }">
                <div>选择主题:</div>
                <div>
                    <li @click="changeColor('#ED4337')" @mouseenter="hoverChgColor('#ED4337', $event)" @mouseleave="unHoverChgColor($event)">red</li>
                    <li @click="changeColor('#79cc92')" @mouseenter="hoverChgColor('#79cc92', $event)" @mouseleave="unHoverChgColor($event)">green</li>
                    <li @click="changeColor('#1890ff')" @mouseenter="hoverChgColor('#1890ff', $event)" @mouseleave="unHoverChgColor($event)">blue</li>
                    <li @click="changeColor('#b090c0')" @mouseenter="hoverChgColor('#b090c0', $event)" @mouseleave="unHoverChgColor($event)">violet</li>
                    <li>
                        <div @click="customColorSelector">自定义</div>
                        <div :style="{ display:'none' }"><input type="color" v-bind:value="this.$store.state.customColor" @change="customChgColor($event)"></div>
                    </li>
                </div>
            </div>
        </li>
        
        <li class="login">
            <div class="login-true" v-show="logined">
                <div class="logined-profile">
                    <div>
                        <profile-icon ref="ProfileIcon" id="profile-img"/>
                    </div>
                </div>
                <div class="logined-dropmenu" :style="{ border:'1px solid '+this.$store.state.customColor }">
                    <div class="dropmenu-detail" v-text="'你好, ' + this.$store.state.user.realname" :style="{ color:this.$store.state.customColor }" @mouseenter="hoverDropMenu($event)" @mouseleave="unHoverDropMenu($event)"></div>
                    <div class="dropmenu-detail" :style="{ color:this.$store.state.customColor }" @mouseenter="hoverDropMenu($event)" @mouseleave="unHoverDropMenu($event)">我的信息</div>
                    <div class="dropmenu-detail" :style="{ color:this.$store.state.customColor }" @mouseenter="hoverDropMenu($event)" @mouseleave="unHoverDropMenu($event)">修改密码</div>
                    <div class="dropmenu-detail" @click="logoff" :style="{ color:this.$store.state.customColor }" @mouseenter="hoverDropMenu($event)" @mouseleave="unHoverDropMenu($event)">退出登录</div>
                </div>
            </div>
            <div class="login-false" v-show="!logined" @click="meetUrl('/login');">
                <div v-bind:style="{ border:'2px solid ' + this.$store.state.customColor }">
                    <svg width="44" height="44" t="1551084863977" style="position:relative;left:1px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3739" xmlns:xlink="http://www.w3.org/1999/xlink">
                        <path v-bind:fill="this.$store.state.customColor" d="M887.413 959.473c-16.764 0-30.404-13.656-30.404-30.42 0-92.336-60.086-194.454-115.887-250.289-11.89-11.84-11.89-31.154-0.05-43.028 5.757-5.757 13.395-8.914 21.523-8.914 8.127 0 15.781 3.157 21.522 8.897 66.563 66.563 133.746 183.089 133.746 293.318-0.03 16.781-13.67 30.436-30.45 30.436z m-751.456 0c-16.764 0-30.42-13.656-30.42-30.42 0-110.244 67.168-226.755 133.73-293.3 39.334-39.317 86.795-69.801 138.344-88.952-78.224-46.987-125.946-130.46-125.946-222.749 0-143.38 116.657-260.02 260.036-260.02 143.364 0 260.02 116.64 260.02 260.02 0 143.397-116.656 260.022-260.02 260.022-3.32 0-6.624-0.082-9.911-0.197-82.492 2.323-160.683 36.03-219.494 94.89-55.835 55.834-115.889 157.95-115.889 250.289-0.048 16.762-13.67 30.417-30.45 30.417z m375.695-436.567c1.374 0 2.699 0.099 4.023 0.295 107.711-2.126 195.208-91.307 195.208-199.132 0-109.838-89.361-199.183-199.198-199.183s-199.197 89.345-199.197 199.165c0 106.763 83.505 194.178 190.104 198.985 2.584-0.08 5.822-0.13 9.06-0.13z">

                        </path>
                    </svg>
                </div>
            </div>
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
            ishome: true,
        }
    },
    methods: {
        customChgColor: function(e) {
            this.$store.commit('chgColor', e.target.value);
            document.querySelector(".color-selector").style.display = "none";
        },
        customColorSelector: function() {
            document.querySelector("input[type=color]").click();
        },
        hoverChgColor: function(color, e) {
            e.target.style.backgroundColor = color;
            e.target.style.color = "#fff";
        },
        unHoverChgColor: function(e) {
            e.target.style.backgroundColor = "#fff";
            e.target.style.color = "#111";
        },
        changeColor: function(color) {
            this.$store.commit('chgColor', color);
            document.querySelector(".color-selector").style.display = "none";
        },
        hoverDropMenu: function(e) {
            e.target.style.color = "#fff";
            e.target.style.backgroundColor = this.$store.state.customColor;
        },
        unHoverDropMenu: function(e) {
            e.target.style.color = this.$store.state.customColor;
            e.target.style.backgroundColor = "#fff";
        },
        meetUrl: function(url) {
            router.push(url);
        },
        showUserMenu: function() {

        },
        logoff: function() {
            this.$setcookie("usession", "", 60);
            window.location.href = "/";
        },
        borderMenu: function(e) {
            e.target.style.border = "2px solid " + this.$store.state.customColor;
            e.target.children[1].style.color = this.$store.state.customColor;
        },
        unBorderMenu: function(e) {
            e.target.style.border = "2px solid #fff";
            e.target.children[1].style.color = "#111";
        },
    },
    mounted() {
        // check session
        let sessid = this.$getcookie("usession");
        if(sessid === "") {
            sessid = "0";
        };
        this.$http
            .get("/api/session/" + sessid)
            .then(resp => {
                if(resp.status === 200){
                    this.logined = !this.logined;
                    if(this.$store.state.user.username !== resp.data["val"]){
                        this.$store.state.user.username = resp.data["val"];
                        this.$refs.ProfileIcon.getProfile();
                    }
                } else {
                    if(router.currentRoute.path !== "/") {
                        router.push("/login");
                    }
                }
            })
            .catch(()=>{
                if(router.currentRoute.path !== "/") {
                    router.push("/login");
                }
            }
        );
        if (router.currentRoute.path !== "/") {
            this.ishome = false;
        } else {
            this.ishome = true;
        };
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
    margin-left: 10px;
}
@media screen and (min-width: 500px){
    .nav > .menu {
        margin-left: 20px;
    }
    .nav > .logo {
        margin-left: 10px;
    }
    .login {
        margin-right: 30px;
    }
}
@media screen and (max-width: 499px){
    .nav > .menu {
        margin-left: 0px;
    }
    .nav > .logo {
        margin-left: 0px;
    }
    .login {
        margin-right: 0px;
    }
    .menu-describe {
        font-size: 11px;
    }
}
.nav > .menu {
    width: 110px;
    border: 2px solid #fff;
    border-radius: 5px;
}
.nav > .menu > .menu-img {
    height: 30px;
    width: 100%;
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
}
.menu-describe {
    width: 100%;
    text-align: center;
}
.login{
    margin-left: auto;
    width:100px;
    height: 50px;
    display: flex;
    align-items: center;
}
.login-true{
    width: 100px;
    margin-left: auto;
}
.login-false{
    width: 100px;
    height: 100%;
    margin-left: auto;
    display: flex;
    justify-content: center;
}
.login-false > div{
    width: 46px;
    height: 46px;
    border-radius: 50%;
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    background-size: 85%;
}
.login-true:hover > .logined-dropmenu{
    display: block;
}
.logined-profile{
    height: 50px;
    width: 100px;
    display: flex;
    justify-content: center;
}
.logined-profile > div {
    height: 50px;
    width: 50px;
    display: block;
}
.logined-dropmenu{
    height: 135px;
    width: 100px;
    display: none;
    position: relative;
    background-color: #fff;
    padding-bottom: 2px;
}
.dropmenu-detail{
    width: 100%;
    height: 30px;
    margin-top: 3px;
    display: flex;
    align-items: center;
    justify-content: center;
}

/* color selector div */
.color-selector{
    display: none;
    position: relative;
    margin-top: 2px;
    width: 100%;
    background-color: #fff;
    padding: 3px;
}
.color-selector > div:first-child {
    height: 30px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.color-selector > div:nth-child(2) > li {
    height: 30px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.color-selector > div:nth-child(2) > li:last-child {
    width: 100%;
}
.color-selector > div:nth-child(2) > li:last-child:hover {
    color: #fff;
    background: -webkit-linear-gradient(left, #ED4337, #79cc92, #1890ff);
}
#li-colorSelector:hover > .color-selector{
    display: block !important;
}
</style>
