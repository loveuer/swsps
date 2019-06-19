<template>
    <div id="container">
        <el-menu
            :default-active="path"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
            :router="true"
        >
            <el-menu-item index="/">
                <font class="logo">Safety</font>
                <font class="logo">Wing</font>
            </el-menu-item>
            <el-menu-item index="/search">搜索备件</el-menu-item>
            <el-menu-item index="/addsp">新增备件</el-menu-item>
            <el-menu-item index="/sphis">备件历史</el-menu-item>
            <el-menu-item index="#status" @click="showNotDoneWarning">备件状态</el-menu-item>
            <el-menu-item index="#admin" @click="showNotDoneWarning">后台管理</el-menu-item>
            <el-menu-item style="float:right;" index="/login" v-if="this.$store.state.user.id === 0">　登 录　</el-menu-item>
            <el-submenu v-if="this.$store.state.user.id !== 0" style="float:right;" index="/user">
                <template slot="title">　{{ this.$store.state.user.realname }}　</template>
                <el-menu-item index="/user/imf">我的信息</el-menu-item>
                <el-menu-item index="/user/chgpswd">修改密码</el-menu-item>
                <el-menu-item index="#logout" @click="logout">退出登录</el-menu-item>
            </el-submenu>
        </el-menu>
    </div>
</template>

<script>
export default {
    props: {
        path: String,
    },
    methods: {
        handleSelect: function(index) {
            if (index === "/" && this.$route.path === "/") {
                console.log("back to 公司主页");
                window.location = "http://safetywing.net:8084";
            };
            return;
        },
        logout: function() {
            this.setCookie("sessid", "", 0);
            this.$router.push("/login");
        },
        setCookie: function(cname, cvalue, expire_min) {
            var d = new Date();
            d.setTime(d.getTime() + (expire_min*60*1000));
            var expires = "expires="+ d.toUTCString();
            document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
        },
        getCookie: function(cname) {
            let name = cname + "=";
            let decodedCookie = decodeURIComponent(document.cookie);
            var ca = decodedCookie.split(';');
            for(var i = 0; i <ca.length; i++) {
                var c = ca[i];
                while (c.charAt(0) == ' ') {
                    c = c.substring(1);
                }
                if (c.indexOf(name) == 0) {
                    return c.substring(name.length, c.length);
                }
            }
            return "";
        },
        showNotDoneWarning: function() {
            this.$message({showClose:true,type:'warning',message:'功能开发中...'});
        },
    },
    mounted() {
        if (this.$store.state.user.id === 0) {
            if (this.getCookie('sessid') === "") {
                this.$router.push('/login');
                return;
            };
            this.$http.get("/api/user/get/realname/bysession")
                .then(resp => {
                    let user = {id: 1, realname: resp.data.realname};
                    this.$store.commit('setUser', user);
                })
                .catch(err => { 
                    console.log(err);
                    switch(err.response.status){
                        case 401:
                            this.$router.push("/login");
                            break;
                        case 500:
                            this.$message({showClose: true, message: '服务器发生未知错误, 重试或联系管理员', type: 'warning',});
                            this.$router.push("/login");
                            break;
                    }
                });
        };
    }
};
</script>

<style scoped>
.logo {
    font-size: 18px;
    font-weight: bold;
}
</style>
