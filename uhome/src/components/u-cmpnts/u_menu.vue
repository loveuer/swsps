<template>
    <div id="container">
        <div @click="meetUrl('/menu1')">menu 1</div>
        <div>menu 2</div>
        <div @click="meetUrl('/')"></div>
        <div>menu 3</div>
        <div>
            <el-dropdown @command="handleYddMenu">
                <span class="el-dropdown-link">
                    傻乎乎tools<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item command="/ydd/salary">工资条邮件</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        meetUrl: function(url) {
            console.log("function: meet url");
            this.$router.push(url);
        },
        handleYddMenu: function(url) {
            this.meetUrl(url);
        },
    },
};
</script>

<style scoped>
#container { 
    height: 50px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 15px;
}
#container > div {
    width: 100px;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 20px;
    cursor: pointer;
    color: #606266;
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
}
#container > div:nth-child(3) {
    background-image: url("~@/assets/img/uhome.png");
    background-repeat: no-repeat;
    background-size: contain;
    background-position: center;
}
</style>
