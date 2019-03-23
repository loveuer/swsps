<template>
    <div id="u-container">
        <div class="u-header">
            <div>
                <div class="header-logo" @click="reloadPage"></div>
                <div class="header-menu">
                    <li @click="reloadPage">
                        <el-button type="text" size="medium" :style="{fontSize:'16px'}">公司首页</el-button>
                    </li>
                    <li @click="meetUrl('/works')">
                        <el-button type="text" size="medium" :style="{fontSize:'16px'}">公司事务</el-button>
                    </li>
                    <li @click="meetUrl('/admin')">
                        <el-button type="text" size="medium" :style="{fontSize:'16px'}">后台管理</el-button>
                    </li>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    components: {
        
    },
    methods: {
        reloadPage: () => {
            window.location.reload(true);
        },
        meetUrl: function(url) {
            this.$router.push(url);
        },
    },
};
</script>

<style scoped>
#u-container{
    min-height: 100%;
    height: 100%;
    width: 100%;
}
.u-header{
    height: 120px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.u-header > div {
    height: 100%;
    width: 90%;
    display: flex;
    align-items: center;
}
.header-logo {
    height: 60px;
    width: 200px;
    background-image: url("~@/assets/img/logo2.png");
    background-repeat: no-repeat;
    background-size: contain;
    background-position: center;
    cursor: pointer;
}
.header-menu{
    margin-left: auto;
    display: flex;
}
.header-menu > li {
    list-style: none;
    margin-right: 30px;
    cursor: pointer;
}
</style>
