<template>
    <el-row>
        <el-menu
            :default-active="defaultActive"
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
                <el-menu-item :index="trainingRecorderCheckIndex">查看</el-menu-item>
                <el-menu-item index="/works/trainingRecorder/modify">修改</el-menu-item>
            </el-submenu>
            <el-submenu index="/works/logRecorder">
                <template slot="title">日志记录</template>
                <el-menu-item index="/works/logRecorder/add">添加</el-menu-item>
                <el-menu-item :index="logRecorderCheckIndex">查看</el-menu-item>
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
                <el-menu-item index="/works/spsRecorder/search">查找备件</el-menu-item>
                <el-menu-item index="/works/spsRecorder/history">备件历史</el-menu-item>
                <el-menu-item index="/works/spsRecorder/stat">备件状态</el-menu-item>
            </el-submenu>
            <el-submenu index="/works/maintRecorder">
                <template slot="title">维护记录</template>
                <el-menu-item index="/works/maintRecorder/add">添加</el-menu-item>
                <el-menu-item index="/works/maintRecorder/check">查看</el-menu-item>
            </el-submenu>
            <el-submenu index="/self" style="float:right;">
                <template slot="title">赵育鹏</template>
                <el-menu-item index="/self/">我的信息</el-menu-item>
                <el-menu-item index="/self/chgpswd">修改密码</el-menu-item>
                <el-menu-item index="/logoff">退出登录</el-menu-item>
            </el-submenu>
        </el-menu>
    </el-row>
</template>

<script>
export default {
    data() {
        return {
            defaultActive: '/works',
        };
    },
    computed: {
        trainingRecorderCheckIndex: () => {
            let dateNow = new Date(Date.now());
            return "/works/trainingRecorder/check/" + dateNow.toISOString().slice(0,10);
        },
        logRecorderCheckIndex: () => {
            let dateNow = new Date(Date.now());
            return "/works/logRecorder/check/" + dateNow.toISOString().slice(0,10);
        },
    },
    methods: {
        
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