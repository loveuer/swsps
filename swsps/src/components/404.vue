<template>
    <div id="container">
        <div class="notfound">
            <div>
                <h1>404 ! oops~~~</h1>
            </div>
            <div>
                <h3>page not found</h3>
            </div>
            <div>
                <el-button @click="meetHome">返回主页</el-button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        meetHome() {
            this.$router.push("/");
        },
    },
}
</script>

<style scoped>
.notfound {
    width: 100%;
    margin-top: 50px;
    display: flex;
    flex-wrap: wrap;
}
.notfound > div {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>
 