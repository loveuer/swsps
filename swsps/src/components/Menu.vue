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
            <el-menu-item index="/history">备件历史</el-menu-item>
            <el-menu-item index="/status">备件状态</el-menu-item>
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
        },
        setCookie: function(cname, cvalue, expire_min) {
            var d = new Date();
            d.setTime(d.getTime() + (expire_min*60*1000));
            var expires = "expires="+ d.toUTCString();
            document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
        },
    },
    mounted() {
        if (this.$store.state.user.id === 0) {
            this.$http.get("/api/user/get/username")
                .then(resp => {

                })
                .catch(err => {
                    switch(err.response.status){
                        case 401:
                            this.$http.router.push("/login");
                            break;
                        case 500:
                            this.$message({showClose: true, message: '服务器发生未知错误, 重试或联系管理员', type: 'warning',});
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
