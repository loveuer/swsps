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
            <el-menu-item style="float:right;" index="/login">　登 录　</el-menu-item>
        </el-menu>
    </div>
</template>

<script>
export default {
    props: {
        path: String,
    },
    methods: {
        "handleSelect": function(index) {
            if (index === "/" && this.$route.path === "/") {
                console.log("back to 公司主页");
                window.location = "http://safetywing.net:8084";
            };
            return;
        },
    },
};
</script>

<style scoped>
.logo {
    font-size: 18px;
    font-weight: bold;
}
</style>
