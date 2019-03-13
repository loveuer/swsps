<template>
    <div class="u-container">
        <div class="u-header">
            <div class="header-logo" @click="changeContent(0)">SafetyWing</div>
            <div class="header-menu">
                <li>
                    <el-button type="text" :style="{fontSize:'16px'}">公司首页</el-button>
                </li>
                <li>
                    <el-dropdown>
                        <span class="el-dropdown-link">
                            <el-button type="text" :style="{fontSize:'16px'}">赵育鹏</el-button>
                            <i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item>我的信息</el-dropdown-item>
                            <el-dropdown-item>修改密码</el-dropdown-item>
                            <el-dropdown-item>退出登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </li>
            </div>
        </div>
        <div class="u-content">
            <div class="u-aside">
                <el-col>
                    <el-menu
                        default-active="0"
                        class="el-menu-vertical-demo"
                        @open="handleOpen"
                        @close="handleClose">
                        <el-menu-item index="1" @click="changeContent(1)">
                            <i class="el-icon-menu"></i>
                            <span slot="title">训练管理</span>
                        </el-menu-item>

                        <el-menu-item index="2" @click="changeContent(2)">
                            <i class="el-icon-document"></i>
                            <span slot="title">日志管理</span>
                        </el-menu-item>

                        <el-menu-item index="3" @click="changeContent(3)">
                            <i class="el-icon-setting"></i>
                            <span slot="title">故障管理</span>
                        </el-menu-item>

                        <el-menu-item index="4" @click="changeContent(4)">
                            <i class="el-icon-setting"></i>
                            <span slot="title">备件管理</span>
                        </el-menu-item>

                        <el-menu-item index="5" @click="changeContent(5)">
                            <i class="el-icon-setting"></i>
                            <span slot="title">系统统计</span>
                        </el-menu-item>

                        <el-menu-item index="6" @click="changeContent(6)">
                            <i class="el-icon-setting"></i>
                            <span slot="title">模拟机信息</span>
                        </el-menu-item>
                    </el-menu>
                </el-col>
            </div>
            <div class="u-main">
                <component :is="workSpace"></component>
            </div>
        </div>
    </div>
</template>

<script>
import router from "../router";

import swSummary from "./Summary.vue"
import trainingRecorder from "./TrainingRecorder.vue";
import logRecoder from "./LogRecorder.vue";
import malfRecoder from "./MalfRecoder.vue";
import spsRecoder from "./SpsRecoder.vue";
import sysRecoder from "./SysRecoder.vue";
import simsImformation from "./SimsImformation.vue";


export default {
    data() {
        return {
            workSpace: "swSummary",
        }
    },
    methods: {
        handleOpen: (key, keyPath) => {
            console.log(key, keyPath);
        },
        handleClose: (key, keyPath) => {
            console.log(key, keyPath);
        },
        changeContent: function(index) {
            let contentMap = {
                0: "swSummary",
                1: "trainingRecorder",
                2: "logRecoder",
                3: "malfRecoder",
                4: "spsRecoder",
                5: "sysRecoder",
                6: "simsImformation",
            };
            this.workSpace = contentMap[index];
        },
    },
    components: {
        "swSummary": swSummary, 
        "trainingRecorder": trainingRecorder, 
        "logRecoder": logRecoder, 
        "malfRecoder": malfRecoder, 
        "spsRecoder": spsRecoder, 
        "sysRecoder": sysRecoder, 
        "simsImformation": simsImformation,
    },
    
};
</script>

<style scoped>
.u-container{
    min-height: 100%;
    width: 100%;
}
.u-header{
    height: 50px;
    width: 100%;
    background-color: #eee;
    display: flex;
}
.header-logo{
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
    font-weight: bold;
    margin-left: 30px;
}
.header-menu{
    margin-left: auto;
    margin-right: 10px;
    display: flex;
    align-items: center;
}
.header-menu > li{
    list-style: none;
    margin-right: 20px;
}
.u-content {
    min-height: 100%;
    width: 100%;
    display: flex;
}
</style>
