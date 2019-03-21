<template>
    <div>
        <el-row>
            <el-menu
                :default-active="workMenuIndex"
                class="el-menu-demo"
                mode="horizontal"
                background-color="#545c64"
                text-color="#fff"
                :router="true"
                active-text-color="#ffd04b">
                <el-menu-item index="/works">
                    <font class="font-logo">Safety</font><font class="font-logo">Wing</font>
                </el-menu-item>
                <el-submenu index="/works/trainingRecorder">
                    <template slot="title">训练记录</template>
                    <el-menu-item index="/works/trainingRecorder/add">添加</el-menu-item>
                    <el-menu-item index="/works/trainingRecorder/check">查看</el-menu-item>
                    <el-menu-item index="/works/trainingRecorder/modify">修改</el-menu-item>
                </el-submenu>
                <el-submenu index="/works/logRecorder">
                    <template slot="title">日志记录</template>
                    <el-menu-item index="/works/logRecorder/add">添加</el-menu-item>
                    <el-menu-item index="/works/logRecorder/check">查看</el-menu-item>
                    <el-menu-item index="/works/logRecorder/modify">修改</el-menu-item>
                </el-submenu>
                <el-submenu index="/works/malfRecorder">
                    <template slot="title">故障记录</template>
                    <el-menu-item index="/works/malfRecorder/add">添加</el-menu-item>
                    <el-menu-item index="/works/malfRecorder/check">查看</el-menu-item>
                    <el-menu-item index="/works/malfRecorder/modify">修改</el-menu-item>
                </el-submenu>
                <el-submenu index="/works/spsRecorder">
                    <template slot="title">备件记录</template>
                    <el-menu-item index="/works/spsRecorder/add">添加</el-menu-item>
                    <el-menu-item index="/works/spsRecorder/check">查看</el-menu-item>
                    <el-menu-item index="/works/spsRecorder/modify">修改</el-menu-item>
                </el-submenu>
                <el-submenu index="/self" style="float:right;">
                    <template slot="title">赵育鹏</template>
                    <el-menu-item index="/self/">我的信息</el-menu-item>
                    <el-menu-item index="/self/chgpswd">修改密码</el-menu-item>
                    <el-menu-item index="/logoff">退出登录</el-menu-item>
                </el-submenu>
            </el-menu>
        </el-row>
    </div>
</template>

<script>
export default {
    data() {
        return {
            workMenuIndex: '0',
        }
    },
    mounted() {
        let workMenuIndexMap = {'traningRecorder':'1', 'logRecorder':'2', 'malfRecorder':'3', 'spsRecorder':'4'};
        let routePath = this.$route.params;
        let pathMode = routePath.mode | 0;
        let pathDetail = routePath.detail | 0;
        let pathKey = routePath.key | 0;
        let pathOther = routePath.other | 0;
        console.log(pathMode, pathDetail, pathKey, pathOther);
    },
    methods: {
        handleOpen:function() {

        },
        handleClose: function() {

        },
    },
};
</script>

<style scoped>
.font-logo{
    height: 60px;
    line-height: 60px;
    font-size: 20px;
    font-weight: bold;
}
.font-logo:first-child{
    color: #7ec0d4;
}
.font-logo:last-child{
    color: #85c59a;
}
</style>
