<template>
    <div id="container">
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{path: '/'}">首页</el-breadcrumb-item>
            <el-breadcrumb-item>余丹丹</el-breadcrumb-item>
            <el-breadcrumb-item>工资条</el-breadcrumb-item>
        </el-breadcrumb>
    </div>
</template>

 <style scoped>
.container {
    width: 100%;
    height: 50px;
    margin-top: 10px;
} 
.container > div {
    width: 50px;
    height: 50px;
    background-image: url("~@/assets/img/uhome.png");
    background-repeat: no-repeat;
    background-size: contain;
    background-position: center;
}
 </style>
 